package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emby/emby-go/internal/config"
)

func TestNewBrandingHandler(t *testing.T) {
	cfg := &config.Config{}
	handler := NewBrandingHandler(cfg)
	if handler == nil {
		t.Fatal("expected non-nil handler")
	}
	if handler.config != cfg {
		t.Error("expected config to be set")
	}
}

func TestGetCss(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/Branding/Css", nil)
	w := httptest.NewRecorder()
	handler.GetCss(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "text/css" {
		t.Errorf("expected text/css, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetJson(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/Branding/Json", nil)
	w := httptest.NewRecorder()
	handler.GetJson(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetImage(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/Branding/Images/test", nil)
	w := httptest.NewRecorder()
	handler.GetImage(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "image/png" {
		t.Errorf("expected image/png, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetBrandingOptions(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding", nil)
	w := httptest.NewRecorder()
	handler.GetBrandingOptions(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetBrandingLogo(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding/logo", nil)
	w := httptest.NewRecorder()
	handler.GetBrandingLogo(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "image/svg+xml" {
		t.Errorf("expected image/svg+xml, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetBrandingBanner(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding/banner", nil)
	w := httptest.NewRecorder()
	handler.GetBrandingBanner(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "image/svg+xml" {
		t.Errorf("expected image/svg+xml, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetBrandingFavicon(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding/favicon", nil)
	w := httptest.NewRecorder()
	handler.GetBrandingFavicon(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "image/svg+xml" {
		t.Errorf("expected image/svg+xml, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetBrandingTheme(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding/theme", nil)
	w := httptest.NewRecorder()
	handler.GetBrandingTheme(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetBrandingCSS(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding/css", nil)
	w := httptest.NewRecorder()
	handler.GetBrandingCSS(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "text/css" {
		t.Errorf("expected text/css, got %s", w.Header().Get("Content-Type"))
	}
}

func TestGetConfiguration(t *testing.T) {
	handler := NewBrandingHandler(&config.Config{})
	req := httptest.NewRequest(http.MethodGet, "/branding/configuration", nil)
	w := httptest.NewRecorder()
	handler.GetConfiguration(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected application/json, got %s", w.Header().Get("Content-Type"))
	}
}