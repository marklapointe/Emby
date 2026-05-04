package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/emby/emby-go/internal/service/transcoding"
	"github.com/go-chi/chi/v5"
)

// PlaybackHandler handles playback-related API endpoints.
type PlaybackHandler struct {
	transcoder *transcoding.Manager
}

// NewPlaybackHandler creates a new playback handler.
func NewPlaybackHandler(transcoder *transcoding.Manager) *PlaybackHandler {
	return &PlaybackHandler{transcoder: transcoder}
}

// GetTranscodeURL handles GET /Videos/{id}/stream
func (h *PlaybackHandler) GetTranscodeURL(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")

	profile := r.URL.Query().Get("profile")
	if profile == "" {
		profile = "default"
	}

	streamInfo, err := h.transcoder.GetStreamURL(videoID, profile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(streamInfo)
}

// GetStream handles GET /Videos/{id}/stream with direct stream support.
func (h *PlaybackHandler) GetStream(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")

	// Get stream parameters from query
	mediaSourceID := r.URL.Query().Get("MediaSourceId")
	audioCodec := r.URL.Query().Get("AudioCodec")
	videoCodec := r.URL.Query().Get("VideoCodec")
	maxVideoBitrate := r.URL.Query().Get("MaxVideoBitrate")
	maxAudioBitrate := r.URL.Query().Get("MaxAudioBitrate")
	container := r.URL.Query().Get("container")
	streamType := r.URL.Query().Get("StreamType")

	// Build transcoding command
	cmd, err := h.transcoder.BuildTranscodeCommand(videoID, mediaSourceID, transcoding.TranscodeConfig{
		AudioCodec:    audioCodec,
		VideoCodec:    videoCodec,
		MaxVideoBitrate: maxVideoBitrate,
		MaxAudioBitrate: maxAudioBitrate,
		Container:     container,
		StreamType:    streamType,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute transcoding
	output, err := h.transcoder.ExecuteTranscode(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Stream output to client
	w.Header().Set("Content-Type", "video/mp2t")
	w.Header().Set("Transfer-Encoding", "chunked")
	io.Copy(w, output)
}

// GetSubtitleStream handles GET /Videos/{id}/Subtitles/{subtitleIndex}/Stream
func (h *PlaybackHandler) GetSubtitleStream(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")
	subtitleIndex := chi.URLParam(r, "subtitleIndex")

	// Get subtitle format from query
	format := r.URL.Query().Get("format")
	if format == "" {
		format = "vtt"
	}

	// Get subtitle stream
	stream, err := h.transcoder.GetSubtitleStream(videoID, subtitleIndex, format)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write(stream)
}

// GetAudioStream handles GET /Videos/{id}/audio
func (h *PlaybackHandler) GetAudioStream(w http.ResponseWriter, r *http.Request) {
	videoID := chi.URLParam(r, "id")

	// Get audio stream parameters
	mediaSourceID := r.URL.Query().Get("MediaSourceId")
	audioCodec := r.URL.Query().Get("AudioCodec")
	maxAudioBitrate := r.URL.Query().Get("MaxAudioBitrate")

	// Build audio transcoding command
	cmd, err := h.transcoder.BuildAudioTranscodeCommand(videoID, mediaSourceID, transcoding.AudioTranscodeConfig{
		AudioCodec:    audioCodec,
		MaxAudioBitrate: maxAudioBitrate,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute transcoding
	output, err := h.transcoder.ExecuteAudioTranscode(cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Stream output to client
	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Transfer-Encoding", "chunked")
	io.Copy(w, output)
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
