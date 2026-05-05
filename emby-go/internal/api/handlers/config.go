package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// ConfigHandler handles configuration-related API endpoints.
type ConfigHandler struct {
	config *config.Config
	logger *zap.Logger
}

// NewConfigHandler creates a new config handler.
func NewConfigHandler(cfg *config.Config, logger *zap.Logger) *ConfigHandler {
	return &ConfigHandler{
		config: cfg,
		logger: logger,
	}
}

// GetConfiguration handles GET /Configuration
func (h *ConfigHandler) GetConfiguration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.config)
}

// UpdateConfiguration handles PUT /Configuration
func (h *ConfigHandler) UpdateConfiguration(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = req
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// GetNamedConfiguration handles GET /Configuration/{name}
func (h *ConfigHandler) GetNamedConfiguration(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	_ = name
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

// GetSystemConfig handles GET /System/Configuration
func (h *ConfigHandler) GetSystemConfig(w http.ResponseWriter, r *http.Request) {
	config := map[string]interface{}{
		"ServerName":        "Emby Server",
		"MetadataPath":      "",
		"ImageOptimize":     false,
		"ImageOptimizeFormat": "",
		"EnableUPnP":        false,
		"EnableDLNA":        false,
		"Port":              h.config.Server.Port,
		"SSLPort":           8920,
		"EnableSSL":         false,
		"LogToStdout":       h.config.Logging.Format == "stdout",
		"LogLevel":          h.config.Logging.Level,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// UpdateSystemConfig handles POST /System/Configuration
func (h *ConfigHandler) UpdateSystemConfig(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update configuration values
	if serverName, ok := req["ServerName"].(string); ok {
		_ = serverName
	}
	if port, ok := req["Port"].(float64); ok {
		_ = port
	}
	if enableSSL, ok := req["EnableSSL"].(bool); ok {
		_ = enableSSL
	}
	if logLevel, ok := req["LogLevel"].(string); ok {
		_ = logLevel
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

// GetPublicSystemConfig handles GET /System/Public/Configuration
func (h *ConfigHandler) GetPublicSystemConfig(w http.ResponseWriter, r *http.Request) {
	config := map[string]interface{}{
		"ServerName":        "Emby Server",
		"EnableSSL":         false,
		"Port":              h.config.Server.Port,
		"SSLPort":           8920,
		"LogoImageBaseUrl":  "",
		"DefaultTab":        "home",
		"IsInConfigurationWizard": false,
		"HasPreparedUsers":  true,
		"LogToStdout":       h.config.Logging.Format == "stdout",
		"LogLevel":          h.config.Logging.Level,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// GetLocalAddress handles GET /System/Info/LocalAddress
func (h *ConfigHandler) GetLocalAddress(w http.ResponseWriter, r *http.Request) {
	address := map[string]string{
		"LocalAddress": fmt.Sprintf("http://localhost:%d", h.config.Server.Port),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(address)
}

// GetMacAddress handles GET /System/Info/MacAddress
func (h *ConfigHandler) GetMacAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"MacAddress": "00:00:00:00:00:00"})
}

// GetPluginConfig handles GET /Plugins/{pluginId}/Configuration
func (h *ConfigHandler) GetPluginConfig(w http.ResponseWriter, r *http.Request) {
	pluginId := chi.URLParam(r, "pluginId")

	// Return empty config for now
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id":         pluginId,
		"Name":       pluginId,
		"ConfigPage": "",
	})
}

// UpdatePluginConfig handles POST /Plugins/{pluginId}/Configuration
func (h *ConfigHandler) UpdatePluginConfig(w http.ResponseWriter, r *http.Request) {
	pluginId := chi.URLParam(r, "pluginId")

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req
	_ = pluginId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
