package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const (
	AppName   = "emby"
	ConfigDir = "emby"
)

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Library  LibraryConfig  `yaml:"library"`
	Logging  LoggingConfig  `yaml:"logging"`
	Stream   StreamConfig   `yaml:"stream_pooling"`
	Wizard   WizardConfig   `yaml:"wizard"`
	Security SecurityConfig `yaml:"security"`
}

type ServerConfig struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	MaxHeaderBytes  int    `yaml:"max_header_bytes"`
	ReadTimeout     int    `yaml:"read_timeout"`
	WriteTimeout    int    `yaml:"write_timeout"`
	PublicPort      int    `yaml:"public_port"`
	PublicHTTPSPort int    `yaml:"public_https_port"`
	CachePath       string `yaml:"cache_path"`
}

type DatabaseConfig struct {
	Type              string `yaml:"type"`
	Path              string `yaml:"path"`
	ConnectionString  string `yaml:"connection_string"`
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Database          string `yaml:"database"`
	MaxOpenConns      int    `yaml:"max_open_conns"`
	MaxIdleConns      int    `yaml:"max_idle_conns"`
	ConnMaxLifetime   int    `yaml:"conn_max_lifetime"`
	EnableWAL         bool   `yaml:"enable_wal"`
	PRAGMAJournalMode string `yaml:"pragma_journal_mode"`
}

