package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// PackageHandler handles package-related API endpoints.
type PackageHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewPackageHandler creates a new package handler.
func NewPackageHandler(cfg *config.Config, repo *repository.ItemRepository) *PackageHandler {
	return &PackageHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetPackages handles GET /Packages
func (h *PackageHandler) GetPackages(w http.ResponseWriter, r *http.Request) {
	packages := []map[string]interface{}{
		{
			"Name":        "Emby Server",
			"Version":     "0.1.0",
			"Overview":    "Emby Server Go Edition",
			"ProductId":   "emby-go-server",
			"ProductName": "Emby Server",
			"TargetRelease": "stable",
			"ReleaseDate": "2026-04-29",
			"DownloadUrl": "https://emby.media",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packages)
}

// GetPackage handles GET /Packages/{id}
func (h *PackageHandler) GetPackage(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("id")

	packageInfo := map[string]interface{}{
		"Name":        "Emby Server",
		"Version":     "0.1.0",
		"Overview":    "Emby Server Go Edition",
		"ProductId":   "emby-go-server",
		"ProductName": "Emby Server",
		"TargetRelease": "stable",
		"ReleaseDate": "2026-04-29",
		"DownloadUrl": "https://emby.media",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packageInfo)
}

// GetPackageInfo handles GET /System/PackageInfo/{os}/{arch}
func (h *PackageHandler) GetPackageInfo(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("os")
	_ = r.URL.Query().Get("arch")

	packageInfo := map[string]interface{}{
		"Name":            "Emby Server",
		"Description":     "Emby Server Go Edition",
		"Overview":        "A media server for organizing and streaming media",
		"ProductId":       "emby-go-server",
		"ProductName":     "Emby Server",
		"Version":         "0.1.0",
		"TargetRelease":   "stable",
		"ReleaseDate":     "2026-04-29",
		"DownloadUrl":     "https://emby.media",
		"DownloadUrlMac":  "https://emby.media",
		"DownloadUrlLinux": "https://emby.media",
		"DownloadUrlWindows": "https://emby.media",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packageInfo)
}

// GetPackageVersions handles GET /Packages/{id}/Versions
func (h *PackageHandler) GetPackageVersions(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("id")

	versions := []map[string]interface{}{
		{"Version": "0.1.0", "ReleaseDate": "2026-04-29", "IsLatest": true},
		{"Version": "0.0.9", "ReleaseDate": "2026-04-22", "IsLatest": false},
		{"Version": "0.0.8", "ReleaseDate": "2026-04-15", "IsLatest": false},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(versions)
}

// GetPackageChangelog handles GET /Packages/{id}/Changelog
func (h *PackageHandler) GetPackageChangelog(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("id")

	changelog := map[string]interface{}{
		"0.1.0": "Initial Go release",
		"0.0.9": "Bug fixes and improvements",
		"0.0.8": "DLNA support added",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(changelog)
}

// GetPackageReleaseNotes handles GET /Packages/{id}/ReleaseNotes
func (h *PackageHandler) GetPackageReleaseNotes(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("id")

	releaseNotes := map[string]interface{}{
		"0.1.0": "Initial Go release with core functionality.",
		"0.0.9": "Bug fixes and improvements.",
		"0.0.8": "DLNA support added.",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(releaseNotes)
}

// GetPackageDownloads handles GET /Packages/{id}/Downloads
func (h *PackageHandler) GetPackageDownloads(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("id")

	downloads := []map[string]interface{}{
		{"Url": "https://emby.media/download/linux", "FileName": "emby-server-linux.tar.gz", "Size": 50000000},
		{"Url": "https://emby.media/download/windows", "FileName": "emby-server-windows.zip", "Size": 55000000},
		{"Url": "https://emby.media/download/macos", "FileName": "emby-server-macos.dmg", "Size": 48000000},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(downloads)
}

// GetPackageChecksums handles GET /Packages/{id}/Checksums
func (h *PackageHandler) GetPackageChecksums(w http.ResponseWriter, r *http.Request) {
	_ = r.URL.Query().Get("id")

	checksums := map[string]interface{}{
		"sha256": "abc123def456",
		"md5":    "123abc456def",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(checksums)
}
