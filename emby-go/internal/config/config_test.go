package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testAPIKey() string {
	return os.Getenv("EMBY_API_KEY")
}

func skipIfNoAPIKey(t *testing.T) {
	key := testAPIKey()
	if key == "" {
		t.Skip("EMBY_API_KEY not set - skipping API key integration test")
	}
}

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

func TestGetConfigDirs(t *testing.T) {
	oldHome := os.Getenv("HOME")
	oldXDG := os.Getenv("XDG_CONFIG_HOME")
	defer func() {
		os.Setenv("HOME", oldHome)
		os.Setenv("XDG_CONFIG_HOME", oldXDG)
	}()

	os.Unsetenv("XDG_CONFIG_HOME")
	os.Setenv("HOME", "/home/testuser")

	dirs := GetConfigDirs()

	if len(dirs) < 4 {
		t.Fatalf("expected at least 4 dirs, got %d", len(dirs))
	}

	foundHomeBased := false
	for _, dir := range dirs {
		if dir == "/home/testuser/.local/etc/emby" {
			foundHomeBased = true
		}
	}
	if !foundHomeBased {
		t.Error("expected to find $HOME/.local/etc/emby")
	}
}

func TestGetConfigDirsWithXDG(t *testing.T) {
	oldHome := os.Getenv("HOME")
	oldXDG := os.Getenv("XDG_CONFIG_HOME")
	defer func() {
		os.Setenv("HOME", oldHome)
		os.Setenv("XDG_CONFIG_HOME", oldXDG)
	}()

	os.Setenv("HOME", "/home/testuser")
	os.Setenv("XDG_CONFIG_HOME", "/home/testuser/.config")

	dirs := GetConfigDirs()

	foundXDG := false
	for _, dir := range dirs {
		if dir == "/home/testuser/.config/emby" {
			foundXDG = true
		}
	}
	if !foundXDG {
		t.Error("expected to find XDG_CONFIG_HOME/emby")
	}
}

func TestGetConfigFiles(t *testing.T) {
	files := GetConfigFiles()

	if len(files) != 8 {
		t.Errorf("expected 8 config files, got %d", len(files))
	}

	fileNames := make(map[string]bool)
	for _, f := range files {
		fileNames[f.Path] = true
	}

	expected := []string{"emby.yaml", "server.yaml", "database.yaml", "library.yaml",
		"logging.yaml", "security.yaml", "stream.yaml", "wizard.yaml"}
	for _, name := range expected {
		if !fileNames[name] {
			t.Errorf("expected to find %s in config files", name)
		}
	}
}

