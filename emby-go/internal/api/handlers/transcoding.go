package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/transcoding"
)

// TranscodingHandler handles transcoding-related API endpoints.
type TranscodingHandler struct {
	transcodingSvc *transcoding.Manager
}

// NewTranscodingHandler creates a new transcoding handler.
func NewTranscodingHandler(transcodingSvc *transcoding.Manager) *TranscodingHandler {
	return &TranscodingHandler{transcodingSvc: transcodingSvc}
}

// GetTranscodingProfiles handles GET /TranscodingProfiles
func (h *TranscodingHandler) GetTranscodingProfiles(w http.ResponseWriter, r *http.Request) {
	profiles := h.transcodingSvc.GetTranscodingProfiles()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

// GetTranscodingProfile handles GET /TranscodingProfiles/{id}
func (h *TranscodingHandler) GetTranscodingProfile(w http.ResponseWriter, r *http.Request) {
	profile := map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// GetActiveTranscodes handles GET /ActiveTranscodes
func (h *TranscodingHandler) GetActiveTranscodes(w http.ResponseWriter, r *http.Request) {
	count := h.transcodingSvc.GetActiveStreamCount()
	transcodes := make([]map[string]interface{}, count)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transcodes)
}

// GetActiveTranscode handles GET /ActiveTranscodes/{id}
func (h *TranscodingHandler) GetActiveTranscode(w http.ResponseWriter, r *http.Request) {
	transcode := map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transcode)
}

// StopTranscode handles POST /ActiveTranscodes/{id}/Stop
func (h *TranscodingHandler) StopTranscode(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}