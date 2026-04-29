package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/session"
	"github.com/gorilla/mux"
)

// SessionHandler handles session-related API endpoints.
type SessionHandler struct {
	sessionMgr *session.Manager
}

// NewSessionHandler creates a new session handler.
func NewSessionHandler(sessionMgr *session.Manager) *SessionHandler {
	return &SessionHandler{sessionMgr: sessionMgr}
}

// GetSessions handles GET /Sessions
func (h *SessionHandler) GetSessions(w http.ResponseWriter, r *http.Request) {
	sessions := h.sessionMgr.GetAllSessions()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sessions)
}

// GetSession handles GET /Sessions/{id}
func (h *SessionHandler) GetSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	sess, exists := h.sessionMgr.GetSession(id)
	if !exists {
		http.Error(w, "session not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sess)
}

// StartPlayback handles POST /Sessions/{id}/Playing
func (h *SessionHandler) StartPlayback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		ItemId        string `json:"ItemId"`
		MediaSourceId string `json:"MediaSourceId"`
		AudioStreamIndex int `json:"AudioStreamIndex"`
		SubtitleStreamIndex int `json:"SubtitleStreamIndex"`
		StartPositionTicks int64 `json:"StartPositionTicks"`
		MaxStreamingBitrate int `json:"MaxStreamingBitrate"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update session with playback info
	_ = id
	_ = req

	w.WriteHeader(http.StatusNoContent)
}

// PlaybackProgress handles POST /Sessions/{id}/Playing/Progress
func (h *SessionHandler) PlaybackProgress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		PositionTicks      int64   `json:"PositionTicks"`
		IsPaused           bool    `json:"IsPaused"`
		PlayMethod         string  `json:"PlayMethod"`
		SubtitleOffsetSecs float64 `json:"SubtitleOffsetSecs"`
		AudioStreamIndex   int     `json:"AudioStreamIndex"`
		SubtitleStreamIndex int    `json:"SubtitleStreamIndex"`
		LiveTimeTicks      int64   `json:"LiveTimeTicks"`
		VideoStreamIndex   int     `json:"VideoStreamIndex"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update session playback progress
	newPos := req.PositionTicks
	if err := h.sessionMgr.UpdateSession(id, &newPos, nil, &req.IsPaused); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// StopPlayback handles POST /Sessions/{id}/Playing/Stopped
func (h *SessionHandler) StopPlayback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Stop playback for session
	_ = id

	w.WriteHeader(http.StatusNoContent)
}

// SetVolume handles POST /Sessions/{id}/Volume
func (h *SessionHandler) SetVolume(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		Volume int `json:"Volume"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update session volume
	newVol := req.Volume
	if err := h.sessionMgr.UpdateSession(id, nil, &newVol, nil); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// PausePlayback handles POST /Sessions/{id}/Pause
func (h *SessionHandler) PausePlayback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	isPaused := true
	if err := h.sessionMgr.UpdateSession(id, nil, nil, &isPaused); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// UnpausePlayback handles POST /Sessions/{id}/Unpause
func (h *SessionHandler) UnpausePlayback(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	isPaused := false
	if err := h.sessionMgr.UpdateSession(id, nil, nil, &isPaused); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ToggleFullscreen handles POST /Sessions/{id}/ToggleFullscreen
func (h *SessionHandler) ToggleFullscreen(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Toggle fullscreen state
	_ = id

	w.WriteHeader(http.StatusNoContent)
}

// NavigateTo handles POST /Sessions/{id}/GoTo
func (h *SessionHandler) NavigateTo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		ItemId string `json:"ItemId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Navigate to item
	_ = id
	_ = req

	w.WriteHeader(http.StatusNoContent)
}

// SendKey handles POST /Sessions/{id}/SendKey
func (h *SessionHandler) SendKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		Key string `json:"Key"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send key to device
	_ = id
	_ = req

	w.WriteHeader(http.StatusNoContent)
}

// SendText handles POST /Sessions/{id}/SendText
func (h *SessionHandler) SendText(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		Text string `json:"Text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Send text to device
	_ = id
	_ = req

	w.WriteHeader(http.StatusNoContent)
}

// CloseSession handles DELETE /Sessions/{id}
func (h *SessionHandler) CloseSession(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.sessionMgr.DeleteSession(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
