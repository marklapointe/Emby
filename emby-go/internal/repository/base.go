package repository

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

// BaseRepository provides common SQLite operations.
type BaseRepository struct {
	db *sql.DB
}

// NewBaseRepository creates a new base repository with the given database connection.
func NewBaseRepository(db *sql.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

// DB returns the underlying *sql.DB for direct access when needed.
func (r *BaseRepository) DB() *sql.DB {
	return r.db
}

// Query executes a SELECT query and returns rows.
func (r *BaseRepository) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return r.db.Query(query, args...)
}

// QueryRow executes a SELECT query that returns a single row.
func (r *BaseRepository) QueryRow(query string, args ...interface{}) *sql.Row {
	return r.db.QueryRow(query, args...)
}

// Exec executes a non-SELECT query.
func (r *BaseRepository) Exec(query string, args ...interface{}) (sql.Result, error) {
	return r.db.Exec(query, args...)
}

// WithTransaction executes a function within a database transaction.
func (r *BaseRepository) WithTransaction(fn func(*sql.Tx) error) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		return err
	}

	return tx.Commit()
}

// Ping checks if the database connection is alive.
func (r *BaseRepository) Ping() error {
	return r.db.Ping()
}
