package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/dlna"
	"github.com/go-chi/chi/v5"
)

type DLNAHandler struct {
	dlnaSvc *dlna.Manager
}

func NewDLNAHandler(dlnaSvc *dlna.Manager) *DLNAHandler {
	return &DLNAHandler{dlnaSvc: dlnaSvc}
}

func (h *DLNAHandler) GetProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *DLNAHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"Id": id,
	})
}

func (h *DLNAHandler) GetProfileInfos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *DLNAHandler) GetDefaultProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}