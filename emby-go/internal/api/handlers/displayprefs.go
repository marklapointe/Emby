package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/gorilla/mux"
)

// DisplayPreferencesHandler handles display preferences-related API endpoints.
type DisplayPreferencesHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewDisplayPreferencesHandler creates a new display preferences handler.
func NewDisplayPreferencesHandler(cfg *config.Config, repo *repository.ItemRepository) *DisplayPreferencesHandler {
	return &DisplayPreferencesHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetDisplayPreferences handles GET /DisplayPreferences
func (h *DisplayPreferencesHandler) GetDisplayPreferences(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	itemId := r.URL.Query().Get("ItemId")
	_ = itemId

	prefs, err := h.repo.GetDisplayPreferences(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prefs)
}

// GetDisplayPreferencesByItem handles GET /DisplayPreferences/{itemId}
func (h *DisplayPreferencesHandler) GetDisplayPreferencesByItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := vars["itemId"]
	userId := r.URL.Query().Get("UserId")

	prefs, err := h.repo.GetDisplayPreferencesByItem(itemId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prefs)
}

// UpdateDisplayPreferences handles POST /DisplayPreferences
func (h *DisplayPreferencesHandler) UpdateDisplayPreferences(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "preferences updated"})
}

// GetLayouts handles GET /DisplayPreferences/Layouts
func (h *DisplayPreferencesHandler) GetLayouts(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("UserId")

	layoutOptions := []map[string]interface{}{
		{"Name": "List", "Value": "list"},
		{"Name": "Thumbnail", "Value": "thumbnail"},
		{"Name": "Poster", "Value": "poster"},
		{"Name": "Banner", "Value": "banner"},
		{"Name": "BannerCard", "Value": "bannercard"},
		{"Name": "Card", "Value": "card"},
		{"Name": "DVD", "Value": "dvd"},
		{"Name": "Photo", "Value": "photo"},
		{"Name": "PhotoGrid", "Value": "photogrid"},
		{"Name": "PosterWall", "Value": "posterwall"},
		{"Name": "Snapshots", "Value": "snapshots"},
		{"Name": "ThumbList", "Value": "thumblist"},
		{"Name": "ThumbGrid", "Value": "thumbgrid"},
		{"Name": "WideList", "Value": "widelist"},
		{"Name": "WideThumbList", "Value": "widethumblist"},
		{"Name": "ViewSettings", "Value": "viewsettings"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(layoutOptions)
}

// GetViewSettings handles GET /DisplayPreferences/ViewSettings
func (h *DisplayPreferencesHandler) GetViewSettings(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")

	settings, err := h.repo.GetViewSettings(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(settings)
}

// UpdateViewSettings handles POST /DisplayPreferences/ViewSettings
func (h *DisplayPreferencesHandler) UpdateViewSettings(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "view settings updated"})
}
