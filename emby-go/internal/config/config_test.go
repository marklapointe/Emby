package config

import (
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	
	if cfg.Server.Host != "0.0.0.0" {
		t.Errorf("expected host 0.0.0.0, got %s", cfg.Server.Host)
	}
	
	if cfg.Server.Port != 8096 {
		t.Errorf("expected port 8096, got %d", cfg.Server.Port)
	}
	
	if cfg.Database.Path != "data/emby-server.db" {
		t.Errorf("expected database path data/emby-server.db, got %s", cfg.Database.Path)
	}
}

func TestLoadConfig(t *testing.T) {
	// Create a temp config file
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	
	configContent := `
server:
  host: "127.0.0.1"
  port: 8080
database:
  path: "/tmp/test.db"
`
	if _, err := tmpFile.WriteString(configContent); err != nil {
		t.Fatalf("failed to write config: %v", err)
	}
	tmpFile.Close()
	
	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}
	
	if cfg.Server.Host != "127.0.0.1" {
		t.Errorf("expected host 127.0.0.1, got %s", cfg.Server.Host)
	}
	
	if cfg.Server.Port != 8080 {
		t.Errorf("expected port 8080, got %d", cfg.Server.Port)
	}
	
	if cfg.Database.Path != "/tmp/test.db" {
		t.Errorf("expected database path /tmp/test.db, got %s", cfg.Database.Path)
	}
}

func TestSaveConfig(t *testing.T) {
	cfg := DefaultConfig()
	cfg.Server.Host = "192.168.1.1"
	cfg.Server.Port = 9000
	
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	tmpFile.Close()
	defer os.Remove(tmpFile.Name())
	
	if err := cfg.SaveConfig(tmpFile.Name()); err != nil {
		t.Fatalf("failed to save config: %v", err)
	}
	
	loadedCfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load saved config: %v", err)
	}
	
	if loadedCfg.Server.Host != "192.168.1.1" {
		t.Errorf("expected host 192.168.1.1, got %s", loadedCfg.Server.Host)
	}
	
	if loadedCfg.Server.Port != 9000 {
		t.Errorf("expected port 9000, got %d", loadedCfg.Server.Port)
	}
}
