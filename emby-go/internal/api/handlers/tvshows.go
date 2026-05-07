package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// TvShowsHandler handles TV shows-related API endpoints.
type TvShowsHandler struct {
	repo *repository.ItemRepository
}

// NewTVShowsHandler creates a new TV shows handler.
func NewTVShowsHandler(repo *repository.ItemRepository) *TvShowsHandler {
	return &TvShowsHandler{
		repo: repo,
	}
}

// GetTVShows handles GET /TvShows
func (h *TvShowsHandler) GetTVShows(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.SearchItems("Series", 50, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"Items": items,
		"TotalRecordCount": len(items),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetTVShow handles GET /TvShows/{id}
func (h *TvShowsHandler) GetTVShow(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := h.repo.GetItem(id)
	if err != nil {
		http.Error(w, "TV show not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// GetNextUp handles GET /Shows/NextUp
func (h *TvShowsHandler) GetNextUp(w http.ResponseWriter, r *http.Request) {
	episodes, err := h.repo.SearchItems("Episode", 20, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"Items":      episodes,
		"NextUpTime": 0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetSeasons handles GET /TvShows/{id}/Seasons
func (h *TvShowsHandler) GetSeasons(w http.ResponseWriter, r *http.Request) {
	seriesId := chi.URLParam(r, "id")

	items, err := h.repo.SearchItems("Season", 50, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter seasons by series
	var seasons []map[string]interface{}
	for _, item := range items {
		if parentId, ok := item["ParentId"].(string); ok && parentId == seriesId {
			seasons = append(seasons, item)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seasons)
}

// GetEpisodes handles GET /TvShows/{id}/Episodes
func (h *TvShowsHandler) GetEpisodes(w http.ResponseWriter, r *http.Request) {
	seriesId := chi.URLParam(r, "id")

	items, err := h.repo.SearchItems("Episode", 100, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter episodes by series
	var episodes []map[string]interface{}
	for _, item := range items {
		if parentId, ok := item["ParentId"].(string); ok && parentId == seriesId {
			episodes = append(episodes, item)
		}
	}

	result := map[string]interface{}{
		"Items": episodes,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetSimilar handles GET /TvShows/{id}/Similar
func (h *TvShowsHandler) GetSimilar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	items, err := h.repo.SearchItems("Series", 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter out current item
	var similar []map[string]interface{}
	for _, item := range items {
		if itemId, ok := item["Id"].(string); ok && itemId != id {
			similar = append(similar, item)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"Items": similar})
}

// GetGenres handles GET /TvShows/Genres
func (h *TvShowsHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}
