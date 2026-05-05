package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/repository"
)

// SearchHandler handles search-related API endpoints.
type SearchHandler struct {
	repo *repository.ItemRepository
}

// NewSearchHandler creates a new search handler.
func NewSearchHandler(repo *repository.ItemRepository) *SearchHandler {
	return &SearchHandler{
		repo: repo,
	}
}

// GetHints handles GET /Search/Hints
func (h *SearchHandler) GetHints(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	items, err := h.repo.SearchItems(query, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	hints := map[string]interface{}{
		"SearchHints": items,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hints)
}

// SearchItems handles GET /Search/Items
func (h *SearchHandler) SearchItems(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	items, err := h.repo.SearchItems(query, 50, 0)
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

// SearchItemsByTerm handles GET /Items/Search
func (h *SearchHandler) SearchItemsByTerm(w http.ResponseWriter, r *http.Request) {
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
