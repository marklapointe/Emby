package media

import (
	"github.com/emby/emby-go/internal/config"
	"go.uber.org/zap"
)

// Manager handles media-related operations.
type Manager struct {
	config *config.Config
	logger *zap.Logger
}

// NewManager creates a new media manager.
func NewManager(cfg *config.Config, logger *zap.Logger) *Manager {
	return &Manager{
		config: cfg,
		logger: logger,
	}
}
