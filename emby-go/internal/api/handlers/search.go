package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// SearchHandler handles search-related API endpoints.
type SearchHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewSearchHandler creates a new search handler.
func NewSearchHandler(cfg *config.Config, repo *repository.ItemRepository) *SearchHandler {
	return &SearchHandler{
		config: cfg,
		repo:   repo,
	}
}

// SearchItems handles GET /Items/Search
func (h *SearchHandler) SearchItems(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.URL.Query().Get("SearchTerm")
	mediaType := r.URL.Query().Get("MediaType")

	_ = searchTerm
	_ = mediaType

	results := []map[string]interface{}{
		{"Name": "Search Result", "Type": "Movie", "Id": "result-1"},
	}

	result := map[string]interface{}{
		"Items": results,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetSearchHints handles GET /Search/Hints
func (h *SearchHandler) GetSearchHints(w http.ResponseWriter, r *http.Request) {
	hints := []map[string]interface{}{
		{"Name": "Hint", "Id": "hint-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hints)
}

// GetArtistPhotos handles GET /Artists/{id}/Photos
func (h *SearchHandler) GetArtistPhotos(w http.ResponseWriter, r *http.Request) {
	photos := []map[string]interface{}{
		{"Name": "Artist Photo", "Id": "photo-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(photos)
}

// GetAlbumArtists handles GET /Music/AlbumArtists
func (h *SearchHandler) GetAlbumArtists(w http.ResponseWriter, r *http.Request) {
	artists := []map[string]interface{}{
		{"Name": "Album Artist", "Id": "artist-1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}
