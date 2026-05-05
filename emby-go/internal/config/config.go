package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config holds the entire application configuration.
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Library  LibraryConfig  `yaml:"library"`
	Logging  LoggingConfig  `yaml:"logging"`
	Stream   StreamConfig   `yaml:"stream_pooling"`
	Wizard   WizardConfig   `yaml:"wizard"`
}

// ServerConfig holds HTTP server settings.
type ServerConfig struct {
	Host           string `yaml:"host"`
	Port           int    `yaml:"port"`
	MaxHeaderBytes int    `yaml:"max_header_bytes"`
	ReadTimeout    int    `yaml:"read_timeout"`
	WriteTimeout   int    `yaml:"write_timeout"`
	PublicPort     int    `yaml:"public_port"`
	PublicHTTPSPort int    `yaml:"public_https_port"`
}

// DatabaseConfig holds SQLite database settings.
type DatabaseConfig struct {
	Path              string `yaml:"path"`
	MaxOpenConns      int    `yaml:"max_open_conns"`
	MaxIdleConns      int    `yaml:"max_idle_conns"`
	ConnMaxLifetime   int    `yaml:"conn_max_lifetime"`
	EnableWAL         bool   `yaml:"enable_wal"`
	PRAGMAJournalMode string `yaml:"pragma_journal_mode"`
}

// LibraryConfig holds media library settings.
type LibraryConfig struct {
	ScanIntervalMinutes int      `yaml:"scan_interval_minutes"`
	EnableAutoDeepScan  bool     `yaml:"enable_auto_deep_scan"`
	ContentTypes        []string `yaml:"content_types"`
	IgnorePaths         []string `yaml:"ignore_paths"`
}

// LoggingConfig holds logging settings.
type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"` // "json" or "console"
}

// StreamConfig holds stream pooling settings.
type StreamConfig struct {
	Enabled              bool     `yaml:"enabled"`
	MaxConcurrentStreams int      `yaml:"max_concurrent_streams"`
	IdleTimeout          string   `yaml:"idle_timeout"`
	HealthCheckInterval  string   `yaml:"health_check_interval"`
	MetricsEnabled       bool     `yaml:"metrics_enabled"`
}

// WizardConfig holds startup wizard settings.
type WizardConfig struct {
	IsCompleted bool   `yaml:"is_completed"`
	ServerName string `yaml:"server_name"`
}

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:            "0.0.0.0",
			Port:            8096,
			MaxHeaderBytes:  1 << 20, // 1 MB
			ReadTimeout:     30,
			WriteTimeout:    30,
			PublicPort:      8096,
			PublicHTTPSPort: 8920,
		},
		Database: DatabaseConfig{
			Path:              "data/emby-server.db",
			MaxOpenConns:      25,
			MaxIdleConns:      10,
			ConnMaxLifetime:   300,
			EnableWAL:         true,
			PRAGMAJournalMode: "WAL",
		},
		Library: LibraryConfig{
			ScanIntervalMinutes: 24,
			EnableAutoDeepScan:  true,
			ContentTypes:        []string{"Video", "Music", "Photos", "Books"},
			IgnorePaths:         []string{".cache", "tmp"},
		},
		Logging: LoggingConfig{
			Level:  "info",
			Format: "json",
		},
		Stream: StreamConfig{
			Enabled:              true,
			MaxConcurrentStreams: 50,
			IdleTimeout:          "5m",
			HealthCheckInterval:  "30s",
			MetricsEnabled:       true,
		},
		Wizard: WizardConfig{
			IsCompleted: false,
			ServerName: "Emby Server",
		},
	}
}

// LoadConfig reads a YAML config file and applies environment variable overrides.
func LoadConfig(path string) (*Config, error) {
	cfg := DefaultConfig()

	if path == "" {
		// Try default locations
		candidates := []string{
			"configs/default.yaml",
			"config.yaml",
			"/etc/emby-server/config.yaml",
		}
		for _, c := range candidates {
			if _, err := os.Stat(c); err == nil {
				path = c
				break
			}
		}
	}

	if path != "" {
		data, err := os.ReadFile(path)
		if err != nil {
			return nil, fmt.Errorf("read config file: %w", err)
		}
		if err := yaml.Unmarshal(data, cfg); err != nil {
			return nil, fmt.Errorf("parse config file: %w", err)
		}
	}

	// Apply environment variable overrides
	if v := os.Getenv("EMBY_SERVER_HOST"); v != "" {
		cfg.Server.Host = v
	}
	if v := os.Getenv("EMBY_SERVER_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &cfg.Server.Port)
	}
	if v := os.Getenv("EMBY_DATABASE_PATH"); v != "" {
		cfg.Database.Path = v
	}
	if v := os.Getenv("EMBY_LOG_LEVEL"); v != "" {
		cfg.Logging.Level = v
	}

	return cfg, nil
}

// SaveConfig writes the config to a YAML file.
func (c *Config) SaveConfig(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("write config file: %w", err)
	}
	return nil
}
