package library

import (
	"context"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/service/scheduled"
	"go.uber.org/zap"
)

type RefreshTask struct {
	scanner *Scanner
	watcher *Watcher
	cfg     *config.Config
	logger  *zap.Logger
}

func NewRefreshTask(cfg *config.Config, logger *zap.Logger, repo *repository.ItemRepository) *RefreshTask {
	return &RefreshTask{
		scanner: NewScanner(cfg, logger, repo),
		watcher: NewWatcher(cfg, logger, repo),
		cfg:     cfg,
		logger:  logger,
	}
}

func (t *RefreshTask) Execute(ctx context.Context) error {
	t.logger.Info("starting library refresh task")
	start := time.Now()

	result, err := t.scanner.ScanLibrary(ctx)
	if err != nil {
		t.logger.Error("library refresh failed", zap.Error(err))
		return err
	}

	t.logger.Info("library refresh completed",
		zap.Int("total", result.TotalItemsFound),
		zap.Int("new", result.NewItems),
		zap.Int("updated", result.UpdatedItems),
		zap.Duration("duration", time.Since(start)),
	)

	return nil
}

func (t *RefreshTask) StartWatcher(ctx context.Context) error {
	return t.watcher.Start(ctx)
}

func (t *RefreshTask) StopWatcher() error {
	return t.watcher.Stop()
}

func RegisterLibraryTasks(mgr *scheduled.Manager, cfg *config.Config, logger *zap.Logger, repo *repository.ItemRepository) {
	task := &scheduled.Task{
		ID:          "library-refresh",
		Name:        "Library Refresh",
		Description: "Scan media library for new and updated content",
		Category:    "Library",
		Options: scheduled.TaskOptions{
			EnableOnStartup: true,
		},
	}
	mgr.RegisterTask(task)
}
