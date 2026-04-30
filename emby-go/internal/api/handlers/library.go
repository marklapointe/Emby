package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emby/emby-go/internal/service/library"
	"github.com/gorilla/mux"
)

// LibraryHandler handles library-related API endpoints.
type LibraryHandler struct {
	scanner *library.Scanner
}

// NewLibraryHandler creates a new library handler.
func NewLibraryHandler(scanner *library.Scanner) *LibraryHandler {
	return &LibraryHandler{scanner: scanner}
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
	// Get query parameters
	ids := r.URL.Query().Get("Ids")
	mediaType := r.URL.Query().Get("MediaType")
	folderID := r.URL.Query().Get("FolderId")
	userId := r.URL.Query().Get("UserId")
	sortBy := r.URL.Query().Get("SortBy")
	sortOrder := r.URL.Query().Get("SortOrder")
	startIndex, _ := strconv.Atoi(r.URL.Query().Get("StartIndex"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	isFavorite := r.URL.Query().Get("IsFavorite")
	isUnaired := r.URL.Query().Get("IsUnaired")
	isNextUp := r.URL.Query().Get("IsNextUp")
	minPremiereDate := r.URL.Query().Get("MinPremiereDate")
	maxPremiereDate := r.URL.Query().Get("MaxPremiereDate")
	minDateCreated := r.URL.Query().Get("MinDateCreated")
	genre := r.URL.Query().Get("Genre")
	artist := r.URL.Query().Get("Artist")
	albumArtist := r.URL.Query().Get("AlbumArtist")
	albumID := r.URL.Query().Get("AlbumId")
	seasonID := r.URL.Query().Get("SeasonId")
	seriesID := r.URL.Query().Get("SeriesId")
	isMovie := r.URL.Query().Get("IsMovie")
	isSeries := r.URL.Query().Get("IsSeries")
	isNews := r.URL.Query().Get("IsNews")
	isKids := r.URL.Query().Get("IsKids")
	isSports := r.URL.Query().Get("IsSports")
	isBoxSet := r.URL.Query().Get("IsBoxSet")
	person := r.URL.Query().Get("Person")
	years := r.URL.Query().Get("Years")
	genreIDs := r.URL.Query().Get("GenreIds")
	officialRatings := r.URL.Query().Get("OfficialRatings")
	enableTotalCount := r.URL.Query().Get("EnableTotalItemCount")
	imageTypeLimit := r.URL.Query().Get("ImageTypeLimit")
	enableImageTypes := r.URL.Query().Get("EnableImageTypes")
	enableImages := r.URL.Query().Get("EnableImages")
	enableUserData := r.URL.Query().Get("EnableUserData")
	fields := r.URL.Query().Get("Fields")
	excludeImageTypes := r.URL.Query().Get("ExcludeImageTypes")
	excludeLocationTypes := r.URL.Query().Get("ExcludeLocationTypes")
	excludeIsFolder := r.URL.Query().Get("ExcludeIsFolder")
	recursive := r.URL.Query().Get("Recursive")
	seasonIDParam := r.URL.Query().Get("SeasonId")
	seriesIDParam := r.URL.Query().Get("SeriesId")
	videoTypes := r.URL.Query().Get("VideoTypes")
	mediaTypes := r.URL.Query().Get("MediaTypes")
	groups := r.URL.Query().Get("Groups")
	enableUserDataParam := r.URL.Query().Get("EnableUserData")
	enableImagesParam := r.URL.Query().Get("EnableImages")
	enableImageTypesParam := r.URL.Query().Get("EnableImageTypes")
	imageTypeLimitParam := r.URL.Query().Get("ImageTypeLimit")
	fieldsParam := r.URL.Query().Get("Fields")
	excludeImageTypesParam := r.URL.Query().Get("ExcludeImageTypes")
	excludeLocationTypesParam := r.URL.Query().Get("ExcludeLocationTypes")
	excludeIsFolderParam := r.URL.Query().Get("ExcludeIsFolder")
	recursiveParam := r.URL.Query().Get("Recursive")
	videoTypesParam := r.URL.Query().Get("VideoTypes")
	mediaTypesParam := r.URL.Query().Get("MediaTypes")
	groupsParam := r.URL.Query().Get("Groups")

	_ = ids
	_ = mediaType
	_ = folderID
	_ = userId
	_ = sortBy
	_ = sortOrder
	_ = startIndex
	_ = limit
	_ = isFavorite
	_ = isUnaired
	_ = isNextUp
	_ = minPremiereDate
	_ = maxPremiereDate
	_ = minDateCreated
	_ = genre
	_ = artist
	_ = albumArtist
	_ = albumID
	_ = seasonID
	_ = seriesID
	_ = isMovie
	_ = isSeries
	_ = isNews
	_ = isKids
	_ = isSports
	_ = isBoxSet
	_ = person
	_ = years
	_ = genreIDs
	_ = officialRatings
	_ = enableTotalCount
	_ = imageTypeLimit
	_ = enableImageTypes
	_ = enableImages
	_ = enableUserData
	_ = fields
	_ = excludeImageTypes
	_ = excludeLocationTypes
	_ = excludeIsFolder
	_ = recursive
	_ = seasonIDParam
	_ = seriesIDParam
	_ = videoTypes
	_ = mediaTypes
	_ = groups
	_ = enableUserDataParam
	_ = enableImagesParam
	_ = enableImageTypesParam
	_ = imageTypeLimitParam
	_ = fieldsParam
	_ = excludeImageTypesParam
	_ = excludeLocationTypesParam
	_ = excludeIsFolderParam
	_ = recursiveParam
	_ = videoTypesParam
	_ = mediaTypesParam
	_ = groupsParam

	// Return empty results for now
	result := map[string]interface{}{
		"Items":        []map[string]interface{}{},
		"TotalCount":   0,
		"StartIndex":   0,
		"HasMoreItems": false,
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
	vars := mux.Vars(r)
	id := vars["id"]

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
	vars := mux.Vars(r)
	id := vars["id"]

	// Remove path from scanner
	h.scanner.RemovePath("/media/" + id)

	w.WriteHeader(http.StatusNoContent)
}

// GetFolderItems handles GET /Library/MediaFolders/{id}/Items
func (h *LibraryHandler) GetFolderItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get query parameters
	mediaType := r.URL.Query().Get("MediaType")
	userId := r.URL.Query().Get("UserId")
	sortBy := r.URL.Query().Get("SortBy")
	sortOrder := r.URL.Query().Get("SortOrder")
	startIndex, _ := strconv.Atoi(r.URL.Query().Get("StartIndex"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("Limit"))
	isFavorite := r.URL.Query().Get("IsFavorite")
	enableTotalCount := r.URL.Query().Get("EnableTotalItemCount")
	enableImages := r.URL.Query().Get("EnableImages")
	enableUserData := r.URL.Query().Get("EnableUserData")
	fields := r.URL.Query().Get("Fields")
	recursive := r.URL.Query().Get("Recursive")
	mediaTypes := r.URL.Query().Get("MediaTypes")

	_ = mediaType
	_ = userId
	_ = sortBy
	_ = sortOrder
	_ = startIndex
	_ = limit
	_ = isFavorite
	_ = enableTotalCount
	_ = enableImages
	_ = enableUserData
	_ = id
	_ = fields
	_ = recursive
	_ = mediaTypes

	// Return empty results for now
	result := map[string]interface{}{
		"Items":        []map[string]interface{}{},
		"TotalCount":   0,
		"StartIndex":   0,
		"HasMoreItems": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// ScanLibrary handles POST /Library/Folders/FullScan
func (h *LibraryHandler) ScanLibrary(w http.ResponseWriter, r *http.Request) {
	// Trigger library scan
	_ = h.scanner

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
	vars := mux.Vars(r)
	id := vars["id"]

	_ = id

	result := map[string]interface{}{
		"Items":        []map[string]interface{}{},
		"TotalCount":   0,
		"StartIndex":   0,
		"HasMoreItems": false,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
