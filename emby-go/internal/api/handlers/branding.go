package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// BrandingHandler handles branding-related API endpoints.
type BrandingHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewBrandingHandler creates a new branding handler.
func NewBrandingHandler(cfg *config.Config, repo *repository.ItemRepository) *BrandingHandler {
	return &BrandingHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetBrandingOptions handles GET /branding
func (h *BrandingHandler) GetBrandingOptions(w http.ResponseWriter, r *http.Request) {
	options := map[string]interface{}{
		"CssOptions":    "",
		"EnableUpdateAvailable": false,
		"LogoImageVersion": 0,
		"BannerImageVersion": 0,
		"FaviconImageVersion": 0,
		"ThemeSet":          "default",
		"ThemeColors":       map[string]string{},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(options)
}

// GetBrandingLogo handles GET /branding/logo
func (h *BrandingHandler) GetBrandingLogo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" width="200" height="50"><text x="10" y="30" font-family="Arial" font-size="20" fill="white">Emby</text></svg>`))
}

// GetBrandingBanner handles GET /branding/banner
func (h *BrandingHandler) GetBrandingBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" width="1200" height="300"><rect width="1200" height="300" fill="#1a1a1a"/><text x="600" y="150" font-family="Arial" font-size="40" fill="white" text-anchor="middle">Emby Server</text></svg>`))
}

// GetBrandingFavicon handles GET /branding/favicon
func (h *BrandingHandler) GetBrandingFavicon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Write([]byte(`<svg xmlns="http://www.w3.org/2000/svg" width="32" height="32"><rect width="32" height="32" fill="#0078d7" rx="4"/><text x="16" y="22" font-family="Arial" font-size="16" fill="white" text-anchor="middle">E</text></svg>`))
}

// GetBrandingTheme handles GET /branding/theme
func (h *BrandingHandler) GetBrandingTheme(w http.ResponseWriter, r *http.Request) {
	theme := map[string]interface{}{
		"name":       "default",
		"color":      "#0078d7",
		"background": "#1a1a1a",
		"text":       "#ffffff",
		"accent":     "#00a4dc",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(theme)
}

// GetBrandingCSS handles GET /branding/css
func (h *BrandingHandler) GetBrandingCSS(w http.ResponseWriter, r *http.Request) {
	css := `
:root {
    --primary-color: #0078d7;
    --background-color: #1a1a1a;
    --text-color: #ffffff;
    --accent-color: #00a4dc;
}

body {
    background-color: var(--background-color);
    color: var(--text-color);
}

.emby-button {
    background-color: var(--primary-color);
    color: var(--text-color);
}
`
	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(css))
}
