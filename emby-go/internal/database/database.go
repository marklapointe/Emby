package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/emby/emby-go/internal/config"
)

// Manager handles database connections and initialization.
type Manager struct {
	db *sql.DB
}

// NewManager creates a new database manager with the given config.
func NewManager(cfg *config.DatabaseConfig) (*Manager, error) {
	// Ensure the directory exists
	dir := filepath.Dir(cfg.Path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("create database dir: %w", err)
		}
	}

	// Remove any existing directory with the same name as the database file
	// (fixes previous bug where MkdirAll created a directory instead of a file)
	if info, err := os.Stat(cfg.Path); err == nil && info.IsDir() {
		if err := os.RemoveAll(cfg.Path); err != nil {
			return nil, fmt.Errorf("remove stale database directory: %w", err)
		}
	}

	// Open SQLite connection
	db, err := sql.Open("sqlite", cfg.Path)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(0) // seconds, 0 = unlimited

	// Enable WAL mode if configured
	if cfg.EnableWAL {
		if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
			return nil, fmt.Errorf("enable WAL: %w", err)
		}
	}

	// Set busy timeout (milliseconds)
	if _, err := db.Exec("PRAGMA busy_timeout=5000"); err != nil {
		return nil, fmt.Errorf("set busy_timeout: %w", err)
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys=ON"); err != nil {
		return nil, fmt.Errorf("enable foreign keys: %w", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &Manager{db: db}, nil
}

// DB returns the underlying *sql.DB.
func (m *Manager) DB() *sql.DB {
	return m.db
}

// Close closes the database connection.
func (m *Manager) Close() error {
	return m.db.Close()
}

// Ping checks if the database is reachable.
func (m *Manager) Ping() error {
	return m.db.Ping()
}
