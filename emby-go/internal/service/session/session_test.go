package session

import (
	"fmt"
	"testing"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/model"
	"go.uber.org/zap"
)

func TestNewManager(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)
	if mgr == nil {
		t.Fatal("expected manager to be created")
	}
}

func TestManager_CreateSession(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	session := &SessionInfo{
		ID:               "session-123",
		Client:           "Emby for Android",
		DeviceName:       "Test Device",
		DisplayName:      "Test User",
		MachineID:        "machine-456",
		LastActivityTime: time.Now(),
		PlayState: model.PlayState{
			PositionTicks:   0,
			VolumePercent:   80,
			IsMuted:         false,
			IsPaused:        false,
		},
	}

	if err := mgr.CreateSession(session); err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	retrieved, exists := mgr.GetSession("session-123")
	if !exists {
		t.Fatal("session not found after creation")
	}

	if retrieved.ID != session.ID {
		t.Errorf("expected ID %s, got %s", session.ID, retrieved.ID)
	}
	if retrieved.Client != session.Client {
		t.Errorf("expected Client %s, got %s", session.Client, retrieved.Client)
	}
}

func TestManager_GetNonExistentSession(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	_, exists := mgr.GetSession("non-existent")
	if exists {
		t.Error("expected session to not exist")
	}
}

func TestManager_UpdateSession(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	session := &SessionInfo{
		ID:               "session-123",
		Client:           "Emby for Android",
		DeviceName:       "Test Device",
		DisplayName:      "Test User",
		MachineID:        "machine-456",
		LastActivityTime: time.Now(),
		PlayState: model.PlayState{
			PositionTicks:   0,
			VolumePercent:   80,
			IsMuted:         false,
			IsPaused:        false,
		},
	}

	if err := mgr.CreateSession(session); err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	// Update the session
	newPosition := int64(1234567890)
	if err := mgr.UpdateSession("session-123", &newPosition, nil, nil); err != nil {
		t.Fatalf("failed to update session: %v", err)
	}

	retrieved, exists := mgr.GetSession("session-123")
	if !exists {
		t.Fatal("session not found after update")
	}

	if retrieved.PlayState.PositionTicks != newPosition {
		t.Errorf("expected position %d, got %d", newPosition, retrieved.PlayState.PositionTicks)
	}
}

func TestManager_DeleteSession(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	session := &SessionInfo{
		ID:               "session-123",
		Client:           "Emby for Android",
		DeviceName:       "Test Device",
		DisplayName:      "Test User",
		MachineID:        "machine-456",
		LastActivityTime: time.Now(),
		PlayState: model.PlayState{
			PositionTicks:   0,
			VolumePercent:   80,
			IsMuted:         false,
			IsPaused:        false,
		},
	}

	if err := mgr.CreateSession(session); err != nil {
		t.Fatalf("failed to create session: %v", err)
	}

	if err := mgr.DeleteSession("session-123"); err != nil {
		t.Fatalf("failed to delete session: %v", err)
	}

	_, exists := mgr.GetSession("session-123")
	if exists {
		t.Error("session should not exist after deletion")
	}
}

func TestManager_GetAllSessions(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	sessions := []*SessionInfo{
		{ID: "session-1", Client: "Client 1", DeviceName: "Device 1", DisplayName: "User 1", MachineID: "machine-1", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
		{ID: "session-2", Client: "Client 2", DeviceName: "Device 2", DisplayName: "User 2", MachineID: "machine-2", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
		{ID: "session-3", Client: "Client 3", DeviceName: "Device 3", DisplayName: "User 3", MachineID: "machine-3", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
	}

	for _, s := range sessions {
		if err := mgr.CreateSession(s); err != nil {
			t.Fatalf("failed to create session: %v", err)
		}
	}

	allSessions := mgr.GetAllSessions()
	if len(allSessions) != len(sessions) {
		t.Errorf("expected %d sessions, got %d", len(sessions), len(allSessions))
	}
}

func TestManager_GetSessionsByDevice(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	sessions := []*SessionInfo{
		{ID: "session-1", Client: "Client 1", DeviceName: "Device 1", DisplayName: "User 1", MachineID: "machine-1", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
		{ID: "session-2", Client: "Client 2", DeviceName: "Device 1", DisplayName: "User 2", MachineID: "machine-2", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
		{ID: "session-3", Client: "Client 3", DeviceName: "Device 2", DisplayName: "User 3", MachineID: "machine-3", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
	}

	for _, s := range sessions {
		if err := mgr.CreateSession(s); err != nil {
			t.Fatalf("failed to create session: %v", err)
		}
	}

	deviceSessions := mgr.GetSessionsByDevice("Device 1")
	if len(deviceSessions) != 2 {
		t.Errorf("expected 2 sessions for Device 1, got %d", len(deviceSessions))
	}
}

func TestManager_GetSessionsByUser(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	sessions := []*SessionInfo{
		{ID: "session-1", Client: "Client 1", DeviceName: "Device 1", DisplayName: "User 1", MachineID: "machine-1", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
		{ID: "session-2", Client: "Client 2", DeviceName: "Device 2", DisplayName: "User 1", MachineID: "machine-2", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
		{ID: "session-3", Client: "Client 3", DeviceName: "Device 3", DisplayName: "User 2", MachineID: "machine-3", LastActivityTime: time.Now(), PlayState: model.PlayState{PositionTicks: 0, VolumePercent: 80, IsMuted: false, IsPaused: false}},
	}

	for _, s := range sessions {
		if err := mgr.CreateSession(s); err != nil {
			t.Fatalf("failed to create session: %v", err)
		}
	}

	userSessions := mgr.GetSessionsByUser("User 1")
	if len(userSessions) != 2 {
		t.Errorf("expected 2 sessions for User 1, got %d", len(userSessions))
	}
}

func TestManager_GetActiveSessionCount(t *testing.T) {
	cfg := &config.Config{}

	logger, _ := zap.NewDevelopment()
	mgr := NewManager(cfg, logger)

	// Create sessions with recent activity
	for i := 0; i < 5; i++ {
		session := &SessionInfo{
			ID:               fmt.Sprintf("session-%d", i),
			Client:           "Emby for Android",
			DeviceName:       "Test Device",
			DisplayName:      "Test User",
			MachineID:        "machine-456",
			LastActivityTime: time.Now(),
			PlayState: model.PlayState{
				PositionTicks: 0,
				VolumePercent: 80,
				IsMuted:       false,
				IsPaused:      false,
			},
		}
		if err := mgr.CreateSession(session); err != nil {
			t.Fatalf("failed to create session: %v", err)
		}
	}

	activeCount := mgr.GetActiveSessionCount()
	if activeCount != 5 {
		t.Errorf("expected 5 active sessions, got %d", activeCount)
	}
}
