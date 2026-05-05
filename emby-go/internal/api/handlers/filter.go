package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
)

// FilterHandler handles filter-related API endpoints.
type FilterHandler struct {
	repo *repository.ItemRepository
}

// NewFilterHandler creates a new filter handler.
func NewFilterHandler(repo *repository.ItemRepository) *FilterHandler {
	return &FilterHandler{
		repo: repo,
	}
}

// GetGenres handles GET /Genres
func (h *FilterHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.repo.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// GetStudios handles GET /Studios
func (h *FilterHandler) GetStudios(w http.ResponseWriter, r *http.Request) {
	studios, err := h.repo.GetStudios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}

// GetYears handles GET /Years
func (h *FilterHandler) GetYears(w http.ResponseWriter, r *http.Request) {
	years, err := h.repo.GetYears()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(years)
}

// GetFilters handles GET /Filters
func (h *FilterHandler) GetFilters(w http.ResponseWriter, r *http.Request) {
	filters := map[string]interface{}{
		"MediaTypes": []string{"Movie", "Series", "Episode", "MusicAlbum", "MusicArtist", "Book", "Photo"},
		"Genres":     []string{},
		"Studios":    []string{},
		"Years":      []int{},
		"OfficialRatings": []string{},
		"SortValues": []string{"SortName", "ProductionYear", "CommunityRating", "CriticRating", "DateCreated", "StartDate"},
		"SortOrders": []string{"Ascending", "Descending"},
		"Filters":    []string{"IsUnaired", "IsFavorite", "IsFavoriteOrLiked", "IsResumable", "IsPlayed", "IsUnplayed"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(filters)
}

// GetFilterGenres handles GET /Filters/Genres
func (h *FilterHandler) GetFilterGenres(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	genres, err := h.repo.GetGenres()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(genres)
}

// GetFilterStudios handles GET /Filters/Studios
func (h *FilterHandler) GetFilterStudios(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	studios, err := h.repo.GetStudios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studios)
}

// GetFilterYears handles GET /Filters/Years
func (h *FilterHandler) GetFilterYears(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	years, err := h.repo.GetYears()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(years)
}

// GetOfficialRatings handles GET /Filters/OfficialRatings
func (h *FilterHandler) GetOfficialRatings(w http.ResponseWriter, r *http.Request) {
	ratings := []string{
		"G", "PG", "PG-13", "R", "NC-17", "NR", "TV-Y", "TV-Y7", "TV-G", "TV-PG", "TV-14", "TV-MA",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ratings)
}

// GetNetworks handles GET /Filters/Networks
func (h *FilterHandler) GetNetworks(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	networks, err := h.repo.GetNetworks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(networks)
}

// GetTags handles GET /Filters/Tags
func (h *FilterHandler) GetTags(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("UserId")
	_ = userId

	tags, err := h.repo.GetTags()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tags)
}