func TestSaveAndLoadSecurityConfig(t *testing.T) {
	cfg := DefaultConfig()
	testAPIKey := "my-secret-api-key"
	cfg.SetAPIKey(testAPIKey)

	tmpDir, err := os.MkdirTemp("", "emby-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	securityPath := filepath.Join(tmpDir, "security.yaml")

	if err := cfg.SaveSecurityConfig(securityPath); err != nil {
		t.Fatalf("failed to save security config: %v", err)
	}

	data, err := os.ReadFile(securityPath)
	if err != nil {
		t.Fatalf("failed to read saved security config: %v", err)
	}

	if !strings.Contains(string(data), testAPIKey) {
		t.Error("saved security config should contain the API key")
	}

	info, err := os.Stat(securityPath)
	if err != nil {
		t.Fatalf("failed to stat security file: %v", err)
	}

	if info.Mode().Perm() != 0600 {
		t.Errorf("expected file permissions 0600, got %o", info.Mode().Perm())
	}
}

func TestHasAPIKey(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.HasAPIKey() {
		t.Error("expected no API key initially")
	}

	cfg.SetAPIKey("test-key")
	if !cfg.HasAPIKey() {
		t.Error("expected API key to be set")
	}
}

func TestAPIKeyEnvOverride(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString("server:\n  host: \"localhost\"\n")
	tmpFile.Close()

	oldAPIKey := os.Getenv("EMBY_API_KEY")
	defer os.Setenv("EMBY_API_KEY", oldAPIKey)
	os.Setenv("EMBY_API_KEY", "env-api-key-123")

	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if cfg.Security.APIKey != "env-api-key-123" {
		t.Errorf("expected API key from env 'env-api-key-123', got '%s'", cfg.Security.APIKey)
	}
}

func TestDatabaseEnvOverrides(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	tmpFile.WriteString("server:\n  host: \"localhost\"\n")
	tmpFile.Close()

	oldType := os.Getenv("EMBY_DATABASE_TYPE")
	oldHost := os.Getenv("EMBY_DATABASE_HOST")
	oldPort := os.Getenv("EMBY_DATABASE_PORT")
	oldUser := os.Getenv("EMBY_DATABASE_USERNAME")
	oldPass := os.Getenv("EMBY_DATABASE_PASSWORD")
	oldName := os.Getenv("EMBY_DATABASE_NAME")
	defer func() {
		os.Setenv("EMBY_DATABASE_TYPE", oldType)
		os.Setenv("EMBY_DATABASE_HOST", oldHost)
		os.Setenv("EMBY_DATABASE_PORT", oldPort)
		os.Setenv("EMBY_DATABASE_USERNAME", oldUser)
		os.Setenv("EMBY_DATABASE_PASSWORD", oldPass)
		os.Setenv("EMBY_DATABASE_NAME", oldName)
	}()

	os.Setenv("EMBY_DATABASE_TYPE", "mysql")
	os.Setenv("EMBY_DATABASE_HOST", "db.example.com")
	os.Setenv("EMBY_DATABASE_PORT", "3307")
	os.Setenv("EMBY_DATABASE_USERNAME", "embyuser")
	os.Setenv("EMBY_DATABASE_PASSWORD", "embypass")
	os.Setenv("EMBY_DATABASE_NAME", "emby_production")

	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if cfg.Database.Type != "mysql" {
		t.Errorf("expected type 'mysql', got '%s'", cfg.Database.Type)
	}
	if cfg.Database.Host != "db.example.com" {
		t.Errorf("expected host 'db.example.com', got '%s'", cfg.Database.Host)
	}
	if cfg.Database.Port != 3307 {
		t.Errorf("expected port 3307, got %d", cfg.Database.Port)
	}
	if cfg.Database.Username != "embyuser" {
		t.Errorf("expected username 'embyuser', got '%s'", cfg.Database.Username)
	}
	if cfg.Database.Password != "embypass" {
		t.Errorf("expected password 'embypass', got '%s'", cfg.Database.Password)
	}
	if cfg.Database.Database != "emby_production" {
		t.Errorf("expected database 'emby_production', got '%s'", cfg.Database.Database)
	}
}

func TestGetConfigSearchPaths(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", "/home/testuser")

	paths := GetConfigSearchPaths()

	if len(paths) == 0 {
		t.Fatal("expected search paths to be returned")
	}

	for _, p := range paths {
		if p == "" {
			t.Error("search path should not be empty")
		}
	}
}

func TestExpandPath(t *testing.T) {
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", "/Users/john")

	tests := []struct {
		input    string
		expected string
	}{
		{"~/data", "/Users/john/data"},
		{"/absolute/path", "/absolute/path"},
		{"relative/path", "relative/path"},
	}

	for _, tt := range tests {
		result := ExpandPath(tt.input)
		if result != tt.expected {
			t.Errorf("ExpandPath(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}

func TestExternalDatabaseConfig(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	configContent := `
database:
  type: "postgres"
  host: "pg.cluster.local"
  port: 5432
  username: "emby"
  password: "secret"
  database: "emby_db"
  max_open_conns: 50
  max_idle_conns: 10
`
	tmpFile.WriteString(configContent)
	tmpFile.Close()

	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if cfg.Database.Type != "postgres" {
		t.Errorf("expected type 'postgres', got '%s'", cfg.Database.Type)
	}
	if cfg.Database.Host != "pg.cluster.local" {
		t.Errorf("expected host 'pg.cluster.local', got '%s'", cfg.Database.Host)
	}
	if cfg.Database.Port != 5432 {
		t.Errorf("expected port 5432, got %d", cfg.Database.Port)
	}
	if cfg.Database.Username != "emby" {
		t.Errorf("expected username 'emby', got '%s'", cfg.Database.Username)
	}
	if cfg.Database.Password != "secret" {
		t.Errorf("expected password 'secret', got '%s'", cfg.Database.Password)
	}
	if cfg.Database.Database != "emby_db" {
		t.Errorf("expected database 'emby_db', got '%s'", cfg.Database.Database)
	}
	if cfg.Database.MaxOpenConns != 50 {
		t.Errorf("expected max_open_conns 50, got %d", cfg.Database.MaxOpenConns)
	}
}

func TestMariaDBConfig(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	configContent := `
database:
  type: "mariadb"
  host: "maria.example.com"
  port: 3306
  username: "emby_user"
  password: "maria_pass"
  database: "emby_live"
`
	tmpFile.WriteString(configContent)
	tmpFile.Close()

	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if cfg.Database.Type != "mariadb" {
		t.Errorf("expected type 'mariadb', got '%s'", cfg.Database.Type)
	}
	if cfg.Database.Host != "maria.example.com" {
		t.Errorf("expected host 'maria.example.com', got '%s'", cfg.Database.Host)
	}
	if cfg.Database.Database != "emby_live" {
		t.Errorf("expected database 'emby_live', got '%s'", cfg.Database.Database)
	}
}

func TestLoadSecurityConfigWithAPIKey(t *testing.T) {
	skipIfNoAPIKey(t)

	apiKey := testAPIKey()
	t.Logf("Testing with API key: %s...", apiKey[:min(8, len(apiKey))]+"...")

	cfg := &SecurityConfig{}
	cfg.APIKey = apiKey

	if cfg.APIKey != apiKey {
		t.Errorf("expected API key '%s', got '%s'", apiKey, cfg.APIKey)
	}
}

func TestAPIKeyConfigSearchPaths(t *testing.T) {
	skipIfNoAPIKey(t)

	apiKey := testAPIKey()
	t.Logf("EMBY_API_KEY is set (%d chars) - testing API key search path resolution", len(apiKey))

	paths := GetConfigSearchPaths()
	if len(paths) == 0 {
		t.Fatal("expected search paths to be returned")
	}

	foundSecurityFile := false
	for _, p := range paths {
		if strings.Contains(p, "security.yaml") {
			foundSecurityFile = true
		}
	}
	if !foundSecurityFile {
		t.Error("expected security.yaml to be in search paths")
	}
}

func TestConfigDirPrecedence(t *testing.T) {
	oldHome := os.Getenv("HOME")
	oldXDG := os.Getenv("XDG_CONFIG_HOME")
	defer func() {
		os.Setenv("HOME", oldHome)
		os.Setenv("XDG_CONFIG_HOME", oldXDG)
	}()

	os.Setenv("HOME", "/home/testuser")
	os.Unsetenv("XDG_CONFIG_HOME")

	dirs := GetConfigDirs()

	if len(dirs) < 4 {
		t.Fatalf("expected at least 4 config dirs, got %d", len(dirs))
	}

	if dirs[0] != "/home/testuser/.local/etc/emby" {
		t.Errorf("expected first dir to be ~/.local/etc/emby, got %s", dirs[0])
	}

	if dirs[len(dirs)-1] != "configs" {
		t.Errorf("expected last dir to be 'configs', got %s", dirs[len(dirs)-1])
	}
}

func TestSecurityConfigFilePermissions(t *testing.T) {
	cfg := DefaultConfig()
	cfg.SetAPIKey("sensitive-key-12345")

	tmpDir, err := os.MkdirTemp("", "emby-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	securityPath := filepath.Join(tmpDir, "security.yaml")

	if err := cfg.SaveSecurityConfig(securityPath); err != nil {
		t.Fatalf("failed to save security config: %v", err)
	}

	info, err := os.Stat(securityPath)
	if err != nil {
		t.Fatalf("failed to stat security file: %v", err)
	}

	if info.Mode().Perm() != 0600 {
		t.Errorf("expected file permissions 0600, got %o", info.Mode().Perm())
	}
}

func TestMultipleConfigFilesOverride(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "emby-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	serverConfig := `
server:
  host: "192.168.1.100"
  port: 8097
`
	if err := os.WriteFile(filepath.Join(tmpDir, "server.yaml"), []byte(serverConfig), 0644); err != nil {
		t.Fatalf("failed to write server config: %v", err)
	}

	databaseConfig := `
database:
  type: "mysql"
  host: "dbserver"
`
	if err := os.WriteFile(filepath.Join(tmpDir, "database.yaml"), []byte(databaseConfig), 0644); err != nil {
		t.Fatalf("failed to write database config: %v", err)
	}

	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", oldHome)

	cfg := DefaultConfig()
	loadConfigFile(filepath.Join(tmpDir, "server.yaml"), cfg)
	loadConfigFile(filepath.Join(tmpDir, "database.yaml"), cfg)

	if cfg.Server.Host != "192.168.1.100" {
		t.Errorf("expected host 192.168.1.100, got %s", cfg.Server.Host)
	}
	if cfg.Server.Port != 8097 {
		t.Errorf("expected port 8097, got %d", cfg.Server.Port)
	}
	if cfg.Database.Type != "mysql" {
		t.Errorf("expected database type mysql, got %s", cfg.Database.Type)
	}
	if cfg.Database.Host != "dbserver" {
		t.Errorf("expected database host dbserver, got %s", cfg.Database.Host)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
