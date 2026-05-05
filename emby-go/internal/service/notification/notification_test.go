package notification

import (
	"testing"
)

func TestNewManager(t *testing.T) {
	m := NewManager()
	if m == nil {
		t.Fatal("NewManager returned nil")
	}
	if m.providers == nil {
		t.Error("providers map not initialized")
	}
	if m.notifications == nil {
		t.Error("notifications slice not initialized")
	}
}

func TestRegisterProvider(t *testing.T) {
	m := NewManager()

	provider := &Provider{
		ID:      "test-provider",
		Name:    "Test Provider",
		Enabled: true,
	}

	m.RegisterProvider(provider)

	providers := m.GetProviders()
	if len(providers) != 1 {
		t.Errorf("expected 1 provider, got %d", len(providers))
	}
}

func TestGetProviders_Empty(t *testing.T) {
	m := NewManager()
	providers := m.GetProviders()
	if len(providers) != 0 {
		t.Errorf("expected 0 providers, got %d", len(providers))
	}
}

func TestSendNotification(t *testing.T) {
	m := NewManager()

	notification := &Notification{
		Title:    "Test Title",
		Message:  "Test Message",
		UserID:   "user-1",
		Type:     "test",
		Severity: "info",
	}

	err := m.SendNotification(notification)
	if err != nil {
		t.Errorf("SendNotification returned error: %v", err)
	}

	if notification.ID == "" {
		t.Error("notification ID is empty after send")
	}
}

func TestGetNotifications(t *testing.T) {
	m := NewManager()

	m.SendNotification(&Notification{
		Title:   "Notification 1",
		UserID:  "user-1",
		Type:    "test",
	})

	m.SendNotification(&Notification{
		Title:   "Notification 2",
		UserID:  "user-2",
		Type:    "test",
	})

	notifications := m.GetNotifications("user-1")
	if len(notifications) != 1 {
		t.Errorf("expected 1 notification for user-1, got %d", len(notifications))
	}

	allNotifications := m.GetNotifications("")
	if len(allNotifications) != 2 {
		t.Errorf("expected 2 notifications total, got %d", len(allNotifications))
	}
}

func TestMarkAsRead(t *testing.T) {
	m := NewManager()

	m.SendNotification(&Notification{
		Title:   "Test",
		UserID:  "user-1",
		Type:    "test",
		IsRead:  false,
	})

	notification := m.notifications[0]
	err := m.MarkAsRead(notification.ID)
	if err != nil {
		t.Errorf("MarkAsRead returned error: %v", err)
	}

	if !m.notifications[0].IsRead {
		t.Error("notification isRead should be true after MarkAsRead")
	}
}

func TestMarkAsRead_NotFound(t *testing.T) {
	m := NewManager()
	err := m.MarkAsRead("non-existent")
	if err == nil {
		t.Error("expected error for non-existent notification")
	}
}

func TestMarkAllAsRead(t *testing.T) {
	m := NewManager()

	m.SendNotification(&Notification{UserID: "user-1", IsRead: false})
	m.SendNotification(&Notification{UserID: "user-1", IsRead: false})

	err := m.MarkAllAsRead("user-1")
	if err != nil {
		t.Errorf("MarkAllAsRead returned error: %v", err)
	}

	for _, n := range m.notifications {
		if !n.IsRead {
			t.Error("all notifications should be read after MarkAllAsRead")
		}
	}
}

func TestDeleteNotification(t *testing.T) {
	m := NewManager()

	m.SendNotification(&Notification{UserID: "user-1", Title: "Test"})

	notification := m.notifications[0]
	err := m.DeleteNotification(notification.ID)
	if err != nil {
		t.Errorf("DeleteNotification returned error: %v", err)
	}

	if len(m.notifications) != 0 {
		t.Errorf("expected 0 notifications, got %d", len(m.notifications))
	}
}

func TestDeleteNotification_NotFound(t *testing.T) {
	m := NewManager()
	err := m.DeleteNotification("non-existent")
	if err == nil {
		t.Error("expected error for non-existent notification")
	}
}

func TestGetUnreadCount(t *testing.T) {
	m := NewManager()

	m.SendNotification(&Notification{UserID: "user-1"})
	m.SendNotification(&Notification{UserID: "user-1"})
	m.SendNotification(&Notification{UserID: "user-1"})

	count := m.GetUnreadCount("user-1")
	if count != 3 {
		t.Errorf("expected 3 unread, got %d", count)
	}
}

func TestGetNotificationTypes(t *testing.T) {
	types := GetNotificationTypes()
	if len(types) == 0 {
		t.Error("expected non-empty notification types")
	}
}