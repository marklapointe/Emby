package repository

import (
	"testing"

	"github.com/emby/emby-go/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUserTestDB(t *testing.T) (*gorm.DB, func()) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	err = db.AutoMigrate(&model.GORMUser{})
	if err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	cleanup := func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}
	return db, cleanup
}

func TestNewUserRepository(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	if repo == nil {
		t.Fatal("expected non-nil repository")
	}
	if repo.db == nil {
		t.Error("expected non-nil db")
	}
}

func TestCreateUser(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Name: "Test User",
	}

	err := repo.CreateUser(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	if user.Id == "" {
		t.Error("expected user ID to be set")
	}
}

func TestCreateUserWithID(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "custom-id",
		Name: "Test User",
	}

	err := repo.CreateUser(user)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	if user.Id != "custom-id" {
		t.Errorf("expected custom ID to be preserved, got %s", user.Id)
	}
}

func TestGetUserByID(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "test-user-1",
		Name: "Test User",
	}
	repo.CreateUser(user)

	found, err := repo.GetUserByID("test-user-1")
	if err != nil {
		t.Fatalf("failed to get user: %v", err)
	}
	if found.Name != "Test User" {
		t.Errorf("expected name 'Test User', got %s", found.Name)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	_, err := repo.GetUserByID("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent user")
	}
}

func TestGetUserByName(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "test-user-2",
		Name: "FindMe",
	}
	repo.CreateUser(user)

	found, err := repo.GetUserByName("FindMe")
	if err != nil {
		t.Fatalf("failed to get user by name: %v", err)
	}
	if found.Id != "test-user-2" {
		t.Errorf("expected id 'test-user-2', got %s", found.Id)
	}
}

func TestGetAllUsers(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	repo.CreateUser(&model.GORMUser{Name: "Alice"})
	repo.CreateUser(&model.GORMUser{Name: "Bob"})

	users, err := repo.GetAllUsers()
	if err != nil {
		t.Fatalf("failed to get all users: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
}

func TestGetAllUsersEmpty(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	users, err := repo.GetAllUsers()
	if err != nil {
		t.Fatalf("failed to get all users: %v", err)
	}
	if len(users) != 0 {
		t.Errorf("expected 0 users, got %d", len(users))
	}
}

func TestUpdateUser(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "update-test",
		Name: "Original Name",
	}
	repo.CreateUser(user)

	user.Name = "Updated Name"
	err := repo.UpdateUser(user)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}

	found, _ := repo.GetUserByID("update-test")
	if found.Name != "Updated Name" {
		t.Errorf("expected 'Updated Name', got %s", found.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "delete-test",
		Name: "Delete Me",
	}
	repo.CreateUser(user)

	err := repo.DeleteUser("delete-test")
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	_, err = repo.GetUserByID("delete-test")
	if err == nil {
		t.Error("expected error after deleting user")
	}
}

func TestSetPassword(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "pw-test",
		Name: "Password Test",
	}
	repo.CreateUser(user)

	err := repo.SetPassword("pw-test", "secret123")
	if err != nil {
		t.Fatalf("failed to set password: %v", err)
	}

	valid, err := repo.ValidatePassword("pw-test", "secret123")
	if err != nil {
		t.Fatalf("failed to validate password: %v", err)
	}
	if !valid {
		t.Error("expected password to be valid")
	}
}

func TestValidatePasswordWrong(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "pw-wrong-test",
		Name: "Wrong Password",
	}
	repo.CreateUser(user)

	repo.SetPassword("pw-wrong-test", "correct")

	valid, err := repo.ValidatePassword("pw-wrong-test", "wrongpassword")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if valid {
		t.Error("expected password to be invalid")
	}
}

func TestAuthenticate(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "auth-test",
		Name: "Auth Test",
	}
	repo.CreateUser(user)
	repo.SetPassword("auth-test", "password123")

	authed, err := repo.Authenticate("Auth Test", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate: %v", err)
	}
	if authed.Id != "auth-test" {
		t.Errorf("expected authed user id 'auth-test', got %s", authed.Id)
	}
}

func TestAuthenticateByEmail(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:           "email-auth-test",
		Name:         "Email Auth",
		EmailAddress: "test@example.com",
	}
	repo.CreateUser(user)
	repo.SetPassword("email-auth-test", "password123")

	authed, err := repo.Authenticate("test@example.com", "password123")
	if err != nil {
		t.Fatalf("failed to authenticate by email: %v", err)
	}
	if authed.Id != "email-auth-test" {
		t.Errorf("expected authed user id 'email-auth-test', got %s", authed.Id)
	}
}

func TestAuthenticateInvalidCredentials(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "invalid-auth",
		Name: "Invalid Auth",
	}
	repo.CreateUser(user)
	repo.SetPassword("invalid-auth", "password123")

	_, err := repo.Authenticate("Invalid Auth", "wrongpassword")
	if err == nil {
		t.Error("expected error for invalid credentials")
	}
}

func TestAuthenticateNoPassword(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "no-pw-auth",
		Name: "No Password User",
	}
	repo.CreateUser(user)

	authed, err := repo.Authenticate("No Password User", "")
	if err != nil {
		t.Fatalf("failed to authenticate with no password: %v", err)
	}
	if authed.Id != "no-pw-auth" {
		t.Errorf("expected authed user id 'no-pw-auth', got %s", authed.Id)
	}
}

func TestUpdateLastLogin(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	user := &model.GORMUser{
		Id:   "lastlogin-test",
		Name: "Last Login Test",
	}
	repo.CreateUser(user)

	err := repo.UpdateLastLogin("lastlogin-test")
	if err != nil {
		t.Fatalf("failed to update last login: %v", err)
	}
}

func TestUserCount(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	repo.CreateUser(&model.GORMUser{Name: "User 1"})
	repo.CreateUser(&model.GORMUser{Name: "User 2"})

	count, err := repo.UserCount()
	if err != nil {
		t.Fatalf("failed to get user count: %v", err)
	}
	if count != 2 {
		t.Errorf("expected count 2, got %d", count)
	}
}

func TestUserCountEmpty(t *testing.T) {
	db, cleanup := setupUserTestDB(t)
	defer cleanup()

	repo := NewUserRepository(db)
	count, err := repo.UserCount()
	if err != nil {
		t.Fatalf("failed to get user count: %v", err)
	}
	if count != 0 {
		t.Errorf("expected count 0, got %d", count)
	}
}