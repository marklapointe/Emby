package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// PlaylistHandler handles playlist-related API endpoints.
type PlaylistHandler struct {
	repo *repository.ItemRepository
}

// NewPlaylistHandler creates a new playlist handler.
func NewPlaylistHandler(repo *repository.ItemRepository) *PlaylistHandler {
	return &PlaylistHandler{
		repo: repo,
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
		http.Error(w, "Playlist not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(playlist)
}

// CreatePlaylist handles POST /Playlists
func (h *PlaylistHandler) CreatePlaylist(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"Name"`
		Ids         string `json:"Ids"`
		UserId      string `json:"UserId"`
		MediaType   string `json:"MediaType"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	playlist, err := h.repo.CreatePlaylist(req.Name, "", req.MediaType, req.UserId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id": playlist.Id,
		"Name": playlist.Name,
	})
}

// UpdatePlaylist handles PUT /Playlists/{id}
func (h *PlaylistHandler) UpdatePlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	var req struct {
		Name     string `json:"Name"`
		Overview string `json:"Overview"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdatePlaylist(playlistId, req.Name, req.Overview); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeletePlaylist handles DELETE /Playlists/{id}
func (h *PlaylistHandler) DeletePlaylist(w http.ResponseWriter, r *http.Request) {
	playlistId := chi.URLParam(r, "id")

	if err := h.repo.DeletePlaylist(playlistId); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

	var req struct {
		Ids   string `json:"Ids"`
		UserId string `json:"UserId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Ids == "" {
		http.Error(w, "Ids is required", http.StatusBadRequest)
		return
	}

	ids := splitIds(req.Ids)
	for i, itemId := range ids {
		if err := h.repo.AddItemToPlaylist(playlistId, itemId, i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

// RemoveFromPlaylist handles DELETE /Playlists/{id}/Items
func (h *PlaylistHandler) RemoveFromPlaylist(w http.ResponseWriter, r *http.Request) {
	chi.URLParam(r, "id")

	var req struct {
		EntryIds string `json:"EntryIds"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.EntryIds == "" {
		http.Error(w, "EntryIds is required", http.StatusBadRequest)
		return
	}

	ids := splitIds(req.EntryIds)
	for _, itemId := range ids {
		if err := h.repo.RemoveItemFromPlaylist(itemId); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func splitIds(ids string) []string {
	if ids == "" {
		return nil
	}
	var result []string
	for _, id := range strings.Split(ids, ",") {
		id = strings.TrimSpace(id)
		if id != "" {
			result = append(result, id)
		}
	}
	return result
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
