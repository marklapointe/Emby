package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/service/notification"
)

// NotificationHandler handles notification-related API endpoints.
type NotificationHandler struct {
	notificationSvc *notification.Manager
}

// NewNotificationHandler creates a new notification handler.
func NewNotificationHandler(notificationSvc *notification.Manager) *NotificationHandler {
	return &NotificationHandler{notificationSvc: notificationSvc}
}

// GetNotifications handles GET /Notifications
func (h *NotificationHandler) GetNotifications(w http.ResponseWriter, r *http.Request) {
	notifications := []map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// GetUnreadNotifications handles GET /Notifications/Unread
func (h *NotificationHandler) GetUnreadNotifications(w http.ResponseWriter, r *http.Request) {
	notifications := []map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notifications)
}

// MarkAsRead handles POST /Notifications/{id}/MarkRead
func (h *NotificationHandler) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// MarkAllAsRead handles POST /Notifications/MarkAllRead
func (h *NotificationHandler) MarkAllAsRead(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// DeleteNotification handles DELETE /Notifications/{id}
func (h *NotificationHandler) DeleteNotification(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

// GetNotificationCount handles GET /Notifications/Count
func (h *NotificationHandler) GetNotificationCount(w http.ResponseWriter, r *http.Request) {
	count := map[string]int{"count": 0}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}

// GetUnreadNotificationCount handles GET /Notifications/UnreadCount
func (h *NotificationHandler) GetUnreadNotificationCount(w http.ResponseWriter, r *http.Request) {
	count := map[string]int{"count": 0}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(count)
}