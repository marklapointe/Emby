package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

type MusicHandler struct {
	repo *repository.ItemRepository
}

func NewMusicHandler(repo *repository.ItemRepository) *MusicHandler {
	return &MusicHandler{repo: repo}
}

func (h *MusicHandler) GetInstantMixFromItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetInstantMixFromArtistId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetInstantMixFromMusicGenreId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetInstantMixFromSong(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetInstantMixFromAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetInstantMixFromPlaylist(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetInstantMixFromMusicGenre(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetSimilarArtists(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *MusicHandler) GetSimilarAlbums(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.GetSimilarItems(id, "Audio", 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
