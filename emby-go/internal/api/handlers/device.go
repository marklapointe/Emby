package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/device"
	"github.com/gorilla/mux"
)

// DeviceHandler handles device-related API endpoints.
type DeviceHandler struct {
	deviceMgr *device.Manager
}

// NewDeviceHandler creates a new device handler.
func NewDeviceHandler(deviceMgr *device.Manager) *DeviceHandler {
	return &DeviceHandler{deviceMgr: deviceMgr}
}

// GetDevices handles GET /Devices
func (h *DeviceHandler) GetDevices(w http.ResponseWriter, r *http.Request) {
	devices := h.deviceMgr.GetDevices()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(devices)
}

// GetDevice handles GET /Devices/{id}
func (h *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	device, exists := h.deviceMgr.GetDevice(id)
	if !exists {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(device)
}

// UpdateDevice handles PUT /Devices/{id}
func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		Name        string `json:"Name"`
		ProductName string `json:"AppName"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.deviceMgr.UpdateDevice(id, req.Name, req.ProductName); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteDevice handles DELETE /Devices/{id}
func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.deviceMgr.RemoveDevice(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetDeviceIcon handles GET /Devices/{id}/Icon
func (h *DeviceHandler) GetDeviceIcon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	device, exists := h.deviceMgr.GetDevice(id)
	if !exists {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	// Return device icon if available
	if device.DeviceProfile != nil && device.DeviceProfile.Name != "" {
		w.Header().Set("Content-Type", "image/png")
		w.Write([]byte("icon_data"))
		return
	}

	http.Error(w, "no icon available", http.StatusNotFound)
}

// GetDeviceProfile handles GET /Devices/{id}/Profile
func (h *DeviceHandler) GetDeviceProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	profile, exists := h.deviceMgr.GetDeviceProfile(id)
	if !exists {
		http.Error(w, "device not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
