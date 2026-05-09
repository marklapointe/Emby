package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
	"github.com/go-chi/chi/v5"
)

// MoviesHandler handles movies-related API endpoints.
type MoviesHandler struct {
	repo *repository.ItemRepository
}

// NewMoviesHandler creates a new movies handler.
func NewMoviesHandler(repo *repository.ItemRepository) *MoviesHandler {
	return &MoviesHandler{
		repo: repo,
	}
}

// GetMovies handles GET /Movies
func (h *MoviesHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.SearchItems("Movie", 50, 0)
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

// GetMovie handles GET /Movies/{id}
func (h *MoviesHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := h.repo.GetItem(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// GetUpcomingMovies handles GET /Movies/Upcoming
func (h *MoviesHandler) GetUpcomingMovies(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.SearchItems("Movie", 20, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GetRecommendations handles GET /Movies/Recommendations
func (h *MoviesHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.SearchItems("Movie", 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"Items": items,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetGenres handles GET /Movies/Genres
func (h *MoviesHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// GetStudios handles GET /Movies/Studios
func (h *MoviesHandler) GetStudios(w http.ResponseWriter, r *http.Request) {
	studios, err := h.repo.GetStudios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}

// GetCollections handles GET /Movies/Collections
func (h *MoviesHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	collections, err := h.repo.GetPlaylists("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collections)
}

// GetSimilar handles GET /Movies/{id}/Similar
func (h *MoviesHandler) GetSimilar(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	item, err := h.repo.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	items, _ := h.repo.SearchItems("Movie", 10, 0)
	filtered := []map[string]interface{}{}
	for _, i := range items {
		if i["Id"] != id {
			filtered = append(filtered, i)
		}
	}

	result := map[string]interface{}{
		"Items": filtered,
		"Item":  item,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *MoviesHandler) GetTrailers(w http.ResponseWriter, r *http.Request) {
	items, err := h.repo.SearchItems("Movie", 50, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result := map[string]interface{}{
		"Items": items,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *MoviesHandler) GetSpecialFeatures(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}
