package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// MoviesHandler handles movies-related API endpoints.
type MoviesHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewMoviesHandler creates a new movies handler.
func NewMoviesHandler(cfg *config.Config, repo *repository.ItemRepository) *MoviesHandler {
	return &MoviesHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetMovies handles GET /Movies
func (h *MoviesHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	movies := []map[string]interface{}{
		{"Name": "Sample Movie", "Type": "Movie", "Id": "movie-1"},
	}

	result := map[string]interface{}{
		"Items": movies,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetUpcomingMovies handles GET /Movies/Upcoming
func (h *MoviesHandler) GetUpcomingMovies(w http.ResponseWriter, r *http.Request) {
	movies := []map[string]interface{}{
		{"Name": "Upcoming Movie", "Type": "Movie", "Id": "movie-2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// GetRecommendations handles GET /Movies/Recommendations
func (h *MoviesHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	recommendations := []map[string]interface{}{
		{"Name": "Recommended Movie", "Type": "Movie", "Id": "movie-3"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}

// GetGenres handles GET /Movies/Genres
func (h *MoviesHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres := []map[string]interface{}{
		{"Name": "Action", "Id": "genre-1"},
		{"Name": "Comedy", "Id": "genre-2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// GetStudios handles GET /Movies/Studios
func (h *MoviesHandler) GetStudios(w http.ResponseWriter, r *http.Request) {
	studios := []map[string]interface{}{
		{"Name": "Studio One", "Id": "studio-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}

// GetCollections handles GET /Movies/Collections
func (h *MoviesHandler) GetCollections(w http.ResponseWriter, r *http.Request) {
	collections := []map[string]interface{}{
		{"Name": "Collection One", "Id": "collection-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(collections)
}

// GetSimilar handles GET /Movies/{id}/Similar
func (h *MoviesHandler) GetSimilar(w http.ResponseWriter, r *http.Request) {
	similar := []map[string]interface{}{
		{"Name": "Similar Movie", "Type": "Movie", "Id": "movie-4"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(similar)
}
