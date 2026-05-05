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
func (r *UserRepository) CreateUserTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL UNIQUE,
		password_hash TEXT,
		primary_image_tag TEXT,
		has_configured_password INTEGER DEFAULT 0,
		has_configured_easy_password INTEGER DEFAULT 0,
		enable_auto_login INTEGER DEFAULT 0,
		last_login_date TEXT,
		created_date TEXT NOT NULL,
		connect_username TEXT,
		policy TEXT,
		configuration TEXT
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
	INSERT INTO users (id, name, password_hash, primary_image_tag, has_configured_password,
		has_configured_easy_password, enable_auto_login, last_login_date, created_date, connect_username)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	_, err := r.Exec(query, user.ID, user.Name, "", user.PrimaryImageTag,
		boolToInt(user.HasConfiguredPassword), boolToInt(user.HasConfiguredEasyPassword),
		boolToInt(user.EnableAutoLogin), user.LastLoginDate, user.CreatedDate, user.ConnectUserName)
	return err
}

// GetUserByID retrieves a user by ID.
func (r *UserRepository) GetUserByID(id string) (*User, error) {
	query := `SELECT id, name, primary_image_tag, has_configured_password, has_configured_easy_password,
		enable_auto_login, last_login_date, created_date, connect_username
		FROM users WHERE id = ?`

	user := &User{}
	var lastLogin, created sql.NullString
	err := r.QueryRow(query, id).Scan(
		&user.ID, &user.Name, &user.PrimaryImageTag, &user.HasConfiguredPassword,
		&user.HasConfiguredEasyPassword, &user.EnableAutoLogin, &lastLogin, &created, &user.ConnectUserName)
	if err != nil {
		return nil, err
	}
	if lastLogin.Valid {
		user.LastLoginDate, _ = time.Parse(time.RFC3339, lastLogin.String)
	}
	if created.Valid {
		user.CreatedDate, _ = time.Parse(time.RFC3339, created.String)
	}
	return user, nil
}

// GetUserByName retrieves a user by name.
func (r *UserRepository) GetUserByName(name string) (*User, error) {
	query := `SELECT id, name, primary_image_tag, has_configured_password, has_configured_easy_password,
		enable_auto_login, last_login_date, created_date, connect_username
		FROM users WHERE name = ?`

	user := &User{}
	var lastLogin, created sql.NullString
	err := r.QueryRow(query, name).Scan(
		&user.ID, &user.Name, &user.PrimaryImageTag, &user.HasConfiguredPassword,
		&user.HasConfiguredEasyPassword, &user.EnableAutoLogin, &lastLogin, &created, &user.ConnectUserName)
	if err != nil {
		return nil, err
	}
	if lastLogin.Valid {
		user.LastLoginDate, _ = time.Parse(time.RFC3339, lastLogin.String)
	}
	if created.Valid {
		user.CreatedDate, _ = time.Parse(time.RFC3339, created.String)
	}
	return user, nil
}

// GetAllUsers retrieves all users.
func (r *UserRepository) GetAllUsers() ([]*User, error) {
	query := `SELECT id, name, primary_image_tag, has_configured_password, has_configured_easy_password,
		enable_auto_login, last_login_date, created_date, connect_username
		FROM users ORDER BY name`

	rows, err := r.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		var lastLogin, created sql.NullString
		err := rows.Scan(
			&user.ID, &user.Name, &user.PrimaryImageTag, &user.HasConfiguredPassword,
			&user.HasConfiguredEasyPassword, &user.EnableAutoLogin, &lastLogin, &created, &user.ConnectUserName)
		if err != nil {
			return nil, err
		}
		if lastLogin.Valid {
			user.LastLoginDate, _ = time.Parse(time.RFC3339, lastLogin.String)
		}
		if created.Valid {
			user.CreatedDate, _ = time.Parse(time.RFC3339, created.String)
		}
		users = append(users, user)
	}
	return users, nil
}

// UpdateUser updates an existing user.
func (r *UserRepository) UpdateUser(user *User) error {
	query := `UPDATE users SET name = ?, primary_image_tag = ?, has_configured_password = ?,
		has_configured_easy_password = ?, enable_auto_login = ?, last_login_date = ?, connect_username = ?
		WHERE id = ?`

	_, err := r.Exec(query, user.Name, user.PrimaryImageTag,
		boolToInt(user.HasConfiguredPassword), boolToInt(user.HasConfiguredEasyPassword),
		boolToInt(user.EnableAutoLogin), user.LastLoginDate.Format(time.RFC3339),
		user.ConnectUserName, user.ID)
	return err
}

// DeleteUser deletes a user by ID.
func (r *UserRepository) DeleteUser(id string) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.Exec(query, id)
	return err
}

// SetPassword sets the password hash for a user.
func (r *UserRepository) SetPassword(userID, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	query := `UPDATE users SET password_hash = ?, has_configured_password = 1 WHERE id = ?`
	_, err = r.Exec(query, string(hash), userID)
	return err
}

// ValidatePassword validates a user's password.
func (r *UserRepository) ValidatePassword(userID, password string) (bool, error) {
	query := `SELECT password_hash FROM users WHERE id = ?`
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
	query := `UPDATE users SET last_login_date = ? WHERE id = ?`
	_, err := r.Exec(query, time.Now().Format(time.RFC3339), userID)
	return err
}

// UserCount returns the total number of users.
func (r *UserRepository) UserCount() (int, error) {
	query := `SELECT COUNT(*) FROM users`
	var count int
	err := r.QueryRow(query).Scan(&count)
	return count, err
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
