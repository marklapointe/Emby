package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
	"time"
)

// Session represents an authenticated user session.
type Session struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	UserName  string    `json:"userName"`
	AccessToken string  `json:"accessToken"`
	IPAddress string    `json:"ipAddress"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt"`
	IsActive  bool      `json:"isActive"`
}

// UserManager handles user authentication.
type UserManager struct {
	sessions map[string]*Session
	mu       sync.RWMutex
}

// NewUserManager creates a new user manager.
func NewUserManager() *UserManager {
	return &UserManager{
		sessions: make(map[string]*Session),
	}
}

// AuthenticateUser authenticates a user and creates a session.
func (m *UserManager) AuthenticateUser(username, password string) (*Session, error) {
	// For now, return a placeholder session
	_ = username
	_ = password

	sessionID := fmt.Sprintf("session-%d", time.Now().UnixNano())
	accessToken := generateToken()

	session := &Session{
		ID:            sessionID,
		AccessToken:   accessToken,
		CreatedAt:     time.Now(),
		ExpiresAt:     time.Now().Add(24 * time.Hour),
		IsActive:      true,
	}

	m.mu.Lock()
	m.sessions[sessionID] = session
	m.mu.Unlock()

	return session, nil
}

// ValidateSession validates a session token.
func (m *UserManager) ValidateSession(token string) (*Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, session := range m.sessions {
		if session.AccessToken == token && session.IsActive {
			if time.Now().After(session.ExpiresAt) {
				session.IsActive = false
				return nil, fmt.Errorf("session expired")
			}
			return session, nil
		}
	}

	return nil, fmt.Errorf("invalid session")
}

// InvalidateSession invalidates a session.
func (m *UserManager) InvalidateSession(token string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, session := range m.sessions {
		if session.AccessToken == token {
			session.IsActive = false
			return nil
		}
	}

	return fmt.Errorf("session not found")
}

// GetSession returns a session by token.
func (m *UserManager) GetSession(token string) (*Session, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, session := range m.sessions {
		if session.AccessToken == token {
			return session, nil
		}
	}

	return nil, fmt.Errorf("session not found")
}

// GetActiveSessions returns all active sessions for a user.
func (m *UserManager) GetActiveSessions(userID string) []*Session {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var sessions []*Session
	for _, session := range m.sessions {
		if session.UserID == userID && session.IsActive {
			sessions = append(sessions, session)
		}
	}

	return sessions
}

// HashPassword hashes a password using SHA-256.
func HashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

// VerifyPassword verifies a password against a hash.
func VerifyPassword(password, hash string) bool {
	return HashPassword(password) == hash
}

// generateToken generates a random access token.
func generateToken() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%d", time.Now().UnixNano()))))
}
