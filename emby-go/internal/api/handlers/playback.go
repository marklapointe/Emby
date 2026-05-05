package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/session"
	"github.com/go-chi/chi/v5"
)

// PlaybackHandler handles playback-related API endpoints.
type PlaybackHandler struct {
	sessionSvc *session.Manager
}

// NewPlaybackHandler creates a new playback handler.
func NewPlaybackHandler(sessionSvc *session.Manager) *PlaybackHandler {
	return &PlaybackHandler{
		sessionSvc: sessionSvc,
	}
}

// SelectPlayback handles POST /Playback/{type}/Selected
func (h *PlaybackHandler) SelectPlayback(w http.ResponseWriter, r *http.Request) {
	playType := chi.URLParam(r, "type")
	_ = playType

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "selected"})
}

// GetFormats handles GET /Playback/{type}/Formats
func (h *PlaybackHandler) GetFormats(w http.ResponseWriter, r *http.Request) {
	playType := chi.URLParam(r, "type")
	_ = playType

	formats := []map[string]interface{}{
		{"Name": "Direct", "Value": "direct"},
		{"Name": "Transcode", "Value": "transcode"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formats)
}

// GetTranscodeURL handles GET /Videos/{id}/stream
func (h *PlaybackHandler) GetTranscodeURL(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")
	_ = videoID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "transcode-not-implemented"})
}

// GetStream handles GET /Videos/{id}/stream with direct stream support.
func (h *PlaybackHandler) GetStream(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")
	_ = videoID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "stream-not-implemented"})
}

// GetSubtitleStream handles GET /Videos/{id}/Subtitles/{subtitleIndex}/Stream
func (h *PlaybackHandler) GetSubtitleStream(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")
	subtitleIndex := chi.URLParam(r, "subtitleIndex")
	_ = videoID
	_ = subtitleIndex

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(""))
}

// GetAudioStream handles GET /Videos/{id}/audio
func (h *PlaybackHandler) GetAudioStream(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")
	_ = videoID

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "audio-not-implemented"})
}

// PostPlaybackProgress handles POST /Sessions/{id}/Playing/Progress
func (h *PlaybackHandler) PostPlaybackProgress(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "id")

	var progress struct {
		PositionTicks   int64 `json:"PositionTicks"`
		IsPaused        bool  `json:"IsPaused"`
		PlayMethod      string `json:"PlayMethod"`
		SubtitleOffset  float64 `json:"SubtitleOffset"`
		AudioStreamIndex int  `json:"AudioStreamIndex"`
	}

	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update session playback progress
	// This would integrate with the session manager
	_ = sessionID
	_ = progress

	w.WriteHeader(http.StatusNoContent)
}

// PostPlaybackStopped handles POST /Sessions/{id}/Playing/Stopped
func (h *PlaybackHandler) PostPlaybackStopped(w http.ResponseWriter, r *http.Request) {
	sessionID := chi.URLParam(r, "id")

	// Stop playback for session
	// This would integrate with the session manager
	_ = sessionID

	w.WriteHeader(http.StatusNoContent)
}
