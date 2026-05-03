package plugin

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// Plugin represents a plugin for Emby Server.
type Plugin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Author      string `json:"author"`
	TargetAPI   string `json:"targetApi"`
	EntryPoint  string `json:"entryPoint"`
	Assembly    string `json:"assembly"`
	ConfigFile  string `json:"configFile"`
	IsActive    bool   `json:"isActive"`
	LastError   string `json:"lastError"`
}

// Manager handles plugin loading and lifecycle.
type Manager struct {
	plugins   map[string]*Plugin
	mu        sync.RWMutex
	configDir string
}

// NewManager creates a new plugin manager.
func NewManager(configDir string) *Manager {
	return &Manager{
		plugins:   make(map[string]*Plugin),
		configDir: configDir,
	}
}

// LoadPlugins scans the plugin directory and loads all plugins.
func (m *Manager) LoadPlugins() error {
	pluginDir := filepath.Join(m.configDir, "plugins")
	
	entries, err := os.ReadDir(pluginDir)
	if err != nil {
		return fmt.Errorf("failed to read plugin directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			pluginPath := filepath.Join(pluginDir, entry.Name())
			if err := m.loadPlugin(pluginPath); err != nil {
				return fmt.Errorf("failed to load plugin %s: %w", entry.Name(), err)
			}
		}
	}

	return nil
}

// loadPlugin loads a single plugin from a directory.
func (m *Manager) loadPlugin(pluginPath string) error {
	manifestPath := filepath.Join(pluginPath, "manifest.json")
	
	data, err := os.ReadFile(manifestPath)
	if err != nil {
		return fmt.Errorf("failed to read manifest: %w", err)
	}

	var plugin Plugin
	if err := json.Unmarshal(data, &plugin); err != nil {
		return fmt.Errorf("failed to parse manifest: %w", err)
	}

	plugin.ID = filepath.Base(pluginPath)
	plugin.IsActive = true
	plugin.ConfigFile = filepath.Join(pluginPath, "config.json")

	m.mu.Lock()
	m.plugins[plugin.ID] = &plugin
	m.mu.Unlock()

	return nil
}

// GetPlugins returns all loaded plugins.
func (m *Manager) GetPlugins() []*Plugin {
	m.mu.RLock()
	defer m.mu.RUnlock()

	plugins := make([]*Plugin, 0, len(m.plugins))
	for _, p := range m.plugins {
		plugins = append(plugins, p)
	}

	return plugins
}

// GetPlugin returns a plugin by ID.
func (m *Manager) GetPlugin(id string) (*Plugin, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	plugin, ok := m.plugins[id]
	if !ok {
		return nil, fmt.Errorf("plugin not found: %s", id)
	}

	return plugin, nil
}

// EnablePlugin enables a plugin.
func (m *Manager) EnablePlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	plugin, ok := m.plugins[id]
	if !ok {
		return fmt.Errorf("plugin not found: %s", id)
	}

	plugin.IsActive = true
	return nil
}

// DisablePlugin disables a plugin.
func (m *Manager) DisablePlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	plugin, ok := m.plugins[id]
	if !ok {
		return fmt.Errorf("plugin not found: %s", id)
	}

	plugin.IsActive = false
	return nil
}

// UninstallPlugin uninstalls a plugin.
func (m *Manager) UninstallPlugin(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	pluginPath := filepath.Join(m.configDir, "plugins", id)
	if err := os.RemoveAll(pluginPath); err != nil {
		return fmt.Errorf("failed to remove plugin directory: %w", err)
	}

	delete(m.plugins, id)
	return nil
}

// InstallPlugin installs a plugin from a URL.
func (m *Manager) InstallPlugin(url string) error {
	// For now, return a placeholder
	_ = url
	return nil
}

// UpdatePlugin updates a plugin to the latest version.
func (m *Manager) UpdatePlugin(id string) error {
	// For now, return a placeholder
	_ = id
	return nil
}
