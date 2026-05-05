package database

import (
	"os"
	"testing"

	"github.com/emby/emby-go/internal/config"
	_ "modernc.org/sqlite"
)

func TestNewManager(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test-*.db")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cfg := &config.DatabaseConfig{
		Path:         tmpfile.Name(),
		MaxOpenConns: 5,
		MaxIdleConns: 2,
		EnableWAL:    true,
	}

	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatalf("NewManager returned error: %v", err)
	}
	if mgr == nil {
		t.Fatal("NewManager returned nil")
	}
	defer mgr.Close()
}

func TestNewManager_CreateDir(t *testing.T) {
	tmpfile := "/tmp/test-nested-dir/subdir/test.db"
	defer os.RemoveAll("/tmp/test-nested-dir")

	cfg := &config.DatabaseConfig{
		Path:         tmpfile,
		MaxOpenConns: 5,
		MaxIdleConns: 2,
	}

	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatalf("NewManager returned error: %v", err)
	}
	defer mgr.Close()
}

func TestDB(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test-*.db")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cfg := &config.DatabaseConfig{
		Path:         tmpfile.Name(),
		MaxOpenConns: 5,
		MaxIdleConns: 2,
	}

	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatalf("NewManager returned error: %v", err)
	}
	defer mgr.Close()

	db := mgr.DB()
	if db == nil {
		t.Error("DB() returned nil")
	}
}

func TestClose(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test-*.db")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cfg := &config.DatabaseConfig{
		Path:         tmpfile.Name(),
		MaxOpenConns: 5,
		MaxIdleConns: 2,
	}

	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatalf("NewManager returned error: %v", err)
	}

	err = mgr.Close()
	if err != nil {
		t.Errorf("Close returned error: %v", err)
	}
}

func TestPing(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test-*.db")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpfile.Close()
	defer os.Remove(tmpfile.Name())

	cfg := &config.DatabaseConfig{
		Path:         tmpfile.Name(),
		MaxOpenConns: 5,
		MaxIdleConns: 2,
	}

	mgr, err := NewManager(cfg)
	if err != nil {
		t.Fatalf("NewManager returned error: %v", err)
	}
	defer mgr.Close()

	err = mgr.Ping()
	if err != nil {
		t.Errorf("Ping returned error: %v", err)
	}
}

func TestNewManager_InvalidPath(t *testing.T) {
	cfg := &config.DatabaseConfig{
		Path:         "/invalid/path/that/does/not/exist/test.db",
		MaxOpenConns: 5,
		MaxIdleConns: 2,
	}

	_, err := NewManager(cfg)
	if err == nil {
		t.Error("expected error for invalid path")
	}
}