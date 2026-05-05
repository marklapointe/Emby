package repository

import (
	"database/sql"
	"fmt"

	"gopkg.in/yaml.v3"
)

// ConfigRepository handles server configuration persistence.
type ConfigRepository struct {
	*BaseRepository
}

// ServerConfig represents the server configuration stored in the database.
type ServerConfig struct {
	ServerName                  string `json:"ServerName"`
	PreferredMetadataLanguage    string `json:"PreferredMetadataLanguage"`
	MetadataCountryCode         string `json:"MetadataCountryCode"`
	UICulture                   string `json:"UICulture"`
	EnableUPnP                  bool   `json:"EnableUPnP"`
	EnableHttps                 bool   `json:"EnableHttps"`
	PublicHttpsPort             int    `json:"PublicHttpsPort"`
	HttpPort                    int    `json:"HttpPort"`
	HttpsPort                   int    `json:"HttpsPort"`
	PublicPort                  int    `json:"PublicPort"`
	PublicMappedPort            int    `json:"PublicMappedPort"`
	IsStartupWizardCompleted     bool   `json:"IsStartupWizardCompleted"`
	CachePath                  string `json:"CachePath"`
	EnableMetrics               bool   `json:"EnableMetrics"`
	EnableRealtimeMonitor       bool   `json:"EnableRealtimeMonitor"`
	WarnedAboutRealtimeMonitor  bool   `json:"WarnedAboutRealtimeMonitor"`
	EnableAutomaticRestart      bool   `json:"EnableAutomaticRestart"`
	EnableServer                bool   `json:"EnableServer"`
	AutoRunWebApp               bool   `json:"AutoRunWebApp"`
	Logger                      string `json:"Logger"`
	LoggerLevel                 string `json:"LoggerLevel"`
}

// NewConfigRepository creates a new config repository.
func NewConfigRepository(db *sql.DB) *ConfigRepository {
	return &ConfigRepository{NewBaseRepository(db)}
}

// CreateConfigTable creates the config table if it doesn't exist.
func (r *ConfigRepository) CreateConfigTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS server_config (
		id INTEGER PRIMARY KEY CHECK (id = 1),
		config_yaml TEXT NOT NULL,
		updated_at TEXT NOT NULL
	);`
	_, err := r.Exec(query)
	return err
}

// GetConfig retrieves the server configuration.
func (r *ConfigRepository) GetConfig() (*ServerConfig, error) {
	query := `SELECT config_yaml FROM server_config WHERE id = 1`

	var configYAML string
	err := r.QueryRow(query).Scan(&configYAML)
	if err == sql.ErrNoRows {
		// Return default config if none exists
		return r.GetDefaultConfig(), nil
	}
	if err != nil {
		return nil, err
	}

	var config ServerConfig
	if err := yaml.Unmarshal([]byte(configYAML), &config); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the server configuration.
func (r *ConfigRepository) SaveConfig(config *ServerConfig) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	query := `
	INSERT INTO server_config (id, config_yaml, updated_at)
	VALUES (1, ?, datetime('now'))
	ON CONFLICT(id) DO UPDATE SET
		config_yaml = excluded.config_yaml,
		updated_at = datetime('now');`

	_, err = r.Exec(query, string(data))
	return err
}

// GetString retrieves a specific config value by key.
func (r *ConfigRepository) GetString(key string) (string, error) {
	config, err := r.GetConfig()
	if err != nil {
		return "", err
	}

	switch key {
	case "ServerName":
		return config.ServerName, nil
	case "PreferredMetadataLanguage":
		return config.PreferredMetadataLanguage, nil
	case "MetadataCountryCode":
		return config.MetadataCountryCode, nil
	case "UICulture":
		return config.UICulture, nil
	}
	return "", fmt.Errorf("unknown config key: %s", key)
}

// GetDefaultConfig returns the default server configuration.
func (r *ConfigRepository) GetDefaultConfig() *ServerConfig {
	return &ServerConfig{
		ServerName:               "Emby Server",
		PreferredMetadataLanguage: "en",
		MetadataCountryCode:      "US",
		UICulture:                "en-US",
		EnableUPnP:               false,
		EnableHttps:              false,
		PublicHttpsPort:          8920,
		HttpPort:                 8096,
		HttpsPort:                8920,
		PublicPort:               8096,
		PublicMappedPort:         8096,
		IsStartupWizardCompleted: false,
		CachePath:                "data/cache",
		EnableMetrics:            false,
		EnableRealtimeMonitor:     true,
		EnableAutomaticRestart:    false,
		EnableServer:             true,
		AutoRunWebApp:            false,
		Logger:                   "console",
		LoggerLevel:              "info",
	}
}
