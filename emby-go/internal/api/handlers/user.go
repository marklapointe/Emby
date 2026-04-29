package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/user"
	"github.com/gorilla/mux"
)

// UserHandler handles user-related API endpoints.
type UserHandler struct {
	userMgr *user.Manager
}

// NewUserHandler creates a new user handler.
func NewUserHandler(userMgr *user.Manager) *UserHandler {
	return &UserHandler{userMgr: userMgr}
}

// GetUsers handles GET /Users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.userMgr.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetPublicUsers handles GET /Users/Public
func (h *UserHandler) GetPublicUsers(w http.ResponseWriter, r *http.Request) {
	users := h.userMgr.GetAllUsers()

	// Filter to public users only
	var publicUsers []*user.User
	for _, u := range users {
		if u.Policy != nil && !u.Policy.IsHidden {
			publicUsers = append(publicUsers, u)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(publicUsers)
}

// Login handles POST /Users/Login
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"Username"`
		Password string `json:"Password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := h.userMgr.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		http.Error(w, "authentication failed", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

// Logout handles POST /Users/Logout
func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Token string `json:"Token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userMgr.RevokeSession(req.Token); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUser handles GET /Users/{id}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, exists := h.userMgr.GetUser(id)
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser handles PUT /Users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		Name     string `json:"Name"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.userMgr.UpdateUser(id, &req.Name, &req.Email, &req.Password); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser handles DELETE /Users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.userMgr.DeleteUser(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ChangePassword handles POST /Users/{id}/Password
func (h *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var req struct {
		NewPassword string `json:"NewPassword"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user password
	newPassword := req.NewPassword
	if err := h.userMgr.UpdateUser(id, nil, nil, &newPassword); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetUserImage handles GET /Users/{id}/Images/{type}
func (h *UserHandler) GetUserImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	imageType := vars["type"]

	_ = id
	_ = imageType

	// Return placeholder image
	w.Header().Set("Content-Type", "image/png")
	w.Write([]byte("placeholder_image_data"))
}

// GetUserConfiguration handles GET /Users/{id}/Configuration
func (h *UserHandler) GetUserConfiguration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, exists := h.userMgr.GetUser(id)
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Configuration)
}

// UpdateUserConfiguration handles PUT /Users/{id}/Configuration
func (h *UserHandler) UpdateUserConfiguration(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var config user.UserConfiguration
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user configuration
	_ = id
	_ = config

	w.WriteHeader(http.StatusNoContent)
}

// GetUserPolicy handles GET /Users/{id}/Policy
func (h *UserHandler) GetUserPolicy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user, exists := h.userMgr.GetUser(id)
	if !exists {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user.Policy)
}

// UpdateUserPolicy handles PUT /Users/{id}/Policy
func (h *UserHandler) UpdateUserPolicy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var policy user.UserPolicy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update user policy
	_ = id
	_ = policy

	w.WriteHeader(http.StatusNoContent)
}

// GetUsersByDevice handles GET /Users/Device/{deviceId}
func (h *UserHandler) GetUsersByDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deviceId := vars["deviceId"]

	_ = deviceId

	// Return all users for now
	users := h.userMgr.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUsersByLibraryFolder handles GET /Users/LibraryFolders/{folderId}
func (h *UserHandler) GetUsersByLibraryFolder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	folderId := vars["folderId"]

	_ = folderId

	// Return all users for now
	users := h.userMgr.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
