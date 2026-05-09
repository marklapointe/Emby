package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RemoteSearchHandler struct{}

func NewRemoteSearchHandler() *RemoteSearchHandler {
	return &RemoteSearchHandler{}
}

func (h *RemoteSearchHandler) GetMovieRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetSeriesRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetTrailerRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetBookRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetGameRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetBoxSetRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetMusicVideoRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetPersonRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetMusicAlbumRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetMusicArtistRemoteSearchResults(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) GetRemoteSearchImage(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *RemoteSearchHandler) ApplySearchCriteria(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	_ = id

	w.WriteHeader(http.StatusNoContent)
}
