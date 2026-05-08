package repository

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

type ConfigRepository struct {
	*BaseRepository
}

type ServerConfig struct {
	ID                           uint      `gorm:"primaryKey"`
	ServerName                   string    `json:"ServerName"`
	PreferredMetadataLanguage    string    `json:"PreferredMetadataLanguage"`
	MetadataCountryCode          string    `json:"MetadataCountryCode"`
	UICulture                    string    `json:"UICulture"`
	EnableUPnP                   bool      `json:"EnableUPnP"`
	EnableHttps                  bool      `json:"EnableHttps"`
	PublicHttpsPort              int       `json:"PublicHttpsPort"`
	HttpPort                     int       `json:"HttpPort"`
	HttpsPort                    int       `json:"HttpsPort"`
	PublicPort                   int       `json:"PublicPort"`
	PublicMappedPort             int       `json:"PublicMappedPort"`
	IsStartupWizardCompleted     bool      `json:"IsStartupWizardCompleted"`
	CachePath                   string    `json:"CachePath"`
	EnableMetrics                bool      `json:"EnableMetrics"`
	EnableRealtimeMonitor        bool      `json:"EnableRealtimeMonitor"`
	WarnedAboutRealtimeMonitor   bool      `json:"WarnedAboutRealtimeMonitor"`
	EnableAutomaticRestart       bool      `json:"EnableAutomaticRestart"`
	EnableServer                 bool      `json:"EnableServer"`
	AutoRunWebApp               bool      `json:"AutoRunWebApp"`
	Logger                       string    `json:"Logger"`
	LoggerLevel                  string    `json:"LoggerLevel"`
	UpdatedAt                    time.Time `json:"UpdatedAt"`
}

func (ServerConfig) TableName() string {
	return "server_config"
}

func NewConfigRepository(db *gorm.DB) *ConfigRepository {
	return &ConfigRepository{NewBaseRepository(db)}
}

func (r *ConfigRepository) CreateConfigTable() error {
	return r.db.AutoMigrate(&ServerConfig{})
}

func (r *ConfigRepository) GetConfig() (*ServerConfig, error) {
	var config ServerConfig
	err := r.db.First(&config, 1).Error
	if err == gorm.ErrRecordNotFound {
		return r.GetDefaultConfig(), nil
	}
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (r *ConfigRepository) SaveConfig(config *ServerConfig) error {
	config.ID = 1
	config.UpdatedAt = time.Now()
	return r.db.Save(config).Error
}

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

func (r *ConfigRepository) GetDefaultConfig() *ServerConfig {
	return &ServerConfig{
		ServerName:                   "Emby Server",
		PreferredMetadataLanguage:     "en",
		MetadataCountryCode:           "US",
		UICulture:                     "en-US",
		EnableUPnP:                   false,
		EnableHttps:                  false,
		PublicHttpsPort:              8920,
		HttpPort:                     8096,
		HttpsPort:                    8920,
		PublicPort:                   8096,
		PublicMappedPort:             8096,
		IsStartupWizardCompleted:     false,
		CachePath:                   "data/cache",
		EnableMetrics:                false,
		EnableRealtimeMonitor:       true,
		EnableAutomaticRestart:       false,
		EnableServer:                true,
		AutoRunWebApp:              false,
		Logger:                      "console",
		LoggerLevel:                 "info",
	}
}

func YAMLToConfig(yamlData []byte) (*ServerConfig, error) {
	var config ServerConfig
	if err := yaml.Unmarshal(yamlData, &config); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &config, nil
}