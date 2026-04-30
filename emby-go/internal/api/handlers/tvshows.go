package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// TvShowsHandler handles TV shows-related API endpoints.
type TvShowsHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewTvShowsHandler creates a new TV shows handler.
func NewTvShowsHandler(cfg *config.Config, repo *repository.ItemRepository) *TvShowsHandler {
	return &TvShowsHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetNextUp handles GET /Shows/NextUp
func (h *TvShowsHandler) GetNextUp(w http.ResponseWriter, r *http.Request) {
	episodes := []map[string]interface{}{
		{"Name": "Next Episode", "Type": "Episode", "Id": "ep-1"},
	}

	result := map[string]interface{}{
		"Items":      episodes,
		"NextUpTime": 0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetShows handles GET /Shows
func (h *TvShowsHandler) GetShows(w http.ResponseWriter, r *http.Request) {
	shows := []map[string]interface{}{
		{"Name": "TV Show", "Type": "Series", "Id": "show-1"},
	}

	result := map[string]interface{}{
		"Items": shows,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetSeasons handles GET /Shows/{seriesId}/Seasons
func (h *TvShowsHandler) GetSeasons(w http.ResponseWriter, r *http.Request) {
	seasons := []map[string]interface{}{
		{"Name": "Season 1", "Id": "season-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seasons)
}

// GetEpisodes handles GET /Shows/{seriesId}/Seasons/{seasonId}/Episodes
func (h *TvShowsHandler) GetEpisodes(w http.ResponseWriter, r *http.Request) {
	episodes := []map[string]interface{}{
		{"Name": "Episode 1", "Id": "ep-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(episodes)
}

// GetStudioShows handles GET /Shows/WithStudio
func (h *TvShowsHandler) GetStudioShows(w http.ResponseWriter, r *http.Request) {
	shows := []map[string]interface{}{
		{"Name": "Studio Show", "Type": "Series", "Id": "show-2"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shows)
}

// GetGenres handles GET /Shows/Genres
func (h *TvShowsHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres := []map[string]interface{}{
		{"Name": "Drama", "Id": "genre-3"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// GetRecommendations handles GET /Shows/Recommendations
func (h *TvShowsHandler) GetRecommendations(w http.ResponseWriter, r *http.Request) {
	recommendations := []map[string]interface{}{
		{"Name": "Recommended Show", "Type": "Series", "Id": "show-3"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recommendations)
}
