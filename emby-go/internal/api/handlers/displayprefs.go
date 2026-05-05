package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// DisplayPrefsHandler handles display preferences-related API endpoints.
type DisplayPrefsHandler struct {
	repo *repository.ItemRepository
}

// NewDisplayPrefsHandler creates a new display prefs handler.
func NewDisplayPrefsHandler(repo *repository.ItemRepository) *DisplayPrefsHandler {
	return &DisplayPrefsHandler{
		repo: repo,
	}
}

// GetDisplayPreferences handles GET /DisplayPreferences/{id}
func (h *DisplayPrefsHandler) GetDisplayPreferences(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	prefs, err := h.repo.GetDisplayPreferences(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prefs)
}

// UpdateDisplayPreferences handles POST /DisplayPreferences/{id}
func (h *DisplayPrefsHandler) UpdateDisplayPreferences(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = id
	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
