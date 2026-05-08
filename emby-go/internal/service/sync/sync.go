package sync

import (
	"go.uber.org/zap"
)

type Manager struct {
	logger *zap.Logger
}

func NewManager(logger *zap.Logger) *Manager {
	return &Manager{logger: logger}
}