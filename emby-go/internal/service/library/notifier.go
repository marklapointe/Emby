package library

import (
	"context"
	"path/filepath"
	"sync"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

// Notifier watches media folders for file system changes.
type Notifier struct {
	config       *config.Config
	logger       *zap.Logger
	repo         *repository.ItemRepository
	watcher      *fsnotify.Watcher
	mu           sync.Mutex
	running      bool
	mediaPaths   []string
	scanQueue    chan string
	scanDelay    time.Duration
}

// NewNotifier creates a new file system notifier.
func NewNotifier(cfg *config.Config, logger *zap.Logger, repo *repository.ItemRepository) *Notifier {
	return &Notifier{
		config:    cfg,
		logger:    logger,
		repo:      repo,
		mediaPaths: cfg.Library.ContentTypes,
		scanDelay:  2 * time.Second,
	}
}

// Start begins watching media folders for changes.
func (n *Notifier) Start(ctx context.Context) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if n.running {
		return nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	n.watcher = watcher
	n.running = true
	n.scanQueue = make(chan string, 100)

	// Watch each media path
	for _, path := range n.mediaPaths {
		if err := watcher.Add(path); err != nil {
			n.logger.Error("failed to watch path", zap.String("path", path), zap.Error(err))
		} else {
			n.logger.Info("watching path", zap.String("path", path))
		}
	}

	// Start scan queue processor
	go n.processScanQueue(ctx)

	// Start watcher event loop
	go n.watchEvents(ctx)

	return nil
}

// Stop stops watching media folders.
func (n *Notifier) Stop() {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.running {
		return
	}

	n.running = false
	if n.watcher != nil {
		n.watcher.Close()
	}
}

// AddPath adds a new path to watch.
func (n *Notifier) AddPath(path string) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.running {
		return nil
	}

	if err := n.watcher.Add(path); err != nil {
		return err
	}

	n.mediaPaths = append(n.mediaPaths, path)
	n.logger.Info("added watch path", zap.String("path", path))
	return nil
}

// RemovePath removes a path from watching.
func (n *Notifier) RemovePath(path string) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.running {
		return
	}

	n.watcher.Remove(path)

	for i, p := range n.mediaPaths {
		if p == path {
			n.mediaPaths = append(n.mediaPaths[:i], n.mediaPaths[i+1:]...)
			break
		}
	}

	n.logger.Info("removed watch path", zap.String("path", path))
}

// watchEvents processes file system events.
func (n *Notifier) watchEvents(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case event, ok := <-n.watcher.Events:
			if !ok {
				return
			}

			// Only care about media file changes
			if n.isMediaFile(event.Name) {
				n.logger.Debug("file system event",
					zap.String("file", event.Name),
					zap.String("op", event.String()))

				// Queue a scan with debounce
				select {
				case n.scanQueue <- event.Name:
				default:
					// Queue full, skip
				}
			}
		case err, ok := <-n.watcher.Errors:
			if !ok {
				return
			}
			n.logger.Error("watcher error", zap.Error(err))
		}
	}
}

// processScanQueue debounces and processes scan requests.
func (n *Notifier) processScanQueue(ctx context.Context) {
	ticker := time.NewTicker(n.scanDelay)
	defer ticker.Stop()

	var lastEvent string
	var timer *time.Timer

	for {
		select {
		case <-ctx.Done():
			return
		case event := <-n.scanQueue:
			lastEvent = event
			if timer == nil {
				timer = time.NewTimer(n.scanDelay)
			} else {
				timer.Reset(n.scanDelay)
			}
		case <-ticker.C:
			if lastEvent != "" {
				n.logger.Info("triggering scan due to file change",
					zap.String("file", lastEvent))
				lastEvent = ""
			}
		case <-timer.C:
			if lastEvent != "" {
				n.logger.Info("triggering scan due to file change",
					zap.String("file", lastEvent))
				lastEvent = ""
			}
		}
	}
}

// isMediaFile checks if a file is a media file.
func (n *Notifier) isMediaFile(path string) bool {
	ext := filepath.Ext(path)
	mediaExtensions := []string{
		".mp4", ".mkv", ".avi", ".mov", ".wmv", ".flv", ".webm",
		".mp3", ".flac", ".wav", ".aac", ".ogg", ".wma",
		".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp",
		".nfo", ".nfo",
	}

	for _, mediaExt := range mediaExtensions {
		if ext == mediaExt {
			return true
		}
	}

	return false
}
