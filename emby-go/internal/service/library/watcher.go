package library

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
)

type Watcher struct {
	config    *config.Config
	logger    *zap.Logger
	repo      *repository.ItemRepository
	watcher   *fsnotify.Watcher
	paths     []string
	mu        sync.RWMutex
	ctx       context.Context
	cancel    context.CancelFunc
	isRunning bool
}

func NewWatcher(cfg *config.Config, logger *zap.Logger, repo *repository.ItemRepository) *Watcher {
	return &Watcher{
		config: cfg,
		logger: logger,
		repo:   repo,
		paths:  cfg.Library.ContentTypes,
	}
}

func (w *Watcher) Start(ctx context.Context) error {
	w.mu.Lock()
	if w.isRunning {
		w.mu.Unlock()
		return nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		w.mu.Unlock()
		return err
	}
	w.watcher = watcher
	w.ctx, w.cancel = context.WithCancel(ctx)
	w.isRunning = true
	w.mu.Unlock()

	for _, mediaType := range w.paths {
		basePath := w.getMediaFolderPath(mediaType)
		if basePath == "" {
			continue
		}
		if err := w.watcher.Add(basePath); err != nil {
			w.logger.Warn("failed to watch path", zap.String("path", basePath), zap.Error(err))
		} else {
			w.logger.Info("watching path", zap.String("path", basePath))
		}
	}

	go w.run()

	return nil
}

func (w *Watcher) Stop() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if !w.isRunning {
		return nil
	}

	if w.cancel != nil {
		w.cancel()
	}
	if w.watcher != nil {
		w.watcher.Close()
	}
	w.isRunning = false
	return nil
}

func (w *Watcher) AddPath(path string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.watcher != nil {
		if err := w.watcher.Add(path); err != nil {
			return err
		}
	}
	w.paths = append(w.paths, path)
	return nil
}

func (w *Watcher) run() {
	for {
		select {
		case <-w.ctx.Done():
			return
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			w.handleEvent(event)
		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			w.logger.Error("watcher error", zap.Error(err))
		}
	}
}

func (w *Watcher) handleEvent(event fsnotify.Event) {
	if event.Has(fsnotify.Create) {
		w.logger.Info("file created", zap.String("path", event.Name))
		w.onFileCreated(event.Name)
	} else if event.Has(fsnotify.Write) {
		w.logger.Debug("file modified", zap.String("path", event.Name))
		w.onFileModified(event.Name)
	} else if event.Has(fsnotify.Remove) {
		w.logger.Info("file removed", zap.String("path", event.Name))
		w.onFileRemoved(event.Name)
	} else if event.Has(fsnotify.Rename) {
		w.logger.Info("file renamed", zap.String("path", event.Name))
		w.onFileRemoved(event.Name)
	}
}

func (w *Watcher) onFileCreated(path string) {
	if !w.isMediaFile(path) {
		return
	}

	info, err := os.Stat(path)
	if err != nil {
		return
	}
	if info.IsDir() {
		w.AddPath(path)
		return
	}

	itemID := generateItemID(path)
	name := filepath.Base(path)
	mediaType := w.getMediaType(path)

	if err := w.repo.InsertItem(itemID, name, path, mediaType); err != nil {
		w.logger.Error("failed to insert item", zap.String("path", path), zap.Error(err))
		return
	}

	w.logger.Info("item added to library", zap.String("name", name), zap.String("type", mediaType))
}

func (w *Watcher) onFileModified(path string) {
	if !w.isMediaFile(path) {
		return
	}
	w.logger.Debug("item modified", zap.String("path", path))
}

func (w *Watcher) onFileRemoved(path string) {
	if !w.isMediaFile(path) {
		return
	}
	w.logger.Info("item removed from library", zap.String("path", path))
}

func (w *Watcher) isMediaFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	mediaExtensions := map[string]bool{
		".mp4": true, ".mkv": true, ".avi": true, ".mov": true, ".wmv": true,
		".flv": true, ".webm": true, ".m4v": true, ".mpg": true, ".mpeg": true,
		".ts": true, ".mts": true, ".m2ts": true, ".vob": true, ".divx": true,
		".mp3": true, ".flac": true, ".aac": true, ".ogg": true, ".wav": true,
		".wma": true, ".m4a": true,
	}
	return mediaExtensions[ext]
}

func (w *Watcher) getMediaFolderPath(mediaType string) string {
	basePath := "media"
	switch strings.ToLower(mediaType) {
	case "video":
		return filepath.Join(basePath, "movies")
	case "music":
		return filepath.Join(basePath, "music")
	case "photos":
		return filepath.Join(basePath, "photos")
	default:
		return filepath.Join(basePath, strings.ToLower(mediaType))
	}
}

func (w *Watcher) getMediaType(path string) string {
	ext := strings.ToLower(filepath.Ext(path))
	audioExts := map[string]bool{
		".mp3": true, ".flac": true, ".aac": true, ".ogg": true, ".wav": true,
		".wma": true, ".m4a": true, ".aiff": true, ".ape": true, ".opus": true,
	}
	if audioExts[ext] {
		return "music"
	}
	return "video"
}