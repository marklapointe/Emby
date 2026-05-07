package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/sync"
	"github.com/go-chi/chi/v5"
)

type SyncHandler struct {
	syncSvc *sync.Manager
}

func NewSyncHandler(syncSvc *sync.Manager) *SyncHandler {
	return &SyncHandler{syncSvc: syncSvc}
}

func (h *SyncHandler) GetJobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *SyncHandler) CreateJob(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *SyncHandler) GetJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id": id,
	})
}

func (h *SyncHandler) DeleteJob(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *SyncHandler) AddItemToJob(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}