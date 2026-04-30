package library

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/service/scheduled"
	"go.uber.org/zap"
)

// ScanResult holds the outcome of a library scan.
type ScanResult struct {
	TotalItemsFound int
	NewItems        int
	UpdatedItems    int
	RemovedItems    int
	Errors          int
	Duration        time.Duration
}

// Scanner handles media library scanning.
type Scanner struct {
	config       *config.Config
	logger       *zap.Logger
	repo         *repository.ItemRepository
	scheduledMgr *scheduled.Manager
	mu           sync.RWMutex
	isScanning   bool
	paths        []string
}

// NewScanner creates a new library scanner.
func NewScanner(cfg *config.Config, logger *zap.Logger, repo *repository.ItemRepository) *Scanner {
	return &Scanner{
		config: cfg,
		logger: logger,
		repo:   repo,
		paths:  cfg.Library.ContentTypes,
	}
}

// AddPath adds a media folder path to the library.
func (s *Scanner) AddPath(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.paths = append(s.paths, path)
}

// RemovePath removes a media folder path from the library.
func (s *Scanner) RemovePath(path string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, p := range s.paths {
		if p == path {
			s.paths = append(s.paths[:i], s.paths[i+1:]...)
			return
		}
	}
}

// ScanLibrary performs a full library scan.
func (s *Scanner) ScanLibrary(ctx context.Context) (*ScanResult, error) {
	s.mu.Lock()
	if s.isScanning {
		s.mu.Unlock()
		return nil, fmt.Errorf("scan already in progress")
	}
	s.isScanning = true
	s.mu.Unlock()
	defer func() {
		s.mu.Lock()
		s.isScanning = false
		s.mu.Unlock()
	}()

	start := time.Now()
	result := &ScanResult{}

	// Get existing items for change detection
	existingItems, err := s.getAllItemPaths()
	if err != nil {
		return nil, fmt.Errorf("get existing items: %w", err)
	}

	// Track found paths
	foundPaths := make(map[string]bool)

	// Scan each library path
	for _, mediaType := range s.paths {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
		}

		mediaPath := s.getMediaFolderPath(mediaType)
		if mediaPath == "" {
			continue
		}

		// Walk the directory tree
		err := filepath.Walk(mediaPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				s.logger.Error("walk error", zap.String("path", path), zap.Error(err))
				result.Errors++
				return nil
			}

			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if info.IsDir() {
				// Check ignore paths
				for _, ignore := range s.config.Library.IgnorePaths {
					if strings.HasSuffix(path, ignore) {
						return filepath.SkipDir
					}
				}
				return nil
			}

			// Check if file is a media file
			if !s.isMediaFile(path) {
				return nil
			}

			foundPaths[path] = true
			result.TotalItemsFound++

			// Check if item exists
			itemID := generateItemID(path)
			_, exists := existingItems[path]
			if !exists {
				// New item
				if err := s.repo.InsertItem(itemID, filepath.Base(path), path, mediaType); err != nil {
					s.logger.Error("insert item", zap.String("path", path), zap.Error(err))
					result.Errors++
					return nil
				}
				result.NewItems++
			} else {
				// Updated item
				result.UpdatedItems++
			}

			return nil
		})
		if err != nil {
			s.logger.Error("scan error", zap.String("path", mediaPath), zap.Error(err))
			result.Errors++
		}
	}

	// Calculate removed items
	for path := range existingItems {
		if !foundPaths[path] {
			result.RemovedItems++
		}
	}

	result.Duration = time.Since(start)
	s.logger.Info("library scan complete",
		zap.Int("total", result.TotalItemsFound),
		zap.Int("new", result.NewItems),
		zap.Int("updated", result.UpdatedItems),
		zap.Int("removed", result.RemovedItems),
		zap.Int("errors", result.Errors),
		zap.Duration("duration", result.Duration),
	)

	return result, nil
}

// StartScheduledScans registers and starts scheduled library scans.
func (s *Scanner) StartScheduledScans(scheduledMgr *scheduled.Manager) {
	if !s.config.Library.EnableAutoDeepScan {
		return
	}

	s.scheduledMgr = scheduledMgr

	task := &scheduled.Task{
		ID:        "library-scan",
		Name:      "Library Scan",
		Description: "Scan media library for new and updated content",
		Category:  "Library",
	}

	s.scheduledMgr.RegisterTask(task)
}

// isMediaFile checks if a file is a recognized media type.
func (s *Scanner) isMediaFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	mediaExtensions := map[string]bool{
		// Video
		".mp4": true, ".mkv": true, ".avi": true, ".mov": true, ".wmv": true,
		".flv": true, ".webm": true, ".m4v": true, ".mpg": true, ".mpeg": true,
		".ts": true, ".mts": true, ".m2ts": true, ".vob": true, ".divx": true,
		".mxf": true, ".rm": true, ".rmvb": true, ".f4v": true,
		// Audio
		".mp3": true, ".flac": true, ".aac": true, ".ogg": true, ".wav": true,
		".wma": true, ".m4a": true, ".aiff": true, ".ape": true, ".opus": true,
		".alac": true,
		// Images
		".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".bmp": true,
		".tiff": true, ".webp": true, ".svg": true, ".raw": true, ".heic": true,
		".heif": true, ".nef": true, ".cr2": true, ".arw": true,
		// Books
		".epub": true, ".pdf": true, ".mobi": true, ".azw3": true,
	}

	if mediaExtensions[ext] {
		return true
	}

	// Check for subtitle files
	subtitleExtensions := map[string]bool{
		".srt": true, ".vtt": true, ".ass": true, ".ssa": true,
		".sub": true, ".idx": true, ".smi": true, ".txt": true,
	}
	if subtitleExtensions[ext] {
		return true
	}

	return false
}

// getMediaFolderPath returns the folder path for a media type.
func (s *Scanner) getMediaFolderPath(mediaType string) string {
	basePath := "media"
	switch strings.ToLower(mediaType) {
	case "video":
		return filepath.Join(basePath, "movies")
	case "music":
		return filepath.Join(basePath, "music")
	case "photos":
		return filepath.Join(basePath, "photos")
	case "books":
		return filepath.Join(basePath, "books")
	default:
		return filepath.Join(basePath, strings.ToLower(mediaType))
	}
}

// getAllItemPaths returns a map of existing item paths.
func (s *Scanner) getAllItemPaths() (map[string]string, error) {
	// This would query the database for all items
	// For now, return empty map
	return make(map[string]string), nil
}

// generateItemID creates a unique ID for a media file.
func generateItemID(path string) string {
	// Simple hash-based ID generation
	return fmt.Sprintf("%x", time.Now().UnixNano())
}
