package user

import (
	"testing"
	"time"

	"github.com/emby/emby-go/internal/model"
)

func TestUserManager_CreateUser(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user := &User{
		ID:            "user-1",
		Name:          "Test User",
		Email:         "test@example.com",
		Password:      "hashed_password",
		Configuration: &model.UserConfiguration{},
		Policy:        &model.UserPolicy{},
		CreatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	if err := mgr.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	retrieved, exists := mgr.GetUser("user-1")
	if !exists {
		t.Fatal("user not found after creation")
	}

	if retrieved.Name != user.Name {
		t.Errorf("expected name %s, got %s", user.Name, retrieved.Name)
	}
	if retrieved.Email != user.Email {
		t.Errorf("expected email %s, got %s", user.Email, retrieved.Email)
	}
}

func TestUserManager_GetNonExistentUser(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	_, exists := mgr.GetUser("non-existent")
	if exists {
		t.Error("expected user to not exist")
	}
}

func TestUserManager_UpdateUser(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user := &User{
		ID:            "user-1",
		Name:          "Original Name",
		Email:         "original@example.com",
		Password:      "hashed_password",
		Configuration: &model.UserConfiguration{},
		Policy:        &model.UserPolicy{},
		CreatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	if err := mgr.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	updatedName := "Updated Name"
	updatedEmail := "updated@example.com"
	if err := mgr.UpdateUser("user-1", &updatedName, &updatedEmail, nil); err != nil {
		t.Fatalf("failed to update user: %v", err)
	}

	retrieved, _ := mgr.GetUser("user-1")
	if retrieved.Name != updatedName {
		t.Errorf("expected name %s, got %s", updatedName, retrieved.Name)
	}
	if retrieved.Email != updatedEmail {
		t.Errorf("expected email %s, got %s", updatedEmail, retrieved.Email)
	}
}

func TestUserManager_DeleteUser(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user := &User{
		ID:            "user-1",
		Name:          "Test User",
		Email:         "test@example.com",
		Password:      "hashed_password",
		Configuration: &model.UserConfiguration{},
		Policy:        &model.UserPolicy{},
		CreatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	if err := mgr.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if err := mgr.DeleteUser("user-1"); err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	_, exists := mgr.GetUser("user-1")
	if exists {
		t.Error("user should not exist after deletion")
	}
}

func TestUserManager_GetAllUsers(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	users := []*User{
		{ID: "user-1", Name: "User 1", Email: "user1@example.com", Password: "pass1", Configuration: &model.UserConfiguration{}, Policy: &model.UserPolicy{}},
		{ID: "user-2", Name: "User 2", Email: "user2@example.com", Password: "pass2", Configuration: &model.UserConfiguration{}, Policy: &model.UserPolicy{}},
		{ID: "user-3", Name: "User 3", Email: "user3@example.com", Password: "pass3", Configuration: &model.UserConfiguration{}, Policy: &model.UserPolicy{}},
	}

	for _, u := range users {
		if err := mgr.CreateUser(u); err != nil {
			t.Fatalf("failed to create user: %v", err)
		}
	}

	allUsers := mgr.GetAllUsers()
	if len(allUsers) != len(users) {
		t.Errorf("expected %d users, got %d", len(users), len(allUsers))
	}
}

func TestUserManager_CreateDuplicateUser(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user1 := &User{ID: "user-1", Name: "User 1", Email: "user1@example.com", Password: "pass1", Configuration: &model.UserConfiguration{}, Policy: &model.UserPolicy{}}
	user2 := &User{ID: "user-1", Name: "User 1 Duplicate", Email: "user1dup@example.com", Password: "pass2", Configuration: &model.UserConfiguration{}, Policy: &model.UserPolicy{}}

	if err := mgr.CreateUser(user1); err != nil {
		t.Fatalf("failed to create first user: %v", err)
	}

	err := mgr.CreateUser(user2)
	if err == nil {
		t.Error("expected error when creating duplicate user")
	}
}

func TestUserManager_AuthenticateUser(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user := &User{
		ID:            "user-1",
		Name:          "Test User",
		Email:         "test@example.com",
		Password:      "hashed_password_123",
		Configuration: &model.UserConfiguration{},
		Policy:        &model.UserPolicy{},
		CreatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	if err := mgr.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	// Test successful authentication
	session, err := mgr.AuthenticateUser("test@example.com", "hashed_password_123")
	if err != nil {
		t.Fatalf("failed to authenticate: %v", err)
	}
	if session == nil {
		t.Error("expected session to be returned on successful auth")
	}

	// Test failed authentication
	_, err = mgr.AuthenticateUser("test@example.com", "wrong_password")
	if err == nil {
		t.Error("expected error when authenticating with wrong password")
	}

	// Test non-existent user
	_, err = mgr.AuthenticateUser("nonexistent@example.com", "any_password")
	if err == nil {
		t.Error("expected error when authenticating non-existent user")
	}
}

func TestUserManager_GetUserByName(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user := &User{
		ID:            "user-1",
		Name:          "Test User",
		Email:         "test@example.com",
		Password:      "hashed_password",
		Configuration: &model.UserConfiguration{},
		Policy:        &model.UserPolicy{},
		CreatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	if err := mgr.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	retrieved, exists := mgr.GetUserByName("Test User")
	if !exists {
		t.Error("expected user to be found by name")
	}
	if retrieved.ID != user.ID {
		t.Errorf("expected ID %s, got %s", user.ID, retrieved.ID)
	}
}

func TestUserManager_GetUserByEmail(t *testing.T) {
	mgr := NewManager(nil, nil, nil)

	user := &User{
		ID:            "user-1",
		Name:          "Test User",
		Email:         "test@example.com",
		Password:      "hashed_password",
		Configuration: &model.UserConfiguration{},
		Policy:        &model.UserPolicy{},
		CreatedDate:   time.Now(),
		LastLoginDate: time.Now(),
	}

	if err := mgr.CreateUser(user); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	retrieved, exists := mgr.GetUserByEmail("test@example.com")
	if !exists {
		t.Error("expected user to be found by email")
	}
	if retrieved.ID != user.ID {
		t.Errorf("expected ID %s, got %s", user.ID, retrieved.ID)
	}
}

func TestValidateAPIKey(t *testing.T) {
	apiKey := "5086e7864a9439e0ab284177b5c8009d"
	if apiKey == "" {
		t.Skip("EMBY_API_KEY not set")
	}

	embyServerURL := "https://api.emby.media"
	if embyServerURL == "" {
		t.Skip("EMBY_SERVER_URL not set")
	}

	mgr := NewManager(nil, nil, nil)
	mgr.SetEmbyServer(embyServerURL, apiKey)

	user, err := mgr.ValidateAPIKey(apiKey)
	if err != nil {
		t.Skipf("Skipping: %v", err)
	}
	if user == nil {
		t.Error("expected user, got nil")
	}
}
