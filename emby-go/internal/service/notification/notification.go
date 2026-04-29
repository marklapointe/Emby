package notification

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Notification represents a server notification.
type Notification struct {
	ID          string    `json:"Id"`
	Name        string    `json:"Name"`
	Text        string    `json:"Text"`
	HTML        string    `json:"Html,omitempty"`
	ImageURL    string    `json:"ImageUrl,omitempty"`
	DateTime    time.Time `json:"DateTime"`
	UserID      string    `json:"UserId,omitempty"`
	IsRead      bool      `json:"IsRead"`
	Severity     string    `json:"Severity"`
	Type         string    `json:"Type"`
}

// Manager handles notification operations.
type Manager struct {
	mu          sync.RWMutex
	notifications map[string]*Notification
	logger      *zap.Logger
}

// NewManager creates a new notification manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		notifications: make(map[string]*Notification),
		logger:        logger,
	}
}

// SendNotification sends a notification to a user.
func (m *Manager) SendNotification(name, text, html, imageURL, userID, severity, notificationType string) (*Notification, error) {
	notification := &Notification{
		ID:         fmt.Sprintf("notif-%d", time.Now().UnixNano()),
		Name:       name,
		Text:       text,
		HTML:       html,
		ImageURL:   imageURL,
		DateTime:   time.Now(),
		UserID:     userID,
		IsRead:     false,
		Severity:   severity,
		Type:       notificationType,
	}

	m.mu.Lock()
	m.notifications[notification.ID] = notification
	m.mu.Unlock()

	if m.logger != nil {
		m.logger.Info("notification sent", zap.String("id", notification.ID), zap.String("name", name))
	}

	return notification, nil
}

// GetNotifications returns all notifications.
func (m *Manager) GetNotifications(userID string) []*Notification {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var notifications []*Notification
	for _, n := range m.notifications {
		if userID == "" || n.UserID == userID {
			notifications = append(notifications, n)
		}
	}
	return notifications
}

// GetUnreadNotifications returns unread notifications for a user.
func (m *Manager) GetUnreadNotifications(userID string) []*Notification {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var notifications []*Notification
	for _, n := range m.notifications {
		if (userID == "" || n.UserID == userID) && !n.IsRead {
			notifications = append(notifications, n)
		}
	}
	return notifications
}

// MarkAsRead marks a notification as read.
func (m *Manager) MarkAsRead(notificationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	n, exists := m.notifications[notificationID]
	if !exists {
		return fmt.Errorf("notification not found: %s", notificationID)
	}

	n.IsRead = true
	return nil
}

// MarkAllAsRead marks all notifications as read.
func (m *Manager) MarkAllAsRead(userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, n := range m.notifications {
		if userID == "" || n.UserID == userID {
			n.IsRead = true
		}
	}
	return nil
}

// DeleteNotification deletes a notification.
func (m *Manager) DeleteNotification(notificationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.notifications[notificationID]; !exists {
		return fmt.Errorf("notification not found: %s", notificationID)
	}

	delete(m.notifications, notificationID)
	return nil
}

// GetNotificationCount returns the number of notifications.
func (m *Manager) GetNotificationCount(userID string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, n := range m.notifications {
		if userID == "" || n.UserID == userID {
			count++
		}
	}
	return count
}

// GetUnreadNotificationCount returns the number of unread notifications.
func (m *Manager) GetUnreadNotificationCount(userID string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, n := range m.notifications {
		if (userID == "" || n.UserID == userID) && !n.IsRead {
			count++
		}
	}
	return count
}
