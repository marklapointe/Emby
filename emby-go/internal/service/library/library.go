package library

import (
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"go.uber.org/zap"
)

// Manager handles library-related operations.
type Manager struct {
	config    *config.Config
	logger    *zap.Logger
	dbManager *database.Manager
}

// NewManager creates a new library manager.
func NewManager(cfg *config.Config, logger *zap.Logger, dbManager *database.Manager) *Manager {
	return &Manager{
		config:    cfg,
		logger:    logger,
		dbManager: dbManager,
	}
}


