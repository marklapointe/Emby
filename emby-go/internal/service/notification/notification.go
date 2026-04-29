package notification

import (
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
)

// Notification represents a notification message.
type Notification struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	Type      string    `json:"type"` // info, warning, error, success
	UserID    string    `json:"userId,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	IsRead    bool      `json:"isRead"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

// Manager handles notification operations.
type Manager struct {
	mu         sync.RWMutex
	notifications map[string][]*Notification
	logger     *zap.Logger
}

// NewManager creates a new notification manager.
func NewManager(logger *zap.Logger) *Manager {
	return &Manager{
		notifications: make(map[string][]*Notification),
		logger:        logger,
	}
}

// SendNotification sends a notification to a user.
func (m *Manager) SendNotification(userID, name, content, notificationType string, data map[string]interface{}) (*Notification, error) {
	notification := &Notification{
		ID:        fmt.Sprintf("notif-%d", time.Now().UnixNano()),
		Name:      name,
		Content:   content,
		Type:      notificationType,
		UserID:    userID,
		CreatedAt: time.Now(),
		IsRead:    false,
		Data:      data,
	}

	m.mu.Lock()
	m.notifications[userID] = append(m.notifications[userID], notification)
	m.mu.Unlock()

	m.logger.Info("notification sent",
		zap.String("userId", userID),
		zap.String("name", name),
	)

	return notification, nil
}

// GetNotifications returns notifications for a user.
func (m *Manager) GetNotifications(userID string, limit, offset int) ([]*Notification, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	notifs := m.notifications[userID]
	if notifs == nil {
		return []*Notification{}, nil
	}

	// Apply pagination
	if offset >= len(notifs) {
		return []*Notification{}, nil
	}

	end := offset + limit
	if end > len(notifs) {
		end = len(notifs)
	}

	return notifs[offset:end], nil
}

// MarkAsRead marks a notification as read.
func (m *Manager) MarkAsRead(userID, notificationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	notifs := m.notifications[userID]
	for i, n := range notifs {
		if n.ID == notificationID {
			notifs[i].IsRead = true
			return nil
		}
	}

	return fmt.Errorf("notification not found: %s", notificationID)
}

// MarkAllAsRead marks all notifications for a user as read.
func (m *Manager) MarkAllAsRead(userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	notifs := m.notifications[userID]
	for i := range notifs {
		notifs[i].IsRead = true
	}

	return nil
}

// GetUnreadCount returns the number of unread notifications for a user.
func (m *Manager) GetUnreadCount(userID string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, n := range m.notifications[userID] {
		if !n.IsRead {
			count++
		}
	}
	return count
}

// BroadcastNotification sends a notification to all users.
func (m *Manager) BroadcastNotification(name, content, notificationType string, data map[string]interface{}) {
	m.mu.RLock()
	userIDs := make([]string, 0, len(m.notifications))
	for userID := range m.notifications {
		userIDs = append(userIDs, userID)
	}
	m.mu.RUnlock()

	for _, userID := range userIDs {
		m.SendNotification(userID, name, content, notificationType, data)
	}
}

// DeleteNotification deletes a notification.
func (m *Manager) DeleteNotification(userID, notificationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	notifs := m.notifications[userID]
	for i, n := range notifs {
		if n.ID == notificationID {
			m.notifications[userID] = append(notifs[:i], notifs[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("notification not found: %s", notificationID)
}

// ClearNotifications clears all notifications for a user.
func (m *Manager) ClearNotifications(userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.notifications, userID)
	return nil
}

// NotificationType constants
const (
	NotificationTypeInformation = "info"
	NotificationTypeWarning     = "warning"
	NotificationTypeError       = "error"
	NotificationTypeSuccess     = "success"
)
