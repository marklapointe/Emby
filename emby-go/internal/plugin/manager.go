package plugin

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
)

// Plugin represents a server plugin.
type Plugin struct {
	ID            string
	Name          string
	Version       string
	Description   string
	Author        string
	Category      string
	EntryPoint    string
	AssemblyPath  string
	IsActive      bool
	Dependencies  []string
	Config        map[string]interface{}
	Logger        *zap.Logger
}

// Manager manages server plugins.
type Manager struct {
	mu       sync.RWMutex
	plugins  map[string]*Plugin
	logger   *zap.Logger
}

// NewManager creates a new plugin manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		plugins: make(map[string]*Plugin),
		logger:  logger,
	}
}

// RegisterPlugin registers a new plugin.
func (m *Manager) RegisterPlugin(p *Plugin) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.plugins[p.ID]; exists {
		return fmt.Errorf("plugin already registered: %s", p.ID)
	}

	m.plugins[p.ID] = p
	m.logger.Info("plugin registered", zap.String("id", p.ID), zap.String("name", p.Name))
	return nil
}

// GetPlugin returns a plugin by ID.
func (m *Manager) GetPlugin(id string) (*Plugin, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	p, exists := m.plugins[id]
	return p, exists
}

// GetPlugins returns all registered plugins.
func (m *Manager) GetPlugins() []*Plugin {
	m.mu.RLock()
	defer m.mu.RUnlock()

	plugins := make([]*Plugin, 0, len(m.plugins))
	for _, p := range m.plugins {
		plugins = append(plugins, p)
	}
	return plugins
}

// ActivatePlugin activates a plugin.
func (m *Manager) ActivatePlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, exists := m.plugins[id]
	if !exists {
		return fmt.Errorf("plugin not found: %s", id)
	}

	p.IsActive = true
	m.logger.Info("plugin activated", zap.String("id", p.ID))
	return nil
}

// DeactivatePlugin deactivates a plugin.
func (m *Manager) DeactivatePlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	p, exists := m.plugins[id]
	if !exists {
		return fmt.Errorf("plugin not found: %s", id)
	}

	p.IsActive = false
	m.logger.Info("plugin deactivated", zap.String("id", p.ID))
	return nil
}

// UnregisterPlugin removes a plugin.
func (m *Manager) UnregisterPlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.plugins[id]; !exists {
		return fmt.Errorf("plugin not found: %s", id)
	}

	delete(m.plugins, id)
	m.logger.Info("plugin unregistered", zap.String("id", id))
	return nil
}

// GetActivePluginCount returns the number of active plugins.
func (m *Manager) GetActivePluginCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, p := range m.plugins {
		if p.IsActive {
			count++
		}
	}
	return count
}
