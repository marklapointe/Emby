package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/service/scheduled"
	"github.com/gorilla/mux"
)

// SystemHandler handles system-related API endpoints.
type SystemHandler struct {
	config       *config.Config
	db           *database.Manager
	scheduledSvc *scheduled.Manager
}

// NewSystemHandler creates a new system handler.
func NewSystemHandler(cfg *config.Config, db *database.Manager, scheduledSvc *scheduled.Manager) *SystemHandler {
	return &SystemHandler{
		config:       cfg,
		db:           db,
		scheduledSvc: scheduledSvc,
	}
}

// GetInfo handles GET /System/Info
func (h *SystemHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	info := map[string]interface{}{
		"Id":                    "emby-go-server-id",
		"ProductName":           "Emby Server",
		"Version":               "0.1.0",
		"OperatingSystem":       runtime.GOOS,
		"OperatingSystemVersion": "unknown",
		"OperatingSystemArchitecture": runtime.GOARCH,
		"OsPackageFamily":     runtime.GOOS,
		"CanSystemRestart":    true,
		"WanAccess":           true,
		"StartupWizardCompleted": true,
		"ConfigDir":           "data",
		"CacheDir":            "data/cache",
		"LogDir":              h.config.Logging.Level,
		"HasPendingRestart":   false,
		"IsInShutDown":        false,
		"IsShuttingDown":      false,
		"ServerName":          "Emby Server",
		"LocalAddress":        fmt.Sprintf("http://localhost:%d", h.config.Server.Port),
		"RemoteAddress":       "unknown",
		"WebSocketPortNumber": h.config.Server.Port,
		"HttpsPortNumber":     8920,
		"SocksProxyAddress":   "",
		"HttpServerPortNumber": h.config.Server.Port,
		"SystemUpdateTime":    "2026-04-29T00:00:00Z",
		"SupportsMultistreaming": true,
		"PackageUrl":          "https://emby.media",
		"EnableUPnP":          false,
		"EnableDLNA":          false,
		"DefaultLayout":       "List",
		"IsNetworkEnabled":    true,
		"RequiresPortForwarding": false,
		"RemoteIP":            "unknown",
		"RemotePort":          0,
		"IsRemote":            false,
		"LogToStdout":         h.config.Logging.Format == "stdout",
		"MinLogLevel":         h.config.Logging.Level,
		"EnableCaseSensitiveId": true,
		"EnableHttps":         false,
		"HttpsPort":           8920,
		"VersionInfo": map[string]interface{}{
			"AssemblyVersion": "0.1.0.0",
			"FileVersion":     "0.1.0.0",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

// GetDailySchedule handles GET /System/Info/DailySchedule
func (h *SystemHandler) GetDailySchedule(w http.ResponseWriter, r *http.Request) {
	schedule := map[string]interface{}{
		"Items": []map[string]interface{}{
			{
				"Name":       "Library Scan",
				"StartHour":  3,
				"EndHour":    4,
				"IsCompleted": true,
			},
			{
				"Name":       "Metadata Refresh",
				"StartHour":  4,
				"EndHour":    5,
				"IsCompleted": true,
			},
			{
				"Name":       "Thumbnail Generation",
				"StartHour":  5,
				"EndHour":    6,
				"IsCompleted": true,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schedule)
}

// Restart handles POST /System/Restart
func (h *SystemHandler) Restart(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "restart initiated",
	})
}

// GetPackageInfo handles GET /System/PackageInfo/{os}/{arch}
func (h *SystemHandler) GetPackageInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	os := vars["os"]
	arch := vars["arch"]

	_ = os
	_ = arch

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
		"DownloadUrlFreeBSD": "https://emby.media",
		"DownloadUrlAndroid": "https://emby.media",
		"DownloadUrlIos":  "https://emby.media",
		"DownloadUrlAppleTv": "https://emby.media",
		"DownloadUrlChromecast": "https://emby.media",
		"DownloadUrlAndroidTV": "https://emby.media",
		"DownloadUrlRoku": "https://emby.media",
		"DownloadUrlSamsungTV": "https://emby.media",
		"DownloadUrlSonyTV": "https://emby.media",
		"DownloadUrlWebOS": "https://emby.media",
		"DownloadUrlXbox": "https://emby.media",
		"DownloadUrlPlayStation": "https://emby.media",
		"DownloadUrlNintendo": "https://emby.media",
		"DownloadUrlSwitch": "https://emby.media",
		"DownloadUrlOuya": "https://emby.media",
		"DownloadUrlFireTv": "https://emby.media",
		"DownloadUrlFireTvStick": "https://emby.media",
		"DownloadUrlShield": "https://emby.media",
		"DownloadUrlPlex": "https://emby.media",
		"DownloadUrlKodi": "https://emby.media",
		"DownloadUrlPlexConnect": "https://emby.media",
		"DownloadUrlPlexPy": "https://emby.media",
		"DownloadUrlPlexPy2": "https://emby.media",
		"DownloadUrlPlexPy3": "https://emby.media",
		"DownloadUrlPlexPy4": "https://emby.media",
		"DownloadUrlPlexPy5": "https://emby.media",
		"DownloadUrlPlexPy6": "https://emby.media",
		"DownloadUrlPlexPy7": "https://emby.media",
		"DownloadUrlPlexPy8": "https://emby.media",
		"DownloadUrlPlexPy9": "https://emby.media",
		"DownloadUrlPlexPy10": "https://emby.media",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(packageInfo)
}

// GetScheduledTasks handles GET /ScheduledTasks
func (h *SystemHandler) GetScheduledTasks(w http.ResponseWriter, r *http.Request) {
	tasks := []map[string]interface{}{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// ExecuteScheduledTask handles POST /ScheduledTasks/Execute/{id}
func (h *SystemHandler) ExecuteScheduledTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["id"]

	err := h.scheduledSvc.ExecuteTask(taskId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "task executed",
	})
}

// GetUsageUsage handles GET /System/Usage
func (h *SystemHandler) GetUsage(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	usage := map[string]interface{}{
		"MemoryUsage":     mem.Alloc,
		"TotalMemory":     mem.TotalAlloc,
		"NumGoroutines":   runtime.NumGoroutine(),
		"NumCPU":          runtime.NumCPU(),
		"Uptime":          "0s",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(usage)
}
