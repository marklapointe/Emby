package session

import (
	"fmt"
	"sync"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/model"
	"go.uber.org/zap"
)

// SessionInfo represents an active playback session.
type SessionInfo struct {
	ID               string                 `json:"Id"`
	Client           string                 `json:"Client"`
	DeviceName       string                 `json:"DeviceName"`
	DisplayName      string                 `json:"DisplayName"`
	MachineID        string                 `json:"MachineId"`
	LastActivityTime time.Time              `json:"LastActivityTime"`
	PlayState        model.PlayState        `json:"PlayState"`
	Location         model.Location         `json:"Location"`
	Capabilities     model.Capabilities     `json:"Capabilities"`
	SupportedMediaTypes []string            `json:"SupportedMediaTypes"`
	RemoteImageURL   string                 `json:"RemoteImageUrl,omitempty"`
	StartTimeTicks   int64                  `json:"StartTimeTicks"`
}

// Manager handles session-related operations.
type Manager struct {
	config *config.Config
	logger *zap.Logger
	mu     sync.RWMutex
	sessions map[string]*SessionInfo
}

// NewManager creates a new session manager.
func NewManager(cfg *config.Config, logger *zap.Logger) *Manager {
	return &Manager{
		config:   cfg,
		logger:   logger,
		sessions: make(map[string]*SessionInfo),
	}
}

// CreateSession creates a new session.
func (m *Manager) CreateSession(session *SessionInfo) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.sessions[session.ID]; exists {
		return fmt.Errorf("session already exists: %s", session.ID)
	}

	m.sessions[session.ID] = session
	if m.logger != nil {
		m.logger.Info("session created", zap.String("id", session.ID))
	}
	return nil
}

// GetSession returns a session by ID.
func (m *Manager) GetSession(id string) (*SessionInfo, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	session, exists := m.sessions[id]
	return session, exists
}

// GetAllSessions returns all sessions.
func (m *Manager) GetAllSessions() []*SessionInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	sessions := make([]*SessionInfo, 0, len(m.sessions))
	for _, session := range m.sessions {
		sessions = append(sessions, session)
	}
	return sessions
}

// UpdateSession updates a session's information.
func (m *Manager) UpdateSession(id string, position *int64, volume *int, isPaused *bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	if position != nil {
		session.PlayState.PositionTicks = *position
	}
	if volume != nil {
		session.PlayState.VolumePercent = *volume
	}
	if isPaused != nil {
		session.PlayState.IsPaused = *isPaused
	}

	session.LastActivityTime = time.Now()
	return nil
}

// DeleteSession deletes a session.
func (m *Manager) DeleteSession(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.sessions[id]; !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	delete(m.sessions, id)
	if m.logger != nil {
		m.logger.Info("session deleted", zap.String("id", id))
	}
	return nil
}

// GetSessionsByDevice returns all sessions for a device.
func (m *Manager) GetSessionsByDevice(deviceName string) []*SessionInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var sessions []*SessionInfo
	for _, session := range m.sessions {
		if session.DeviceName == deviceName {
			sessions = append(sessions, session)
		}
	}
	return sessions
}

// GetSessionsByUser returns all sessions for a user.
func (m *Manager) GetSessionsByUser(displayName string) []*SessionInfo {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var sessions []*SessionInfo
	for _, session := range m.sessions {
		if session.DisplayName == displayName {
			sessions = append(sessions, session)
		}
	}
	return sessions
}

// GetActiveSessionCount returns the number of active sessions.
func (m *Manager) GetActiveSessionCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	count := 0
	for _, session := range m.sessions {
		if time.Since(session.LastActivityTime) < 24*time.Hour {
			count++
		}
	}
	return count
}

// StartPlayback starts playback for a session.
func (m *Manager) StartPlayback(id string, itemId string, mediaSourceId string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	session.PlayState.PositionTicks = 0
	session.PlayState.IsPaused = false
	session.LastActivityTime = time.Now()

	return nil
}

// StopPlayback stops playback for a session.
func (m *Manager) StopPlayback(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	session.PlayState.PositionTicks = 0
	session.PlayState.IsPaused = false
	session.LastActivityTime = time.Now()

	return nil
}

// SendCommand sends a command to a session.
func (m *Manager) SendCommand(id string, command string, args map[string]interface{}) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	if m.logger != nil {
		m.logger.Info("session command", zap.String("id", id), zap.String("command", command))
	}

	session.LastActivityTime = time.Now()
	return nil
}

// AddUserToSession adds a user to a session.
func (m *Manager) AddUserToSession(id string, userId string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	if m.logger != nil {
		m.logger.Info("user added to session", zap.String("sessionId", id), zap.String("userId", userId))
	}

	session.LastActivityTime = time.Now()
	return nil
}

// RemoveUserFromSession removes a user from a session.
func (m *Manager) RemoveUserFromSession(id string, userId string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	session, exists := m.sessions[id]
	if !exists {
		return fmt.Errorf("session not found: %s", id)
	}

	if m.logger != nil {
		m.logger.Info("user removed from session", zap.String("sessionId", id), zap.String("userId", userId))
	}

	session.LastActivityTime = time.Now()
	return nil
}
