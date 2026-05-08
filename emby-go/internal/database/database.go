package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/emby/emby-go/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(cfg *config.DatabaseConfig) (*Manager, error) {
	dir := filepath.Dir(cfg.Path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("create database dir: %w", err)
		}
	}

	if info, err := os.Stat(cfg.Path); err == nil && info.IsDir() {
		if err := os.RemoveAll(cfg.Path); err != nil {
			return nil, fmt.Errorf("remove stale database directory: %w", err)
		}
	}

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	}

	db, err := gorm.Open(sqlite.Open(cfg.Path), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying db: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	if cfg.EnableWAL {
		sqlDB.Exec("PRAGMA journal_mode=WAL")
	}

	sqlDB.Exec("PRAGMA busy_timeout=5000")
	sqlDB.Exec("PRAGMA foreign_keys=ON")

	return &Manager{db: db}, nil
}

func (m *Manager) DB() *gorm.DB {
	return m.db
}

func (m *Manager) SQLDB() (*sql.DB, error) {
	return m.db.DB()
}

func (m *Manager) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (m *Manager) Ping() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}