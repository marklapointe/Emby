package repository

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// UserRepository handles user data persistence.
type UserRepository struct {
	*BaseRepository
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{NewBaseRepository(db)}
}

// User represents a user in the database.
type User struct {
	ID                string    `json:"Id"`
	Name              string    `json:"Name"`
	PrimaryImageTag   string    `json:"PrimaryImageTag,omitempty"`
	HasConfiguredPassword bool  `json:"HasConfiguredPassword"`
	HasConfiguredEasyPassword bool `json:"HasConfiguredEasyPassword"`
	EnableAutoLogin   bool      `json:"EnableAutoLogin"`
	LastLoginDate     time.Time `json:"LastLoginDate,omitempty"`
	CreatedDate       time.Time `json:"CreatedDate"`
	ConnectUserName   string    `json:"ConnectUserName,omitempty"`
}

// CreateUserTable creates the users table if it doesn't exist.
// Uses the Users table from ItemRepository schema for compatibility.
func (r *UserRepository) CreateUserTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS Users (
		Id TEXT PRIMARY KEY,
		Name TEXT NOT NULL,
		Username TEXT,
		EmailAddress TEXT,
		LoginUsername TEXT,
		LoginPassword TEXT,
		InvalidLoginAttemptCount INTEGER,
		LastLoginDate TEXT,
		LastActivityDate TEXT,
		AuthenticationProviderID TEXT,
		PrimaryImageTag TEXT,
		Policy TEXT
	);`
	_, err := r.Exec(query)
	return err
}

// CreateUser creates a new user.
func (r *UserRepository) CreateUser(user *User) error {
	if user.ID == "" {
		user.ID = fmt.Sprintf("user-%d", time.Now().UnixNano())
	}
	if user.CreatedDate.IsZero() {
		user.CreatedDate = time.Now()
	}

	query := `
	INSERT INTO Users (Id, Name, Username, EmailAddress, LoginUsername, LoginPassword,
		InvalidLoginAttemptCount, LastLoginDate, LastActivityDate, AuthenticationProviderID,
		PrimaryImageTag, Policy)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := r.Exec(query, user.ID, user.Name, "", "", "", "",
		0, user.LastLoginDate.Format(time.RFC3339), "", "", user.PrimaryImageTag, "")
	return err
}

// GetUserByID retrieves a user by ID.
func (r *UserRepository) GetUserByID(id string) (*User, error) {
	query := `SELECT Id, Name, Username, EmailAddress, LoginUsername, LoginPassword,
		InvalidLoginAttemptCount, LastLoginDate, LastActivityDate, AuthenticationProviderID,
		PrimaryImageTag, Policy
		FROM Users WHERE Id = ?`

	user := &User{}
	var lastLogin, lastActivity, authProviderID, policy sql.NullString
	var invalidLoginCount sql.NullInt64
	err := r.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.PrimaryImageTag,
		&invalidLoginCount, &lastLogin, &lastActivity, &authProviderID, &policy)
	if err != nil {
		return nil, err
	}
	if lastLogin.Valid {
		user.LastLoginDate, _ = time.Parse(time.RFC3339, lastLogin.String)
	}
	return user, nil
}

// GetUserByName retrieves a user by name.
func (r *UserRepository) GetUserByName(name string) (*User, error) {
	query := `SELECT Id, Name, PrimaryImageTag FROM Users WHERE Name = ?`

	user := &User{}
	err := r.QueryRow(query, name).Scan(&user.ID, &user.Name, &user.PrimaryImageTag)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAllUsers retrieves all users.
func (r *UserRepository) GetAllUsers() ([]*User, error) {
	query := `SELECT Id, Name, PrimaryImageTag FROM Users ORDER BY Name`

	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.Name, &user.PrimaryImageTag)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if users == nil {
		users = []*User{}
	}
	return users, nil
}

// UpdateUser updates an existing user.
func (r *UserRepository) UpdateUser(user *User) error {
	query := `UPDATE Users SET Name = ?, PrimaryImageTag = ? WHERE Id = ?`

	_, err := r.Exec(query, user.Name, user.PrimaryImageTag, user.ID)
	return err
}

// DeleteUser deletes a user by ID.
func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM Users WHERE Id = ?`
	_, err := r.Exec(query, id)
	return err
}

// SetPassword sets the password hash for a user.
func (r *UserRepository) SetPassword(userID, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	query := `UPDATE Users SET LoginPassword = ?, InvalidLoginAttemptCount = 0 WHERE Id = ?`
	_, err = r.Exec(query, string(hash), userID)
	return err
}

// ValidatePassword validates a user's password.
func (r *UserRepository) ValidatePassword(userID, password string) (bool, error) {
	query := `SELECT LoginPassword FROM Users WHERE Id = ?`
	var hash string
	err := r.QueryRow(query, userID).Scan(&hash)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateLastLogin updates the last login date for a user.
func (r *UserRepository) UpdateLastLogin(userID string) error {
	query := `UPDATE Users SET LastLoginDate = ? WHERE Id = ?`
	_, err := r.Exec(query, time.Now().Format(time.RFC3339), userID)
	return err
}

// UserCount returns the total number of users.
func (r *UserRepository) UserCount() (int, error) {
	query := `SELECT COUNT(*) FROM Users`
	var count int
	err := r.QueryRow(query).Scan(&count)
	return count, err
}
