package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// PlaylistHandler handles playlist-related API endpoints.
type PlaylistHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewPlaylistHandler creates a new playlist handler.
func NewPlaylistHandler(cfg *config.Config, repo *repository.ItemRepository) *PlaylistHandler {
	return &PlaylistHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetPlaylists handles GET /Playlists
func (h *PlaylistHandler) GetPlaylists(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	isFavorite := r.URL.Query().Get("IsFavorite")
	_ = isFavorite

	playlists, err := h.repo.GetPlaylists(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlists)
}

// GetPlaylist handles GET /Playlists/{id}
func (h *PlaylistHandler) GetPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	playlist, err := h.repo.GetPlaylist(playlistId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlist)
}

// CreatePlaylist handles POST /Playlists
func (h *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "playlist created"})
}

// UpdatePlaylist handles PUT /Playlists/{id}
func (h *PlaylistHandler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req
	_ = playlistId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "playlist updated"})
}

// DeletePlaylist handles DELETE /Playlists/{id}
func (h *PlaylistHandler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	_ = playlistId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "playlist deleted"})
}

// GetPlaylistItems handles GET /Playlists/{id}/Items
func (h *PlaylistHandler) GetPlaylistItems(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	items, err := h.repo.GetPlaylistItems(playlistId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// AddToPlaylist handles POST /Playlists/{id}/Items
func (h *PlaylistHandler) AddToPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req
	_ = playlistId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "item added"})
}

// RemoveFromPlaylist handles DELETE /Playlists/{id}/Items
func (h *PlaylistHandler) RemoveFromPlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	var req map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = req
	_ = playlistId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "item removed"})
}

// GetDynamicPlaylist handles GET /Playlists/Dynamic
func (h *PlaylistHandler) GetDynamicPlaylist(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	// Return dynamic playlist items
	items := []map[string]interface{}{
		{"Name": "Recently Added", "Type": "playlist"},
		{"Name": "Frequently Played", "Type": "playlist"},
		{"Name": "Continue Watching", "Type": "playlist"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
