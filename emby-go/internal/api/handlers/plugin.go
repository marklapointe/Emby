package handlers

import (
	"encoding/json"
	"net/http"
)

type PluginHandler struct{}

func NewPluginHandler() *PluginHandler {
	return &PluginHandler{}
}

func (h *PluginHandler) GetPlugins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}

func (h *PluginHandler) DeletePlugin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *PluginHandler) GetPluginConfiguration(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{})
}

func (h *PluginHandler) GetSecurityInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"IsMBSupporter": true,
		"IsSuperUser":   true,
	})
}

func (h *PluginHandler) GetReleased(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode([]interface{}{})
}