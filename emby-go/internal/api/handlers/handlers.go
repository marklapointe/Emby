package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/media"
	"github.com/emby/emby-go/internal/service/session"
	"github.com/emby/emby-go/internal/service/user"
	"github.com/go-chi/chi/v5"
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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Name": "Media Library",
		"Id":   "root",
		"MediaType": "folder",
	})
}

// GetItems handles GET /Library/Items
func (h *LibraryHandler) GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Items":      []interface{}{},
		"TotalCount": 0,
		"StartIndex": 0,
	})
}

// ScanLibrary handles POST /Library/Root/Scan
func (h *LibraryHandler) ScanLibrary(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := h.scanner.ScanLibrary(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

// MediaHandler handles media-related API endpoints.
type MediaHandler struct {
	mediaSvc *media.Manager
}

// NewMediaHandler creates a new media handler.
func NewMediaHandler(mediaSvc *media.Manager) *MediaHandler {
	return &MediaHandler{mediaSvc: mediaSvc}
}

// GetMediaInfo handles GET /Items/{id}/MediaStreams
func (h *MediaHandler) GetMediaInfo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id":   id,
		"Name": "Media Info",
	})
}

// GetPlaybackInfo handles GET /Items/{id}/PlaybackInfo
func (h *MediaHandler) GetPlaybackInfo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id":       id,
		"MediaSources": []interface{}{},
	})
}

// SessionHandler handles session-related API endpoints.
type SessionHandler struct {
	sessionSvc *session.Manager
}

// NewSessionHandler creates a new session handler.
func NewSessionHandler(sessionSvc *session.Manager) *SessionHandler {
	return &SessionHandler{sessionSvc: sessionSvc}
}

// GetSessions handles GET /Sessions
func (h *SessionHandler) GetSessions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Sessions": []interface{}{},
	})
}

// SendCommand handles POST /Sessions/{id}/Playing/{command}
func (h *SessionHandler) SendCommand(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	command := chi.URLParam(r, "command")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"SessionId": id,
		"Command":   command,
	})
}

// UserHandler handles user-related API endpoints.
type UserHandler struct {
	userSvc *user.Manager
}

// NewUserHandler creates a new user handler.
func NewUserHandler(userSvc *user.Manager) *UserHandler {
	return &UserHandler{userSvc: userSvc}
}

// GetUsers handles GET /Users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Users": []interface{}{},
	})
}

// GetUser handles GET /Users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id":   id,
		"Name": "User",
	})
}

// PostUser handles POST /Users/Public/Login
func (h *UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var req map[string]string
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id":            "user-1",
		"AccessToken":   "token-123",
		"Name":          req["username"],
		"ServerId":      "server-1",
		"ExpirationDate": "2025-12-31T00:00:00Z",
	})
}
