package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/config"
	"go.uber.org/zap"
)

func TestNewConfigHandler(t *testing.T) {
	cfg := &config.Config{}
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
	if handler.config != cfg {
		t.Error("expected config to be set")
	}
	if handler.logger != logger {
		t.Error("expected logger to be set")
	}
}

func TestConfigHandlerGetConfiguration(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/Configuration", nil)
	w := httptest.NewRecorder()
	handler.GetConfiguration(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestUpdateConfiguration(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	body := []byte(`{"key": "value"}`)
	req := httptest.NewRequest(http.MethodPut, "/Configuration", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.UpdateConfiguration(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestUpdateConfigurationInvalidJSON(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	body := []byte(`invalid json`)
	req := httptest.NewRequest(http.MethodPut, "/Configuration", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.UpdateConfiguration(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestGetNamedConfiguration(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/Configuration/test", nil)
	w := httptest.NewRecorder()
	handler.GetNamedConfiguration(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetSystemConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/System/Configuration", nil)
	w := httptest.NewRecorder()
	handler.GetSystemConfig(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestUpdateSystemConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	body := []byte(`{"ServerName": "Test", "Port": 8097}`)
	req := httptest.NewRequest(http.MethodPost, "/System/Configuration", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.UpdateSystemConfig(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}

func TestGetPublicSystemConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/System/Public/Configuration", nil)
	w := httptest.NewRecorder()
	handler.GetPublicSystemConfig(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetLocalAddress(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/System/Info/LocalAddress", nil)
	w := httptest.NewRecorder()
	handler.GetLocalAddress(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetMacAddress(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/System/Info/MacAddress", nil)
	w := httptest.NewRecorder()
	handler.GetMacAddress(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetPluginConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	req := httptest.NewRequest(http.MethodGet, "/Plugins/test-plugin/Configuration", nil)
	w := httptest.NewRecorder()
	handler.GetPluginConfig(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestUpdatePluginConfig(t *testing.T) {
	cfg := config.DefaultConfig()
	logger := zap.NewNop()
	handler := NewConfigHandler(cfg, logger)
	body := []byte(`{"key": "value"}`)
	req := httptest.NewRequest(http.MethodPost, "/Plugins/test-plugin/Configuration", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handler.UpdatePluginConfig(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
}