package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// StartupHandler handles startup wizard-related API endpoints.
type StartupHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewStartupHandler creates a new startup handler.
func NewStartupHandler(cfg *config.Config, repo *repository.ItemRepository) *StartupHandler {
	return &StartupHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetStartupConfig handles GET /startup/config
func (h *StartupHandler) GetStartupConfig(w http.ResponseWriter, r *http.Request) {
	config := map[string]interface{}{
		"IsInConfigurationWizard": true,
		"HasPreparedUsers":        false,
		"LogToStdout":             h.config.Logging.Format == "stdout",
		"LogLevel":                h.config.Logging.Level,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(config)
}

// PostStartupConfig handles POST /startup/config
func (h *StartupHandler) PostStartupConfig(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "config saved"})
}

// GetStartupDashboardInfo handles GET /startup/dashboard
func (h *StartupHandler) GetStartupDashboardInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"ServerName":        "Emby Server",
		"Version":           "0.1.0",
		"OperatingSystem":   "Linux",
		"OperatingSystemVersion": "unknown",
		"OperatingSystemArchitecture": "x64",
		"IsInConfigurationWizard": true,
		"HasPreparedUsers":  false,
		"LogToStdout":       h.config.Logging.Format == "stdout",
		"LogLevel":          h.config.Logging.Level,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// PostUser handles POST /startup/user
func (h *StartupHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "user created"})
}

// GetStartupLanguage handles GET /startup/language
func (h *StartupHandler) GetStartupLanguage(w http.ResponseWriter, r *http.Request) {
	languages := []map[string]interface{}{
		{"Name": "English", "Code": "en"},
		{"Name": "Spanish", "Code": "es"},
		{"Name": "French", "Code": "fr"},
		{"Name": "German", "Code": "de"},
		{"Name": "Italian", "Code": "it"},
		{"Name": "Portuguese", "Code": "pt"},
		{"Name": "Russian", "Code": "ru"},
		{"Name": "Japanese", "Code": "ja"},
		{"Name": "Korean", "Code": "ko"},
		{"Name": "Chinese", "Code": "zh"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languages)
}

// GetStartupTimezone handles GET /startup/timezone
func (h *StartupHandler) GetStartupTimezone(w http.ResponseWriter, r *http.Request) {
	timezones := []map[string]interface{}{
		{"Name": "UTC", "Offset": 0},
		{"Name": "America/New_York", "Offset": -5},
		{"Name": "America/Chicago", "Offset": -6},
		{"Name": "America/Denver", "Offset": -7},
		{"Name": "America/Los_Angeles", "Offset": -8},
		{"Name": "Europe/London", "Offset": 0},
		{"Name": "Europe/Paris", "Offset": 1},
		{"Name": "Europe/Berlin", "Offset": 1},
		{"Name": "Asia/Tokyo", "Offset": 9},
		{"Name": "Asia/Shanghai", "Offset": 8},
		{"Name": "Australia/Sydney", "Offset": 10},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(timezones)
}

// GetStartupRemoteAccess handles GET /startup/remoteaccess
func (h *StartupHandler) GetStartupRemoteAccess(w http.ResponseWriter, r *http.Request) {
	remoteAccess := map[string]interface{}{
		"EnableRemoteAccess": true,
		"Port":               h.config.Server.Port,
		"SSLPort":            8920,
		"EnableSSL":          false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(remoteAccess)
}

// PostStartupRemoteAccess handles POST /startup/remoteaccess
func (h *StartupHandler) PostStartupRemoteAccess(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "remote access configured"})
}
