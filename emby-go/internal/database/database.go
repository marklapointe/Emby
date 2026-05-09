package database

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/emby/emby-go/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Manager struct {
	db *gorm.DB
}

func NewManager(cfg *config.DatabaseConfig) (*Manager, error) {
	var db *gorm.DB
	var err error

	switch cfg.Type {
	case "mysql", "mariadb":
		db, err = openMySQL(cfg)
	case "postgres", "postgresql":
		db, err = openPostgres(cfg)
	case "sqlite", "":
		db, err = openSQLite(cfg)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", cfg.Type)
	}

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("get underlying db: %w", err)
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)

	return &Manager{db: db}, nil
}

func openSQLite(cfg *config.DatabaseConfig) (*gorm.DB, error) {
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
		return nil, fmt.Errorf("open sqlite: %w", err)
	}

	sqlDB, _ := db.DB()
	if sqlDB != nil {
		if cfg.EnableWAL {
			sqlDB.Exec("PRAGMA journal_mode=WAL")
		}
		sqlDB.Exec("PRAGMA busy_timeout=5000")
		sqlDB.Exec("PRAGMA foreign_keys=ON")
	}

	return db, nil
}

func openMySQL(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := buildMySQLDSN(cfg)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("open mysql: %w", err)
	}
	return db, nil
}

func buildMySQLDSN(cfg *config.DatabaseConfig) string {
	if cfg.ConnectionString != "" {
		return cfg.ConnectionString
	}
	host := cfg.Host
	if host == "" {
		host = "localhost"
	}
	port := cfg.Port
	if port == 0 {
		port = 3306
	}
	user := cfg.Username
	if user == "" {
		user = "root"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, cfg.Password, host, port, cfg.Database)
}

func openPostgres(cfg *config.DatabaseConfig) (*gorm.DB, error) {
	dsn := buildPostgresDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	return db, nil
}

func buildPostgresDSN(cfg *config.DatabaseConfig) string {
	if cfg.ConnectionString != "" {
		return cfg.ConnectionString
	}
	host := cfg.Host
	if host == "" {
		host = "localhost"
	}
	port := cfg.Port
	if port == 0 {
		port = 5432
	}
	user := cfg.Username
	if user == "" {
		user = "postgres"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, cfg.Password, cfg.Database)
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