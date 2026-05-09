package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

type SubtitleHandler struct {
	repo *repository.ItemRepository
}

func NewSubtitleHandler(repo *repository.ItemRepository) *SubtitleHandler {
	return &SubtitleHandler{repo: repo}
}

func (h *SubtitleHandler) GetRemoteSubtitles(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
	_ = id
}

func (h *SubtitleHandler) SearchRemoteSubtitles(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ItemId   string `json:"ItemId"`
		Language string `json:"Language"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	subtitles, err := h.repo.SearchSubtitles(req.ItemId, req.Language, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subtitles)
}

func (h *SubtitleHandler) DownloadRemoteSubtitles(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var req struct {
		ItemId string `json:"ItemId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.repo.DownloadSubtitle(req.ItemId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SubtitleHandler) DeleteSubtitle(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	itemId := chi.URLParam(r, "itemId")

	err := h.repo.DeleteSubtitle(itemId, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *SubtitleHandler) GetSubtitlePlaylist(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}
