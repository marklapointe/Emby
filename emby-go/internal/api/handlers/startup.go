package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/emby/emby-go/internal/version"
	"github.com/emby/emby-go/internal/model"
	"github.com/emby/emby-go/internal/repository"
	"go.uber.org/zap"
)

// StartupHandler handles startup wizard-related API endpoints.
type StartupHandler struct {
	configRepo *repository.ConfigRepository
	userRepo   *repository.UserRepository
	logger     *zap.Logger
}

// NewStartupHandler creates a new startup handler.
func NewStartupHandler(configRepo *repository.ConfigRepository, userRepo *repository.UserRepository, logger *zap.Logger) *StartupHandler {
	if logger == nil {
		logger = zap.NewNop()
	}
	return &StartupHandler{
		configRepo: configRepo,
		userRepo:   userRepo,
		logger:     logger,
	}
}

// IsFirstRun handles GET /Startup/First
// Returns whether the server needs initial setup
func (h *StartupHandler) IsFirstRun(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		config = h.configRepo.GetDefaultConfig()
	}

	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		h.logger.Error("failed to get users", zap.Error(err))
		users = nil
	}

	result := model.IsFirstRunResult{
		IsFirstRun:  !config.IsStartupWizardCompleted,
		HasPassword: len(users) > 0 && users[0].HasConfiguredPassword,
		HasUsername: len(users) > 0 && users[0].Name != "",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetOptions handles GET /Startup/Options
func (h *StartupHandler) GetOptions(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
	}

	options := model.StartupOptions{
		EnableUPnP: config.EnableUPnP,
		EnableDLNA: false, // DLNA is disabled in this implementation
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

// Complete handles POST /Startup/Complete
// Marks the startup wizard as completed
func (h *StartupHandler) Complete(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	config.IsStartupWizardCompleted = true

	if err := h.configRepo.SaveConfig(config); err != nil {
		h.logger.Error("failed to save config", zap.Error(err))
		http.Error(w, "failed to save config", http.StatusInternalServerError)
		return
	}

	h.logger.Info("startup wizard completed")

	w.WriteHeader(http.StatusNoContent)
}

// GetStartupConfig handles GET /Startup/Configuration
// Returns the current startup configuration
func (h *StartupHandler) GetStartupConfig(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	startupConfig := model.StartupConfiguration{
		UICulture:                   config.UICulture,
		MetadataCountryCode:        config.MetadataCountryCode,
		PreferredMetadataLanguage:   config.PreferredMetadataLanguage,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startupConfig)
}

// PostStartupConfig handles POST /Startup/Configuration
// Updates the startup configuration
func (h *StartupHandler) PostStartupConfig(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DEBUG PostStartupConfig: method=%s contentType=%s\n", r.Method, r.Header.Get("Content-Type"))

	var req model.StartupConfiguration

	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		fmt.Println("DEBUG: handling as JSON")
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.logger.Warn("failed to decode json request", zap.Error(err))
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
	} else {
		fmt.Println("DEBUG: handling as form")
		if err := r.ParseForm(); err != nil {
			h.logger.Warn("failed to parse form request", zap.Error(err))
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		req.UICulture = r.Form.Get("UICulture")
		req.MetadataCountryCode = r.Form.Get("MetadataCountryCode")
		req.PreferredMetadataLanguage = r.Form.Get("PreferredMetadataLanguage")
	}

	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Update config with values from request
	if req.UICulture != "" {
		config.UICulture = req.UICulture
	}
	if req.MetadataCountryCode != "" {
		config.MetadataCountryCode = req.MetadataCountryCode
	}
	if req.PreferredMetadataLanguage != "" {
		config.PreferredMetadataLanguage = req.PreferredMetadataLanguage
	}

	if err := h.configRepo.SaveConfig(config); err != nil {
		h.logger.Error("failed to save config", zap.Error(err))
		http.Error(w, "failed to save config", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetStartupUser handles GET /Startup/User
// Returns the initial user being configured during wizard
func (h *StartupHandler) GetStartupUser(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		h.logger.Error("failed to get users", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(model.StartupUser{})
		return
	}

	user := users[0]
	startupUser := model.StartupUser{
		Name:            user.Name,
		ConnectUserName: user.ConnectUserName,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(startupUser)
}

// PostUser handles POST /Startup/User
// Creates or updates the initial user during wizard
func (h *StartupHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var req model.StartupUser

	contentType := r.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.logger.Warn("failed to decode json request", zap.Error(err))
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
	} else {
		if err := r.ParseForm(); err != nil {
			h.logger.Warn("failed to parse form request", zap.Error(err))
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}
		req.Name = r.Form.Get("Name")
		req.ConnectUserName = r.Form.Get("ConnectUserName")
	}

	users, err := h.userRepo.GetAllUsers()
	if err != nil {
		h.logger.Error("failed to get users", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	var user *repository.User
	if len(users) > 0 {
		user = users[0]
		user.Name = req.Name
		user.ConnectUserName = req.ConnectUserName
		if err := h.userRepo.UpdateUser(user); err != nil {
			h.logger.Error("failed to update user", zap.Error(err))
			http.Error(w, "failed to update user", http.StatusInternalServerError)
			return
		}
	} else {
		// Create new user
		user = &repository.User{
			Name:            req.Name,
			ConnectUserName: req.ConnectUserName,
		}
		if err := h.userRepo.CreateUser(user); err != nil {
			h.logger.Error("failed to create user", zap.Error(err))
			http.Error(w, "failed to create user", http.StatusInternalServerError)
			return
		}
	}

	result := model.UpdateStartupUserResult{
		UserLinkResult: nil, // Connect linking not implemented
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetStartupLanguage handles GET /Localization/Options
// Returns available localization options
func (h *StartupHandler) GetStartupLanguage(w http.ResponseWriter, r *http.Request) {
	languages := []map[string]interface{}{
		{"Name": "English", "Value": "en"},
		{"Name": "Arabic", "Value": "ar"},
		{"Name": "Belarusian", "Value": "be-BY"},
		{"Name": "Bulgarian", "Value": "bg-BG"},
		{"Name": "Catalan", "Value": "ca"},
		{"Name": "Chinese (Simplified)", "Value": "zh-CN"},
		{"Name": "Chinese (Traditional)", "Value": "zh-TW"},
		{"Name": "Czech", "Value": "cs"},
		{"Name": "Danish", "Value": "da"},
		{"Name": "Dutch", "Value": "nl"},
		{"Name": "Finnish", "Value": "fi"},
		{"Name": "French", "Value": "fr"},
		{"Name": "German", "Value": "de"},
		{"Name": "Greek", "Value": "el"},
		{"Name": "Hebrew", "Value": "he"},
		{"Name": "Hindi", "Value": "hi-IN"},
		{"Name": "Hungarian", "Value": "hu"},
		{"Name": "Indonesian", "Value": "id"},
		{"Name": "Italian", "Value": "it"},
		{"Name": "Japanese", "Value": "ja"},
		{"Name": "Korean", "Value": "ko"},
		{"Name": "Norwegian", "Value": "nb"},
		{"Name": "Polish", "Value": "pl"},
		{"Name": "Portuguese (Brazil)", "Value": "pt-BR"},
		{"Name": "Portuguese (Portugal)", "Value": "pt-PT"},
		{"Name": "Romanian", "Value": "ro"},
		{"Name": "Russian", "Value": "ru"},
		{"Name": "Slovak", "Value": "sk"},
		{"Name": "Spanish", "Value": "es"},
		{"Name": "Swedish", "Value": "sv"},
		{"Name": "Thai", "Value": "th"},
		{"Name": "Turkish", "Value": "tr"},
		{"Name": "Ukrainian", "Value": "uk"},
		{"Name": "Vietnamese", "Value": "vi"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(languages)
}

// GetStartupRemoteAccess handles GET /Startup/RemoteAccess
// Returns remote access configuration
func (h *StartupHandler) GetStartupRemoteAccess(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	remoteAccess := model.StartupRemoteAccess{
		EnableRemoteAccess:        !config.EnableUPnP, // Inverse of EnableUPnP for some reason
		EnableAutomaticPortMapping: config.EnableUPnP,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(remoteAccess)
}

// PostStartupRemoteAccess handles POST /Startup/RemoteAccess
// Updates remote access configuration
func (h *StartupHandler) PostStartupRemoteAccess(w http.ResponseWriter, r *http.Request) {
	var req model.StartupRemoteAccess
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Warn("failed to decode request", zap.Error(err))
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	// Note: In the C# code, EnableRemoteAccess is inverted for EnableUPnP
	config.EnableUPnP = req.EnableAutomaticPortMapping

	if err := h.configRepo.SaveConfig(config); err != nil {
		h.logger.Error("failed to save config", zap.Error(err))
		http.Error(w, "failed to save config", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetStartupDashboardInfo handles GET /Startup/Dashboard
func (h *StartupHandler) GetStartupDashboardInfo(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	info := map[string]interface{}{
		"ServerName":                    config.ServerName,
		"Version":                       version.Version,
		"OperatingSystem":                "Linux",
		"OperatingSystemVersion":         "unknown",
		"OperatingSystemArchitecture":    "x64",
		"IsInConfigurationWizard":        !config.IsStartupWizardCompleted,
		"HasPreparedUsers":              true,
		"LogToStdout":                   true,
		"LogLevel":                      "info",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// GetSystemInfoPublic handles GET /System/Info/Public
func (h *StartupHandler) GetSystemInfoPublic(w http.ResponseWriter, r *http.Request) {
	config, err := h.configRepo.GetConfig()
	if err != nil {
		h.logger.Error("failed to get config", zap.Error(err))
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	info := map[string]interface{}{
		"ServerName":        config.ServerName,
		"Version":           version.Version,
		"OperatingSystem":   "Linux",
		"Id":                "emby-go-server-id",
		"LocalAddress":      "http://localhost:8096",
		"LocalAddressV6":    "http://[::1]:8096",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// SetLibraryPath handles setting a library path during wizard
func (h *StartupHandler) SetLibraryPath(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Path string `json:"Path"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// wizardSettings handles GET /wizardsettings
func (h *StartupHandler) GetWizardSettings(w http.ResponseWriter, r *http.Request) {
	wizardSettings := map[string]interface{}{
		"HasWizardCompleted": false,
		"EnableWizard":        true,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wizardSettings)
}
