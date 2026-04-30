package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/media"
	"github.com/gorilla/mux"
)

// MediaHandler handles media-related API endpoints.
type MediaHandler struct {
	mediaMgr *media.Manager
}

// NewMediaHandler creates a new media handler.
func NewMediaHandler(mediaMgr *media.Manager) *MediaHandler {
	return &MediaHandler{mediaMgr: mediaMgr}
}

// GetItem handles GET /Items/{id}
func (h *MediaHandler) GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	mediaInfo, err := h.mediaMgr.GetMediaInfo(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mediaInfo)
}

// GetStream handles GET /Items/{id}/Stream
func (h *MediaHandler) GetStream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	mediaSource, err := h.mediaMgr.GetMediaSource(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Content-Disposition", "inline; filename="+mediaSource.Name)
	w.Write([]byte("stream_data"))
}

// GetSubtitles handles GET /Items/{id}/Subtitles
func (h *MediaHandler) GetSubtitles(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	mediaInfo, err := h.mediaMgr.GetMediaInfo(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mediaInfo.SubtitleStreams)
}

// GetSubtitleStream handles GET /Items/{id}/Subtitles/{index}/Stream
func (h *MediaHandler) GetSubtitleStream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]
	index := vars["index"]

	_ = itemID
	_ = index

	w.Header().Set("Content-Type", "text/vtt")
	w.Write([]byte("#VTT\n\n00:00:00.000 --> 00:00:05.000\nSubtitle text"))
}

// GetAudioStream handles GET /Items/{id}/Audio
func (h *MediaHandler) GetAudioStream(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemID := vars["id"]

	mediaSource, err := h.mediaMgr.GetMediaSource(itemID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "inline; filename="+mediaSource.Name+".mp3")
	w.Write([]byte("audio_data"))
}
