package auth

import (
	"testing"
)

func TestNewUserManager(t *testing.T) {
	m := NewUserManager()
	if m == nil {
		t.Fatal("NewUserManager returned nil")
	}
	if m.sessions == nil {
		t.Error("sessions map not initialized")
	}
}

func TestHashPassword(t *testing.T) {
	hash := HashPassword("testpassword")
	if hash == "" {
		t.Error("hash is empty")
	}
	if hash == "testpassword" {
		t.Error("hash should not be plain password")
	}
}

func TestHashPassword_Consistency(t *testing.T) {
	hash1 := HashPassword("testpassword")
	hash2 := HashPassword("testpassword")
	if hash1 != hash2 {
		t.Error("same password should produce same hash")
	}
}

func TestVerifyPassword(t *testing.T) {
	hash := HashPassword("testpassword")

	if !VerifyPassword("testpassword", hash) {
		t.Error("VerifyPassword failed for correct password")
	}

	if VerifyPassword("wrongpassword", hash) {
		t.Error("VerifyPassword should fail for wrong password")
	}
}

func TestAuthenticateUser(t *testing.T) {
	m := NewUserManager()

	session, err := m.AuthenticateUser("testuser", "testpass")
	if err != nil {
		t.Errorf("AuthenticateUser returned error: %v", err)
	}
	if session == nil {
		t.Fatal("session is nil")
	}
	if session.ID == "" {
		t.Error("session ID is empty")
	}
	if session.AccessToken == "" {
		t.Error("access token is empty")
	}
}

func TestValidateSession(t *testing.T) {
	m := NewUserManager()

	session, _ := m.AuthenticateUser("testuser", "testpass")
	validated, err := m.ValidateSession(session.AccessToken)

	if err != nil {
		t.Errorf("ValidateSession returned error: %v", err)
	}
	if validated.UserID != session.UserID {
		t.Error("validated session userID mismatch")
	}
}

func TestValidateSession_Invalid(t *testing.T) {
	m := NewUserManager()
	_, err := m.ValidateSession("invalid-token")
	if err == nil {
		t.Error("expected error for invalid token")
	}
}

func TestInvalidateSession(t *testing.T) {
	m := NewUserManager()

	session, _ := m.AuthenticateUser("testuser", "testpass")
	err := m.InvalidateSession(session.AccessToken)
	if err != nil {
		t.Errorf("InvalidateSession returned error: %v", err)
	}

	_, err = m.ValidateSession(session.AccessToken)
	if err == nil {
		t.Error("expected error after invalidation")
	}
}

func TestGetSession(t *testing.T) {
	m := NewUserManager()

	session, _ := m.AuthenticateUser("testuser", "testpass")
	got, err := m.GetSession(session.AccessToken)

	if err != nil {
		t.Errorf("GetSession returned error: %v", err)
	}
	if got.ID != session.ID {
		t.Errorf("expected ID '%s', got '%s'", session.ID, got.ID)
	}
}

func TestGetActiveSessions(t *testing.T) {
	m := NewUserManager()

	session1, _ := m.AuthenticateUser("user1", "pass")
	session2, _ := m.AuthenticateUser("user2", "pass")

	sessions := m.GetActiveSessions("")
	if len(sessions) < 2 {
		t.Errorf("expected at least 2 sessions, got %d", len(sessions))
	}

	_ = session1
	_ = session2
}