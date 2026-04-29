package library

import (
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/go-chi/chi/v5"
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

// RegisterRoutes registers library API routes.
func (m *Manager) RegisterRoutes(r chi.Router) {
	// TODO: Add library endpoints
}
