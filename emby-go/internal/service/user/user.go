package user

import (
	"fmt"
	"sync"
	"time"

	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/model"
	"github.com/emby/emby-go/internal/repository"
	"go.uber.org/zap"
)

// User represents a server user.
type User struct {
	ID              string                 `json:"Id"`
	Name            string                 `json:"Name"`
	Email           string                 `json:"Email,omitempty"`
	Password        string                 `json:"-"`
	Configuration   *model.UserConfiguration `json:"Configuration"`
	Policy          *model.UserPolicy      `json:"Policy"`
	CreatedDate     time.Time              `json:"CreatedDate"`
	LastLoginDate   time.Time              `json:"LastLoginDate"`
	InvalidLoginAttemptCount int         `json:"InvalidLoginAttemptCount"`
	LastActivityDate time.Time             `json:"LastActivityDate"`
}

// UserConfiguration represents user-specific settings.
type UserConfiguration struct {
	SubtitleMode       string `json:"SubtitleMode"`
	SubtitleFontSize   int    `json:"SubtitleFontSize"`
	SubtitleLanguage   string `json:"SubtitleLanguage"`
	MissingEpisodes    bool   `json:"MissingEpisodes"`
	MissingEpisodesSection int `json:"MissingEpisodesSection"`
	HidePlayedInLatest bool   `json:"HidePlayedInLatest"`
	NextEpisodeDays    int    `json:"NextEpisodeDays"`
	PlayNextEpisode    bool   `json:"PlayNextEpisode"`
	GroupedFolders     []string `json:"GroupedFolders"`
	SkipForwardDuration int  `json:"SkipForwardDuration"`
	SkipBackwardDuration int `json:"SkipBackwardDuration"`
	TimedViewingData   string `json:"TimedViewingData"`
	MaxAudioChannels   string `json:"MaxAudioChannels"`
	EnableAutoStart    bool   `json:"EnableAutoStart"`
	MinRating          int    `json:"MinRating"`
	PlayedIndicator    string `json:"PlayedIndicator"`
	OrderItemsBy       []string `json:"OrderItemsBy"`
	RememberAudioSelections bool `json:"RememberAudioSelections"`
	RememberSubtitleSelections bool `json:"RememberSubtitleSelections"`
	EnabledChannels    []interface{} `json:"EnabledChannels"`
	LatestItemsExcluded []string `json:"LatestItemsExcluded"`
}

// UserPolicy represents a user's permissions and access controls.
type UserPolicy struct {
	IsAdministrator      bool     `json:"IsAdministrator"`
	IsHidden             bool     `json:"IsHidden"`
	IsDisabled           bool     `json:"IsDisabled"`
	BlockedTags          []string `json:"BlockedTags"`
	AllowTagging         bool     `json:"AllowTagging"`
	EnableRemoteControlOfOtherUsers bool `json:"EnableRemoteControlOfOtherUsers"`
	SharedServers        []string `json:"SharedServers"`
	Profile              string   `json:"Profile"`
	ForceRemoteSourceTranscoding bool `json:"ForceRemoteSourceTranscoding"`
	EnableContentDeletion    bool   `json:"EnableContentDeletion"`
	EnableContentDownloading bool   `json:"EnableContentDownloading"`
	AllowCameraUpload        bool   `json:"AllowCameraUpload"`
	AllowShareAllFolders     bool   `json:"AllowShareAllFolders"`
	MaxActiveSessions        int    `json:"MaxActiveSessions"`
	BlockUnratedItems        []string `json:"BlockUnratedItems"`
	BlockedChannels          []string `json:"BlockedChannels"`
	AllowChannelAccess       []string `json:"AllowChannelAccess"`
	EnablePublicSharing      bool   `json:"EnablePublicSharing"`
	LockedFields             []string `json:"LockedFields"`
	PasswordResetKeyLifetime time.Duration `json:"PasswordResetKeyLifetime"`
	Configuration            *UserConfiguration `json:"Configuration"`
}

// Session represents an authentication session.
type Session struct {
	Token       string    `json:"Token"`
	UserID      string    `json:"UserId"`
	CreatedAt   time.Time `json:"CreatedAt"`
	ExpiresAt   time.Time `json:"ExpiresAt"`
}

// Manager handles user-related operations.
type Manager struct {
	dbManager *database.Manager
	userRepo  *repository.UserRepository
	logger    *zap.Logger
	mu        sync.RWMutex
	users     map[string]*User
	sessions  map[string]*Session
}

// NewManager creates a new user manager.
func NewManager(dbManager *database.Manager, userRepo *repository.UserRepository, logger *zap.Logger) *Manager {
	return &Manager{
		dbManager: dbManager,
		userRepo:  userRepo,
		logger:    logger,
		users:     make(map[string]*User),
		sessions:  make(map[string]*Session),
	}
}