type LibraryConfig struct {
	ScanIntervalMinutes int      `yaml:"scan_interval_minutes"`
	EnableAutoDeepScan bool     `yaml:"enable_auto_deep_scan"`
	ContentTypes       []string `yaml:"content_types"`
	IgnorePaths        []string `yaml:"ignore_paths"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
}

type StreamConfig struct {
	Enabled              bool   `yaml:"enabled"`
	MaxConcurrentStreams int    `yaml:"max_concurrent_streams"`
	IdleTimeout          string `yaml:"idle_timeout"`
	HealthCheckInterval  string `yaml:"health_check_interval"`
	MetricsEnabled       bool   `yaml:"metrics_enabled"`
}

type WizardConfig struct {
	IsCompleted bool   `yaml:"is_completed"`
	ServerName  string `yaml:"server_name"`
}

type SecurityConfig struct {
	APIKey        string `yaml:"api_key"`
	APIKeyHash    string `yaml:"api_key_hash"`
	EmbyServerURL string `yaml:"emby_server_url"`
}

type ConfigFile struct {
	Path   string
	Target string
}

func GetConfigDirs() []string {
	var dirs []string

	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		dirs = append(dirs, filepath.Join(xdg, ConfigDir))
	}

	if home := os.Getenv("HOME"); home != "" {
		dirs = append(dirs, filepath.Join(home, ".local", "etc", AppName))
	}

	dirs = append(dirs, "/etc/"+AppName)
	dirs = append(dirs, "/usr/local/etc/"+AppName)
	dirs = append(dirs, "configs")

	return dirs
}

func GetConfigFiles() []ConfigFile {
	return []ConfigFile{
		{Target: "full", Path: "emby.yaml"},
		{Target: "server", Path: "server.yaml"},
		{Target: "database", Path: "database.yaml"},
		{Target: "library", Path: "library.yaml"},
		{Target: "logging", Path: "logging.yaml"},
		{Target: "security", Path: "security.yaml"},
		{Target: "stream", Path: "stream.yaml"},
		{Target: "wizard", Path: "wizard.yaml"},
	}
}

func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Host:            "0.0.0.0",
			Port:            8096,
			MaxHeaderBytes:  1 << 20,
			ReadTimeout:     30,
			WriteTimeout:    30,
			PublicPort:      8096,
			PublicHTTPSPort: 8920,
		},
		Database: DatabaseConfig{
			Type:              "sqlite",
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
		Security: SecurityConfig{},
	}
}

func LoadConfig(overridePath string) (*Config, error) {
	cfg := DefaultConfig()

	if overridePath != "" {
		if err := loadConfigFile(overridePath, cfg); err != nil {
			return nil, err
		}
		applyEnvOverrides(cfg)
		return cfg, nil
	}

	dirs := GetConfigDirs()
	files := GetConfigFiles()

	for _, dir := range dirs {
		for _, file := range files {
			fullPath := filepath.Join(dir, file.Path)
			if _, err := os.Stat(fullPath); err == nil {
				if err := loadConfigFile(fullPath, cfg); err != nil {
					return nil, err
				}
			}
		}
	}

	applyEnvOverrides(cfg)
	return cfg, nil
}

func loadConfigFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config file %s: %w", path, err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("parse config file %s: %w", path, err)
	}

	return nil
}

func applyEnvOverrides(cfg *Config) {
	if v := os.Getenv("EMBY_SERVER_HOST"); v != "" {
		cfg.Server.Host = v
	}
	if v := os.Getenv("EMBY_SERVER_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &cfg.Server.Port)
	}
	if v := os.Getenv("EMBY_DATABASE_TYPE"); v != "" {
		cfg.Database.Type = v
	}
	if v := os.Getenv("EMBY_DATABASE_PATH"); v != "" {
		cfg.Database.Path = v
	}
	if v := os.Getenv("EMBY_DATABASE_CONNECTION_STRING"); v != "" {
		cfg.Database.ConnectionString = v
	}
	if v := os.Getenv("EMBY_DATABASE_HOST"); v != "" {
		cfg.Database.Host = v
	}
	if v := os.Getenv("EMBY_DATABASE_PORT"); v != "" {
		fmt.Sscanf(v, "%d", &cfg.Database.Port)
	}
	if v := os.Getenv("EMBY_DATABASE_USERNAME"); v != "" {
		cfg.Database.Username = v
	}
	if v := os.Getenv("EMBY_DATABASE_PASSWORD"); v != "" {
		cfg.Database.Password = v
	}
	if v := os.Getenv("EMBY_DATABASE_NAME"); v != "" {
		cfg.Database.Database = v
	}
	if v := os.Getenv("EMBY_LOG_LEVEL"); v != "" {
		cfg.Logging.Level = v
	}
	if v := os.Getenv("EMBY_API_KEY"); v != "" {
		cfg.Security.APIKey = v
	}
}

func GetFirstFoundConfig() (string, error) {
	dirs := GetConfigDirs()
	files := GetConfigFiles()

	for _, dir := range dirs {
		for _, file := range files {
			fullPath := filepath.Join(dir, file.Path)
			if _, err := os.Stat(fullPath); err == nil {
				return fullPath, nil
			}
		}
	}
	return "", nil
}

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

func (c *Config) SaveSecurityConfig(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create config dir: %w", err)
	}

	data, err := yaml.Marshal(c.Security)
	if err != nil {
		return fmt.Errorf("marshal security config: %w", err)
	}

	if err := os.WriteFile(path, data, 0600); err != nil {
		return fmt.Errorf("write security config: %w", err)
	}
	return nil
}

func GetSecurityConfigPath(dir string) string {
	return filepath.Join(dir, "security.yaml")
}

func LoadSecurityConfig() (*SecurityConfig, error) {
	cfg := &SecurityConfig{}

	dirs := GetConfigDirs()
	securityFile := "security.yaml"

	for _, dir := range dirs {
		path := filepath.Join(dir, securityFile)
		if _, err := os.Stat(path); err == nil {
			data, err := os.ReadFile(path)
			if err != nil {
				return nil, fmt.Errorf("read security config: %w", err)
			}
			if err := yaml.Unmarshal(data, cfg); err != nil {
				return nil, fmt.Errorf("parse security config: %w", err)
			}
			return cfg, nil
		}
	}

	cfg.APIKey = os.Getenv("EMBY_API_KEY")
	return cfg, nil
}

func (c *Config) HasAPIKey() bool {
	return c.Security.APIKey != ""
}

func (c *Config) SetAPIKey(key string) {
	c.Security.APIKey = key
}

func GetConfigSearchPaths() []string {
	var result []string
	dirs := GetConfigDirs()
	files := GetConfigFiles()

	for _, dir := range dirs {
		for _, file := range files {
			result = append(result, filepath.Join(dir, file.Path))
		}
	}
	return result
}

func GetConfigSearchPathsWithDescriptions() []string {
	var result []string
	dirs := GetConfigDirs()

	descriptions := map[string]string{
		"XDG_CONFIG_HOME/.emby":    "User config (XDG standard)",
		"$HOME/.local/etc/emby":    "User config (Unix convention)",
		"/etc/emby":                "System config",
		"/usr/local/etc/emby":       "System config (local)",
		"configs":                  "Local development",
	}

	for _, dir := range dirs {
		desc := descriptions[dir]
		if desc == "" {
			desc = dir
		}
		result = append(result, fmt.Sprintf("%s (%s)", dir, desc))
	}
	return result
}

func ExpandPath(path string) string {
	if strings.HasPrefix(path, "~/") {
		if home := os.Getenv("HOME"); home != "" {
			path = filepath.Join(home, path[2:])
		}
	}
	return os.Expand(path, func(s string) string {
		return os.Getenv(s)
	})
}
