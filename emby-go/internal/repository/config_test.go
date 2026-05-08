package repository

import (
	"testing"
)

func TestServerConfigTableName(t *testing.T) {
	cfg := ServerConfig{}
	if cfg.TableName() != "server_config" {
		t.Errorf("expected 'server_config', got %s", cfg.TableName())
	}
}

func TestNewConfigRepository(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	if repo == nil {
		t.Fatal("expected non-nil repository")
	}
}

func TestConfigRepository_CreateConfigTable(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	err = repo.CreateConfigTable()
	if err != nil {
		t.Fatalf("failed to create config table: %v", err)
	}
}

func TestConfigRepository_GetConfig(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	repo.CreateConfigTable()

	config, err := repo.GetConfig()
	if err != nil {
		t.Fatalf("failed to get config: %v", err)
	}
	if config == nil {
		t.Fatal("expected non-nil config")
	}
	if config.ServerName != "Emby Server" {
		t.Errorf("expected 'Emby Server', got %s", config.ServerName)
	}
}

func TestConfigRepository_GetConfigNotFound(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	repo.CreateConfigTable()

	config, err := repo.GetConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if config.ServerName != "Emby Server" {
		t.Errorf("expected default ServerName, got %s", config.ServerName)
	}
}

func TestConfigRepository_SaveConfig(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	repo.CreateConfigTable()

	config := &ServerConfig{
		ServerName: "Test Server",
		HttpPort:   8097,
	}
	err = repo.SaveConfig(config)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	saved, err := repo.GetConfig()
	if err != nil {
		t.Fatalf("failed to get saved config: %v", err)
	}
	if saved.ServerName != "Test Server" {
		t.Errorf("expected 'Test Server', got %s", saved.ServerName)
	}
	if saved.HttpPort != 8097 {
		t.Errorf("expected port 8097, got %d", saved.HttpPort)
	}
}

func TestConfigRepository_GetString(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	repo.CreateConfigTable()
	repo.SaveConfig(&ServerConfig{ServerName: "StringTest"})

	val, err := repo.GetString("ServerName")
	if err != nil {
		t.Fatalf("failed to get string: %v", err)
	}
	if val != "StringTest" {
		t.Errorf("expected 'StringTest', got %s", val)
	}
}

func TestConfigRepository_GetStringUnknownKey(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)
	repo.CreateConfigTable()

	_, err = repo.GetString("UnknownKey")
	if err == nil {
		t.Error("expected error for unknown key")
	}
}

func TestConfigRepository_GetDefaultConfig(t *testing.T) {
	db, err := createMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}
	repo := NewConfigRepository(db)

	config := repo.GetDefaultConfig()
	if config.ServerName != "Emby Server" {
		t.Errorf("expected 'Emby Server', got %s", config.ServerName)
	}
	if config.HttpPort != 8096 {
		t.Errorf("expected 8096, got %d", config.HttpPort)
	}
	if config.PreferredMetadataLanguage != "en" {
		t.Errorf("expected 'en', got %s", config.PreferredMetadataLanguage)
	}
}

func TestYAMLToConfigInvalid(t *testing.T) {
	_, err := YAMLToConfig([]byte("invalid: yaml: :data:"))
	if err == nil {
		t.Error("expected error for invalid yaml")
	}
}