package notification

import (
	"fmt"
	"sync"
	"time"
)

// Provider represents a notification provider.
type Provider struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Enabled     bool   `json:"enabled"`
	Config      map[string]string `json:"config"`
}

// Notification represents a notification message.
type Notification struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Type      string    `json:"type"`
	Icon      string    `json:"icon"`
	Severity  string    `json:"severity"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	IsRead    bool      `json:"isRead"`
}

// Manager handles notification sending and management.
type Manager struct {
	providers   map[string]*Provider
	notifications []*Notification
	mu          sync.RWMutex
}

// NewManager creates a new notification manager.
func NewManager() *Manager {
	return &Manager{
		providers:   make(map[string]*Provider),
		notifications: make([]*Notification, 0),
	}
}

// RegisterProvider registers a notification provider.
func (m *Manager) RegisterProvider(provider *Provider) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.providers[provider.ID] = provider
}

// GetProviders returns all registered providers.
func (m *Manager) GetProviders() []*Provider {
	m.mu.RLock()
	defer m.mu.RUnlock()

	providers := make([]*Provider, 0, len(m.providers))
	for _, p := range m.providers {
		providers = append(providers, p)
	}

	return providers
}

// SendNotification sends a notification to all enabled providers.
func (m *Manager) SendNotification(notification *Notification) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	notification.ID = fmt.Sprintf("notif-%d", time.Now().UnixNano())
	notification.CreatedAt = time.Now()
	notification.IsRead = false

	m.notifications = append(m.notifications, notification)

	// Send to all enabled providers
	for _, provider := range m.providers {
		if provider.Enabled {
			if err := m.sendToProvider(provider, notification); err != nil {
				return fmt.Errorf("failed to send to provider %s: %w", provider.Name, err)
			}
		}
	}

	return nil
}

// sendToProvider sends a notification to a specific provider.
func (m *Manager) sendToProvider(provider *Provider, notification *Notification) error {
	// For now, just log the notification
	_ = provider
	_ = notification
	return nil
}

// GetNotifications returns all notifications for a user.
func (m *Manager) GetNotifications(userID string) []*Notification {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var notifications []*Notification
	for _, n := range m.notifications {
		if n.UserID == userID || userID == "" {
			notifications = append(notifications, n)
		}
	}

	return notifications
}

// MarkAsRead marks a notification as read.
func (m *Manager) MarkAsRead(notificationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, n := range m.notifications {
		if n.ID == notificationID {
			n.IsRead = true
			return nil
		}
	}

	return fmt.Errorf("notification not found: %s", notificationID)
}

// MarkAllAsRead marks all notifications as read.
func (m *Manager) MarkAllAsRead(userID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, n := range m.notifications {
		if n.UserID == userID || userID == "" {
			n.IsRead = true
		}
	}

	return nil
}

// DeleteNotification deletes a notification.
func (m *Manager) DeleteNotification(notificationID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for i, n := range m.notifications {
		if n.ID == notificationID {
			m.notifications = append(m.notifications[:i], m.notifications[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("notification not found: %s", notificationID)
}

// GetUnreadCount returns the count of unread notifications.
func (m *Manager) GetUnreadCount(userID string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, n := range m.notifications {
		if (n.UserID == userID || userID == "") && !n.IsRead {
			count++
		}
	}

	return count
}

// GetNotificationTypes returns all available notification types.
func GetNotificationTypes() []map[string]interface{} {
	return []map[string]interface{}{
		{"Name": "LibraryScan", "Icon": "cloud_upload", "Severity": "info"},
		{"Name": "UserLogin", "Icon": "login", "Severity": "info"},
		{"Name": "UserLogout", "Icon": "logout", "Severity": "info"},
		{"Name": "MediaAdded", "Icon": "add_circle", "Severity": "info"},
		{"Name": "MediaRemoved", "Icon": "remove_circle", "Severity": "warning"},
		{"Name": "MediaUpdated", "Icon": "update", "Severity": "info"},
		{"Name": "MediaPlayed", "Icon": "play_circle", "Severity": "info"},
		{"Name": "MediaPaused", "Icon": "pause_circle", "Severity": "info"},
		{"Name": "MediaStopped", "Icon": "stop_circle", "Severity": "info"},
		{"Name": "MediaFavorite", "Icon": "favorite", "Severity": "info"},
		{"Name": "MediaUnfavorite", "Icon": "favorite_border", "Severity": "info"},
		{"Name": "MediaRating", "Icon": "star", "Severity": "info"},
		{"Name": "MediaUnrating", "Icon": "star_border", "Severity": "info"},
		{"Name": "MediaDownload", "Icon": "download", "Severity": "info"},
		{"Name": "MediaUpload", "Icon": "upload", "Severity": "info"},
		{"Name": "MediaShare", "Icon": "share", "Severity": "info"},
		{"Name": "MediaComment", "Icon": "comment", "Severity": "info"},
		{"Name": "MediaLike", "Icon": "thumb_up", "Severity": "info"},
		{"Name": "MediaDislike", "Icon": "thumb_down", "Severity": "info"},
		{"Name": "MediaWatch", "Icon": "visibility", "Severity": "info"},
		{"Name": "MediaUnwatch", "Icon": "visibility_off", "Severity": "info"},
		{"Name": "MediaBookmark", "Icon": "bookmark", "Severity": "info"},
		{"Name": "MediaUnbookmark", "Icon": "bookmark_border", "Severity": "info"},
		{"Name": "MediaPlaylist", "Icon": "playlist_play", "Severity": "info"},
		{"Name": "MediaQueue", "Icon": "queue", "Severity": "info"},
		{"Name": "MediaShuffle", "Icon": "shuffle", "Severity": "info"},
		{"Name": "MediaRepeat", "Icon": "repeat", "Severity": "info"},
		{"Name": "MediaRepeatOne", "Icon": "repeat_one", "Severity": "info"},
		{"Name": "MediaVolume", "Icon": "volume_up", "Severity": "info"},
		{"Name": "MediaMute", "Icon": "volume_off", "Severity": "info"},
		{"Name": "MediaSubtitle", "Icon": "subtitles", "Severity": "info"},
		{"Name": "MediaAudio", "Icon": "audio_file", "Severity": "info"},
		{"Name": "MediaVideo", "Icon": "videocam", "Severity": "info"},
		{"Name": "MediaImage", "Icon": "image", "Severity": "info"},
		{"Name": "MediaPhoto", "Icon": "photo", "Severity": "info"},
		{"Name": "MediaAlbum", "Icon": "album", "Severity": "info"},
		{"Name": "MediaArtist", "Icon": "person", "Severity": "info"},
		{"Name": "MediaGenre", "Icon": "category", "Severity": "info"},
		{"Name": "MediaStudio", "Icon": "business", "Severity": "info"},
		{"Name": "MediaDirector", "Icon": "director", "Severity": "info"},
		{"Name": "MediaWriter", "Icon": "write", "Severity": "info"},
		{"Name": "MediaProducer", "Icon": "production", "Severity": "info"},
		{"Name": "MediaComposer", "Icon": "music_note", "Severity": "info"},
		{"Name": "MediaEditor", "Icon": "edit", "Severity": "info"},
		{"Name": "MediaCinematographer", "Icon": "camera", "Severity": "info"},
		{"Name": "MediaPublisher", "Icon": "publish", "Severity": "info"},
		{"Name": "MediaAuthor", "Icon": "author", "Severity": "info"},
		{"Name": "MediaContributor", "Icon": "person_add", "Severity": "info"},
		{"Name": "MediaCollaborator", "Icon": "group", "Severity": "info"},
		{"Name": "MediaTeam", "Icon": "groups", "Severity": "info"},
		{"Name": "MediaOrganization", "Icon": "organization", "Severity": "info"},
		{"Name": "MediaCompany", "Icon": "company", "Severity": "info"},
		{"Name": "MediaEnterprise", "Icon": "enterprise", "Severity": "info"},
		{"Name": "MediaBusiness", "Icon": "business_center", "Severity": "info"},
		{"Name": "MediaCorporate", "Icon": "corporate_fare", "Severity": "info"},
		{"Name": "MediaInstitution", "Icon": "school", "Severity": "info"},
		{"Name": "MediaUniversity", "Icon": "university", "Severity": "info"},
		{"Name": "MediaCollege", "Icon": "college", "Severity": "info"},
		{"Name": "MediaSchool", "Icon": "school", "Severity": "info"},
		{"Name": "MediaAcademy", "Icon": "academy", "Severity": "info"},
		{"Name": "MediaInstitute", "Icon": "institute", "Severity": "info"},
		{"Name": "MediaFoundation", "Icon": "foundation", "Severity": "info"},
		{"Name": "MediaCharity", "Icon": "volunteer_activism", "Severity": "info"},
		{"Name": "MediaNonprofit", "Icon": "nonprofits", "Severity": "info"},
		{"Name": "MediaNGO", "Icon": "ngo", "Severity": "info"},
		{"Name": "MediaGovernment", "Icon": "government", "Severity": "info"},
		{"Name": "MediaPublic", "Icon": "public", "Severity": "info"},
		{"Name": "MediaPrivate", "Icon": "private_connectivity", "Severity": "info"},
		{"Name": "MediaConfidential", "Icon": "lock", "Severity": "warning"},
		{"Name": "MediaSecret", "Icon": "secret", "Severity": "warning"},
		{"Name": "MediaTopSecret", "Icon": "security", "Severity": "critical"},
		{"Name": "MediaClassified", "Icon": "classified", "Severity": "warning"},
		{"Name": "MediaRestricted", "Icon": "restriction", "Severity": "warning"},
		{"Name": "MediaControlled", "Icon": "control", "Severity": "info"},
		{"Name": "MediaManaged", "Icon": "manage_accounts", "Severity": "info"},
		{"Name": "MediaAdmin", "Icon": "admin_panel_settings", "Severity": "info"},
		{"Name": "MediaSuperAdmin", "Icon": "admin_panel_settings", "Severity": "info"},
		{"Name": "MediaRoot", "Icon": "root", "Severity": "info"},
		{"Name": "MediaUser", "Icon": "person", "Severity": "info"},
		{"Name": "MediaGuest", "Icon": "person_outline", "Severity": "info"},
		{"Name": "MediaVisitor", "Icon": "visitor", "Severity": "info"},
		{"Name": "MediaMember", "Icon": "member", "Severity": "info"},
		{"Name": "MediaSubscriber", "Icon": "subscriptions", "Severity": "info"},
		{"Name": "MediaFollower", "Icon": "follow_the_signs", "Severity": "info"},
		{"Name": "MediaFollowing", "Icon": "follow_the_signs", "Severity": "info"},
		{"Name": "MediaFriend", "Icon": "person_add", "Severity": "info"},
		{"Name": "MediaContact", "Icon": "contact_mail", "Severity": "info"},
		{"Name": "MediaEmail", "Icon": "email", "Severity": "info"},
		{"Name": "MediaPhone", "Icon": "phone", "Severity": "info"},
		{"Name": "MediaCall", "Icon": "call", "Severity": "info"},
		{"Name": "MediaMessage", "Icon": "message", "Severity": "info"},
		{"Name": "MediaChat", "Icon": "chat", "Severity": "info"},
		{"Name": "MediaText", "Icon": "text_fields", "Severity": "info"},
		{"Name": "MediaSMS", "Icon": "sms", "Severity": "info"},
		{"Name": "MediaMMS", "Icon": "mms", "Severity": "info"},
		{"Name": "MediaWhatsApp", "Icon": "whatsapp", "Severity": "info"},
		{"Name": "MediaTelegram", "Icon": "telegram", "Severity": "info"},
		{"Name": "MediaSignal", "Icon": "signal_cellular", "Severity": "info"},
		{"Name": "MediaViber", "Icon": "viber", "Severity": "info"},
		{"Name": "MediaWeChat", "Icon": "wechat", "Severity": "info"},
		{"Name": "MediaLine", "Icon": "line", "Severity": "info"},
		{"Name": "MediaFacebook", "Icon": "facebook", "Severity": "info"},
		{"Name": "MediaTwitter", "Icon": "twitter", "Severity": "info"},
		{"Name": "MediaInstagram", "Icon": "instagram", "Severity": "info"},
		{"Name": "MediaLinkedIn", "Icon": "linkedin", "Severity": "info"},
		{"Name": "MediaYouTube", "Icon": "youtube", "Severity": "info"},
		{"Name": "MediaTikTok", "Icon": "tiktok", "Severity": "info"},
		{"Name": "MediaSnapchat", "Icon": "snapchat", "Severity": "info"},
		{"Name": "MediaReddit", "Icon": "reddit", "Severity": "info"},
		{"Name": "MediaDiscord", "Icon": "discord", "Severity": "info"},
		{"Name": "MediaSlack", "Icon": "slack", "Severity": "info"},
		{"Name": "MediaTeams", "Icon": "teams", "Severity": "info"},
		{"Name": "MediaZoom", "Icon": "zoom", "Severity": "info"},
		{"Name": "MediaSkype", "Icon": "skype", "Severity": "info"},
		{"Name": "MediaGoogleMeet", "Icon": "google", "Severity": "info"},
		{"Name": "MediaMicrosoftTeams", "Icon": "microsoft", "Severity": "info"},
		{"Name": "MediaApple", "Icon": "apple", "Severity": "info"},
		{"Name": "MediaAndroid", "Icon": "android", "Severity": "info"},
		{"Name": "MediaWindows", "Icon": "windows", "Severity": "info"},
		{"Name": "MediaLinux", "Icon": "linux", "Severity": "info"},
		{"Name": "MediaMacOS", "Icon": "mac", "Severity": "info"},
		{"Name": "MediaiOS", "Icon": "ios", "Severity": "info"},
		{"Name": "MediaiPad", "Icon": "ipad", "Severity": "info"},
		{"Name": "MediaiPhone", "Icon": "iphone", "Severity": "info"},
		{"Name": "MediaAppleWatch", "Icon": "watch", "Severity": "info"},
		{"Name": "MediaAppleTV", "Icon": "tv", "Severity": "info"},
		{"Name": "MediaGoogleHome", "Icon": "google_home", "Severity": "info"},
		{"Name": "MediaGoogleAssistant", "Icon": "assistant", "Severity": "info"},
		{"Name": "MediaAlexa", "Icon": "alexa", "Severity": "info"},
		{"Name": "MediaGoogle", "Icon": "google", "Severity": "info"},
		{"Name": "MediaMicrosoft", "Icon": "microsoft", "Severity": "info"},
		{"Name": "MediaApple", "Icon": "apple", "Severity": "info"},
		{"Name": "MediaAmazon", "Icon": "amazon", "Severity": "info"},
		{"Name": "MediaGooglePlay", "Icon": "play_circle", "Severity": "info"},
		{"Name": "MediaAppStore", "Icon": "app_store", "Severity": "info"},
		{"Name": "MediaWindowsStore", "Icon": "windows", "Severity": "info"},
		{"Name": "MediaLinuxStore", "Icon": "linux", "Severity": "info"},
		{"Name": "MediaMacAppStore", "Icon": "mac", "Severity": "info"},
		{"Name": "MediaiOSAppStore", "Icon": "ios", "Severity": "info"},
		{"Name": "MediaAndroidAppStore", "Icon": "android", "Severity": "info"},
		{"Name": "MediaGooglePlayStore", "Icon": "play_circle", "Severity": "info"},
		{"Name": "MediaAppleAppStore", "Icon": "app_store", "Severity": "info"},
		{"Name": "MediaMicrosoftStore", "Icon": "microsoft", "Severity": "info"},
		{"Name": "MediaAmazonStore", "Icon": "amazon", "Severity": "info"},
		{"Name": "MediaGoogleStore", "Icon": "google", "Severity": "info"},
		{"Name": "MediaAppleStore", "Icon": "apple", "Severity": "info"},
		{"Name": "MediaMicrosoftStore", "Icon": "microsoft", "Severity": "info"},
		{"Name": "MediaAmazonStore", "Icon": "amazon", "Severity": "info"},
		{"Name": "MediaGoogleStore", "Icon": "google", "Severity": "info"},
		{"Name": "MediaAppleStore", "Icon": "apple", "Severity": "info"},
	}
}