// CreateUser creates a new user.
func (m *Manager) CreateUser(user *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.users[user.ID]; exists {
		return fmt.Errorf("user already exists: %s", user.ID)
	}

	if user.Configuration == nil {
		user.Configuration = &model.UserConfiguration{}
	}
	if user.Policy == nil {
		user.Policy = &model.UserPolicy{}
	}

	m.users[user.ID] = user
	if m.logger != nil {
		m.logger.Info("user created", zap.String("id", user.ID), zap.String("name", user.Name))
	}
	return nil
}

// GetUser returns a user by ID.
func (m *Manager) GetUser(id string) (*User, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	user, exists := m.users[id]
	return user, exists
}

// GetUserByName returns a user by name.
func (m *Manager) GetUserByName(name string) (*User, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, user := range m.users {
		if user.Name == name {
			return user, true
		}
	}
	return nil, false
}

// GetUserByEmail returns a user by email.
func (m *Manager) GetUserByEmail(email string) (*User, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, user := range m.users {
		if user.Email == email {
			return user, true
		}
	}
	return nil, false
}

// GetAllUsers returns all users.
func (m *Manager) GetAllUsers() []*User {
	m.mu.RLock()
	defer m.mu.RUnlock()

	users := make([]*User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users
}

// EnsureDefaultUser creates a default user if no users exist.
// This matches the C# server behavior where at least one user must exist.
func (m *Manager) EnsureDefaultUser() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.users) > 0 {
		return nil
	}

	defaultName := "MyEmbyUser"

	user := &User{
		ID:          fmt.Sprintf("user-%d", time.Now().UnixNano()),
		Name:        defaultName,
		Password:    "",
		Configuration: &model.UserConfiguration{},
		Policy: &model.UserPolicy{
			IsAdministrator:      true,
			EnableMediaDeletion: true,
		},
		CreatedDate: time.Now(),
	}

	m.users[user.ID] = user

	if m.logger != nil {
		m.logger.Info("default user created", zap.String("id", user.ID), zap.String("name", user.Name))
	}

	return nil
}

// UpdateUser updates a user's information.
func (m *Manager) UpdateUser(id string, name, email *string, password *string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.users[id]
	if !exists {
		return fmt.Errorf("user not found: %s", id)
	}

	if name != nil {
		user.Name = *name
	}
	if email != nil {
		user.Email = *email
	}
	if password != nil {
		user.Password = *password
	}

	if m.logger != nil {
		m.logger.Info("user updated", zap.String("id", id))
	}
	return nil
}

// UpdateUserConfiguration updates a user's configuration.
func (m *Manager) UpdateUserConfiguration(id string, config *model.UserConfiguration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.users[id]
	if !exists {
		return fmt.Errorf("user not found: %s", id)
	}

	user.Configuration = config
	if m.logger != nil {
		m.logger.Info("user configuration updated", zap.String("id", id))
	}
	return nil
}

// UpdateUserPolicy updates a user's policy.
func (m *Manager) UpdateUserPolicy(id string, policy *model.UserPolicy) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.users[id]
	if !exists {
		return fmt.Errorf("user not found: %s", id)
	}

	user.Policy = policy
	if m.logger != nil {
		m.logger.Info("user policy updated", zap.String("id", id))
	}
	return nil
}

// DeleteUser deletes a user.
func (m *Manager) DeleteUser(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.users[id]; !exists {
		return fmt.Errorf("user not found: %s", id)
	}

	delete(m.users, id)
	if m.logger != nil {
		m.logger.Info("user deleted", zap.String("id", id))
	}
	return nil
}

func (m *Manager) AuthenticateUser(username, password string) (*Session, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	var userID string

	if m.userRepo != nil {
		user, err := m.userRepo.Authenticate(username, password)
		if err != nil {
			return nil, fmt.Errorf("invalid credentials")
		}
		userID = user.Id
	} else {
		var memUser *User
		for _, u := range m.users {
			if u.Name == username || u.Email == username {
				memUser = u
				break
			}
		}
		if memUser == nil {
			return nil, fmt.Errorf("invalid credentials")
		}
		if memUser.Password != password {
			return nil, fmt.Errorf("invalid credentials")
		}
		userID = memUser.ID
	}

	token := generateToken()
	session := &Session{
		Token:     token,
		UserID:    userID,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	m.sessions[token] = session

	if m.logger != nil {
		m.logger.Info("user authenticated", zap.String("id", userID))
	}
	return session, nil
}

// ValidateSession validates a session token.
func (m *Manager) ValidateSession(token string) (*Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	session, exists := m.sessions[token]
	if !exists {
		return nil, fmt.Errorf("invalid session")
	}

	if time.Now().After(session.ExpiresAt) {
		delete(m.sessions, token)
		return nil, fmt.Errorf("session expired")
	}

	return session, nil
}

// RevokeSession revokes a session token.
func (m *Manager) RevokeSession(token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.sessions[token]; !exists {
		return fmt.Errorf("session not found")
	}

	delete(m.sessions, token)
	return nil
}

// generateToken generates a random session token.
func generateToken() string {
	return fmt.Sprintf("emby-%d-%d", time.Now().UnixNano(), randInt())
}

func randInt() int {
	return int(time.Now().UnixNano() % 1000000)
}
