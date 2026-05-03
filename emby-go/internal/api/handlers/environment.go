package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/repository"
)

// EnvironmentHandler handles environment-related API endpoints.
type EnvironmentHandler struct {
	config *config.Config
	repo   *repository.ItemRepository
}

// NewEnvironmentHandler creates a new environment handler.
func NewEnvironmentHandler(cfg *config.Config, repo *repository.ItemRepository) *EnvironmentHandler {
	return &EnvironmentHandler{
		config: cfg,
		repo:   repo,
	}
}

// GetEnvironmentInfo handles GET /Environment
func (h *EnvironmentHandler) GetEnvironmentInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"OperatingSystem":    runtime.GOOS,
		"OperatingSystemVersion": "unknown",
		"OperatingSystemArchitecture": runtime.GOARCH,
		"ProcessorCount":     runtime.NumCPU(),
		"HasIPv6":            true,
		"HasHttp2":           true,
		"FileProtocol":       "file",
		"PackageOperatingSystem": runtime.GOOS,
		"SystemTimeZone":     getTimeZone(),
		"LocalAddress":       "http://localhost:" + string(rune(h.config.Server.Port)),
		"PublicAddress":      "http://localhost:" + string(rune(h.config.Server.Port)),
		"ServerName":         "Emby Server",
		"Version":            "0.1.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// GetEnvironmentPaths handles GET /Environment/Paths
func (h *EnvironmentHandler) GetEnvironmentPaths(w http.ResponseWriter, r *http.Request) {
	paths := map[string]interface{}{
		"ConfigDir":        "/etc/emby-server",
		"DataDir":          "/var/lib/emby-server",
		"CacheDir":         "/var/cache/emby-server",
		"LogDir":           "/var/log/emby-server",
		"MetadataDir":      "/var/lib/emby-server/metadata",
		"TranscodeDir":     "/var/lib/emby-server/transcode",
		"PluginConfigDir":  "/etc/emby-server/plugins",
		"PluginDataDir":    "/var/lib/emby-server/plugins",
		"CrashReportDir":   "/var/lib/emby-server/crashes",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paths)
}

// GetEnvironmentVariables handles GET /Environment/Variables
func (h *EnvironmentHandler) GetEnvironmentVariables(w http.ResponseWriter, r *http.Request) {
	vars := make(map[string]string)
	for _, env := range os.Environ() {
		// Filter out sensitive variables
		if !isSensitive(env) {
			pair := splitEnv(env)
			vars[pair[0]] = pair[1]
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vars)
}

// GetEnvironmentProcessInfo handles GET /Environment/ProcessInfo
func (h *EnvironmentHandler) GetEnvironmentProcessInfo(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	info := map[string]interface{}{
		"ProcessId":       os.Getpid(),
		"ProcessName":     os.Args[0],
		"WorkingDirectory": getWorkingDir(),
		"CommandLine":     os.Args,
		"MemoryUsage":     mem.Alloc,
		"TotalMemory":     mem.TotalAlloc,
		"NumGoroutines":   runtime.NumGoroutine(),
		"NumCPU":          runtime.NumCPU(),
		"NumGC":           0,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// GetEnvironmentNetworkInfo handles GET /Environment/NetworkInfo
func (h *EnvironmentHandler) GetEnvironmentNetworkInfo(w http.ResponseWriter, r *http.Request) {
	networkInfo := map[string]interface{}{
		"HasIPv4": true,
		"HasIPv6": true,
		"HasHttp2": true,
		"LocalIPv4Addresses": []string{"127.0.0.1", "192.168.1.1"},
		"LocalIPv6Addresses": []string{"::1", "fe80::1"},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(networkInfo)
}

// GetEnvironmentDiskInfo handles GET /Environment/DiskInfo
func (h *EnvironmentHandler) GetEnvironmentDiskInfo(w http.ResponseWriter, r *http.Request) {
	diskInfo := map[string]interface{}{
		"Drives": []map[string]interface{}{
			{
				"Name":        "/",
				"Path":        "/",
				"TotalSize":   1000000000,
				"FreeSpace":   500000000,
				"IsRemovable": false,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(diskInfo)
}

// isSensitive checks if an environment variable is sensitive.
func isSensitive(env string) bool {
	sensitive := []string{"PASSWORD", "SECRET", "TOKEN", "KEY", "CREDENTIAL"}
	for _, s := range sensitive {
		if len(env) > len(s) && env[len(env)-len(s):] == s {
			return true
		}
	}
	return false
}

// splitEnv splits an environment variable into key and value.
func splitEnv(env string) []string {
	for i := 0; i < len(env); i++ {
		if env[i] == '=' {
			return []string{env[:i], env[i+1:]}
		}
	}
	return []string{env, ""}
}

// getTimeZone returns the system time zone.
func getTimeZone() string {
	_, zone := os.LookupEnv("TZ")
	if zone {
		return os.Getenv("TZ")
	}
	return "UTC"
}

// getWorkingDir returns the current working directory.
func getWorkingDir() string {
	dir, _ := os.Getwd()
	return dir
}
