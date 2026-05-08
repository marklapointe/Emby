package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
)

func TestNewLocalizationHandler(t *testing.T) {
	logger := zap.NewNop()
	handler := NewLocalizationHandler(logger)
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
	if handler.logger != logger {
		t.Error("expected logger to be set")
	}
}

func TestGetLocalization(t *testing.T) {
	logger := zap.NewNop()
	handler := NewLocalizationHandler(logger)
	req := httptest.NewRequest(http.MethodGet, "/Localization/en", nil)
	w := httptest.NewRecorder()
	handler.GetLocalization(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestLocalizationGetCultures(t *testing.T) {
	logger := zap.NewNop()
	handler := NewLocalizationHandler(logger)
	req := httptest.NewRequest(http.MethodGet, "/Localization/Cultures", nil)
	w := httptest.NewRecorder()
	handler.GetCultures(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestLocalizationGetCountries(t *testing.T) {
	logger := zap.NewNop()
	handler := NewLocalizationHandler(logger)
	req := httptest.NewRequest(http.MethodGet, "/Localization/Countries", nil)
	w := httptest.NewRecorder()
	handler.GetCountries(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}