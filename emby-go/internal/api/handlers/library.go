package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/go-chi/chi/v5"
)

// LibraryHandler handles library-related API endpoints.
type LibraryHandler struct {
	scanner *library.Scanner
	repo    *repository.ItemRepository
}

// NewLibraryHandler creates a new library handler.
func NewLibraryHandler(scanner *library.Scanner, repo *repository.ItemRepository) *LibraryHandler {
	return &LibraryHandler{scanner: scanner, repo: repo}
}

// GetLibraryRoot handles GET /Library/Root
func (h *LibraryHandler) GetLibraryRoot(w http.ResponseWriter, r *http.Request) {
	root := map[string]interface{}{
		"Name":        "Media Library",
		"Type":        "Folder",
		"ItemId":      "",
		"Path":        "/media",
		"Children":    []map[string]interface{}{
			{"Name": "Movies", "Type": "Folder", "ItemId": "movies"},
			{"Name": "TV Shows", "Type": "Folder", "ItemId": "tvshows"},
			{"Name": "Music", "Type": "Folder", "ItemId": "music"},
			{"Name": "Photos", "Type": "Folder", "ItemId": "photos"},
			{"Name": "Books", "Type": "Folder", "ItemId": "books"},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(root)
}

// GetItems handles GET /Library/Items
func (h *LibraryHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	ids := r.URL.Query().Get("Ids")
	mediaType := r.URL.Query().Get("MediaType")
	folderID := r.URL.Query().Get("FolderId")
	userId := r.URL.Query().Get("UserId")
	sortBy := r.URL.Query().Get("SortBy")
	startIndex, _ := strconv.Atoi(r.URL.Query().Get("StartIndex"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	if limit <= 0 || limit > 100 {
		limit = 50
	}

	var items []map[string]interface{}
	var totalCount int

	if ids != "" {
		item, err := h.repo.GetItem(ids)
		if err == nil && item != nil {
			items = []map[string]interface{}{item}
			totalCount = 1
		}
	} else if h.repo != nil {
		if searchQuery := r.URL.Query().Get("SearchTerm"); searchQuery != "" {
			var err error
			items, err = h.repo.SearchItems(searchQuery, limit, startIndex)
			if err != nil {
				items = []map[string]interface{}{}
			}
			totalCount = len(items)
		} else {
			counts, err := h.repo.GetTotalItemCounts()
			if err == nil {
				for _, c := range counts {
					totalCount += c
				}
			}
			items = []map[string]interface{}{}
		}
	} else {
		items = []map[string]interface{}{}
	}

	_ = mediaType
	_ = folderID
	_ = userId
	_ = sortBy

	result := map[string]interface{}{
		"Items":        items,
		"TotalCount":   totalCount,
		"StartIndex":   startIndex,
		"HasMoreItems": startIndex+len(items) < totalCount,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetMediaFolders handles GET /Library/MediaFolders
func (h *LibraryHandler) GetMediaFolders(w http.ResponseWriter, r *http.Request) {
	folders := []map[string]interface{}{
		{"Name": "Movies", "Path": "/media/movies", "Id": "movies"},
		{"Name": "TV Shows", "Path": "/media/tvshows", "Id": "tvshows"},
		{"Name": "Music", "Path": "/media/music", "Id": "music"},
		{"Name": "Photos", "Path": "/media/photos", "Id": "photos"},
		{"Name": "Books", "Path": "/media/books", "Id": "books"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

// CreateMediaFolder handles POST /Library/MediaFolders
func (h *LibraryHandler) CreateMediaFolder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name   string `json:"Name"`
		Path   string `json:"Path"`
		IsHidden bool `json:"IsHidden"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Add path to scanner
	h.scanner.AddPath(req.Path)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Name":   req.Name,
		"Path":   req.Path,
		"IsHidden": req.IsHidden,
	})
}

// GetMediaFolder handles GET /Library/MediaFolders/{id}
func (h *LibraryHandler) GetMediaFolder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	folder := map[string]interface{}{
		"Name":   id,
		"Path":   "/media/" + id,
		"Id":     id,
		"IsHidden": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folder)
}

// DeleteMediaFolder handles DELETE /Library/MediaFolders/{id}
func (h *LibraryHandler) DeleteMediaFolder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Remove path from scanner
	h.scanner.RemovePath("/media/" + id)

	w.WriteHeader(http.StatusNoContent)
}

// GetFolderItems handles GET /Library/MediaFolders/{id}/Items
func (h *LibraryHandler) GetFolderItems(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	startIndex, _ := strconv.Atoi(r.URL.Query().Get("StartIndex"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	if limit <= 0 || limit > 100 {
		limit = 50
	}
	mediaType := r.URL.Query().Get("MediaType")

	var items []map[string]interface{}
	if h.repo != nil {
		var err error
		items, err = h.repo.GetItemsByParent(id, mediaType, limit, startIndex)
		if err != nil {
			items = []map[string]interface{}{}
		}
	} else {
		items = []map[string]interface{}{}
	}

	result := map[string]interface{}{
		"Items":        items,
		"TotalCount":   len(items),
		"StartIndex":   startIndex,
		"HasMoreItems": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// ScanLibrary handles POST /Library/Folders/FullScan
func (h *LibraryHandler) ScanLibrary(w http.ResponseWriter, r *http.Request) {
	if h.scanner == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	go func() {
		_, _ = h.scanner.ScanLibrary(r.Context())
	}()

	w.WriteHeader(http.StatusNoContent)
}

// GetVirtualFolder handles GET /Library/VirtualFolders
func (h *LibraryHandler) GetVirtualFolders(w http.ResponseWriter, r *http.Request) {
	folders := []map[string]interface{}{
		{"Name": "Movies", "CollectionType": "movies", "Id": "movies"},
		{"Name": "TV Shows", "CollectionType": "tvshows", "Id": "tvshows"},
		{"Name": "Music", "CollectionType": "music", "Id": "music"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(folders)
}

// GetVirtualFolderItems handles GET /Library/VirtualFolders/{id}/Items
func (h *LibraryHandler) GetVirtualFolderItems(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	startIndex, _ := strconv.Atoi(r.URL.Query().Get("StartIndex"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	if limit <= 0 || limit > 100 {
		limit = 50
	}

	var items []map[string]interface{}
	if h.repo != nil {
		var err error
		items, err = h.repo.GetItemsByParent(id, "", limit, startIndex)
		if err != nil {
			items = []map[string]interface{}{}
		}
	} else {
		items = []map[string]interface{}{}
	}

	result := map[string]interface{}{
		"Items":        items,
		"TotalCount":   len(items),
		"StartIndex":   startIndex,
		"HasMoreItems": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
