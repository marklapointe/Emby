package api

import (
	"net/http"

	"github.com/emby/emby-go/internal/api/handlers"
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/service/device"
	"github.com/emby/emby-go/internal/service/image"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/media"
	"github.com/emby/emby-go/internal/service/notification"
	"github.com/emby/emby-go/internal/service/scheduled"
	"github.com/emby/emby-go/internal/service/session"
	"github.com/emby/emby-go/internal/service/transcoding"
	"github.com/emby/emby-go/internal/service/user"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Router manages API routes and handlers.
type Router struct {
	config      *config.Config
	logger      *zap.Logger
	dbManager   *database.Manager
	userSvc     *user.Manager
	librarySvc  *library.Manager
	mediaSvc    *media.Manager
	sessionSvc  *session.Manager
	deviceSvc   *device.Manager
	imageSvc    *image.Manager
	notificationSvc *notification.Manager
	scheduledSvc    *scheduled.Manager
	transcodingSvc   *transcoding.Manager
}

// NewRouter creates a new API router.
func NewRouter(cfg *config.Config, logger *zap.Logger, dbManager *database.Manager) *Router {
	return &Router{
		config:    cfg,
		logger:    logger,
		dbManager: dbManager,
	}
}

// RegisterRoutes registers all API routes on the given chi router.
func (r *Router) RegisterRoutes(router *chi.Mux) {
	// Initialize services
	r.userSvc = user.NewManager(r.dbManager, r.logger)
	r.librarySvc = library.NewManager(r.config, r.logger, r.dbManager)
	r.mediaSvc = media.NewManager(r.config, r.logger)
	r.sessionSvc = session.NewManager(r.config, r.logger)
	r.deviceSvc = device.NewManager(r.logger)
	r.imageSvc = image.NewManager(r.logger)
	r.notificationSvc = notification.NewManager()
	r.scheduledSvc = scheduled.NewManager(r.logger)
	r.transcodingSvc = transcoding.NewManager(r.config, r.logger)

	// Register routes on the provided router
	r.registerLibraryRoutes(router)
	r.registerSessionRoutes(router)
	r.registerUserRoutes(router)
	r.registerDeviceRoutes(router)
	r.registerImageRoutes(router)
	r.registerMediaRoutes(router)
	r.registerNotificationRoutes(router)
	r.registerScheduledTaskRoutes(router)
	r.registerTranscodingRoutes(router)

	// Register health endpoint
	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
}

// registerLibraryRoutes registers library-related routes.
func (r *Router) registerLibraryRoutes(router *chi.Mux) {
	libHandler := handlers.NewLibraryHandler(library.NewScanner(r.config, r.logger, nil))

	r.registerRoute(router, http.MethodGet, "/Library/Root", libHandler.GetLibraryRoot)
	r.registerRoute(router, http.MethodGet, "/Library/Items", libHandler.GetItems)
	r.registerRoute(router, http.MethodGet, "/Library/MediaFolders", libHandler.GetMediaFolders)
	r.registerRoute(router, http.MethodPost, "/Library/MediaFolders", libHandler.CreateMediaFolder)
	r.registerRoute(router, http.MethodGet, "/Library/MediaFolders/{id}", libHandler.GetMediaFolder)
	r.registerRoute(router, http.MethodDelete, "/Library/MediaFolders/{id}", libHandler.DeleteMediaFolder)
	r.registerRoute(router, http.MethodGet, "/Library/MediaFolders/{id}/Items", libHandler.GetFolderItems)
	r.registerRoute(router, http.MethodPost, "/Library/Folders/FullScan", libHandler.ScanLibrary)
	r.registerRoute(router, http.MethodGet, "/Library/VirtualFolders", libHandler.GetVirtualFolders)
	r.registerRoute(router, http.MethodGet, "/Library/VirtualFolders/{id}/Items", libHandler.GetVirtualFolderItems)
}

// registerSessionRoutes registers session-related routes.
func (r *Router) registerSessionRoutes(router *chi.Mux) {
	sessionHandler := handlers.NewSessionHandler(r.sessionSvc)

	r.registerRoute(router, http.MethodGet, "/Sessions", sessionHandler.GetSessions)
	r.registerRoute(router, http.MethodGet, "/Sessions/{id}", sessionHandler.GetSession)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/Playing", sessionHandler.StartPlayback)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/Playing/Progress", sessionHandler.PlaybackProgress)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/Playing/Stopped", sessionHandler.StopPlayback)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/Volume", sessionHandler.SetVolume)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/Pause", sessionHandler.PausePlayback)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/Unpause", sessionHandler.UnpausePlayback)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/ToggleFullscreen", sessionHandler.ToggleFullscreen)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/GoTo", sessionHandler.NavigateTo)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/SendKey", sessionHandler.SendKey)
	r.registerRoute(router, http.MethodPost, "/Sessions/{id}/SendText", sessionHandler.SendText)
	r.registerRoute(router, http.MethodDelete, "/Sessions/{id}", sessionHandler.CloseSession)
}

// registerUserRoutes registers user-related routes.
func (r *Router) registerUserRoutes(router *chi.Mux) {
	userHandler := handlers.NewUserHandler(r.userSvc)

	r.registerRoute(router, http.MethodGet, "/Users", userHandler.GetUsers)
	r.registerRoute(router, http.MethodGet, "/Users/Public", userHandler.GetPublicUsers)
	r.registerRoute(router, http.MethodPost, "/Users/Login", userHandler.Login)
	r.registerRoute(router, http.MethodPost, "/Users/Logout", userHandler.Logout)
	r.registerRoute(router, http.MethodGet, "/Users/{id}", userHandler.GetUser)
	r.registerRoute(router, http.MethodPut, "/Users/{id}", userHandler.UpdateUser)
	r.registerRoute(router, http.MethodDelete, "/Users/{id}", userHandler.DeleteUser)
	r.registerRoute(router, http.MethodPost, "/Users/{id}/Password", userHandler.ChangePassword)
	r.registerRoute(router, http.MethodGet, "/Users/{id}/Images/{type}", userHandler.GetUserImage)
	r.registerRoute(router, http.MethodGet, "/Users/{id}/Configuration", userHandler.GetUserConfiguration)
	r.registerRoute(router, http.MethodPut, "/Users/{id}/Configuration", userHandler.UpdateUserConfiguration)
	r.registerRoute(router, http.MethodGet, "/Users/{id}/Policy", userHandler.GetUserPolicy)
	r.registerRoute(router, http.MethodPut, "/Users/{id}/Policy", userHandler.UpdateUserPolicy)
	r.registerRoute(router, http.MethodGet, "/Users/Device/{deviceId}", userHandler.GetUsersByDevice)
	r.registerRoute(router, http.MethodGet, "/Users/LibraryFolders/{folderId}", userHandler.GetUsersByLibraryFolder)
}

// registerDeviceRoutes registers device-related routes.
func (r *Router) registerDeviceRoutes(router *chi.Mux) {
	deviceHandler := handlers.NewDeviceHandler(r.deviceSvc)

	r.registerRoute(router, http.MethodGet, "/Devices", deviceHandler.GetDevices)
	r.registerRoute(router, http.MethodGet, "/Devices/{id}", deviceHandler.GetDevice)
	r.registerRoute(router, http.MethodPut, "/Devices/{id}", deviceHandler.UpdateDevice)
	r.registerRoute(router, http.MethodDelete, "/Devices/{id}", deviceHandler.DeleteDevice)
	r.registerRoute(router, http.MethodGet, "/Devices/{id}/Icon", deviceHandler.GetDeviceIcon)
	r.registerRoute(router, http.MethodGet, "/Devices/{id}/Profile", deviceHandler.GetDeviceProfile)
}

// registerImageRoutes registers image-related routes.
func (r *Router) registerImageRoutes(router *chi.Mux) {
	imageHandler := handlers.NewImageHandler(r.imageSvc)

	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}", imageHandler.GetItemImage)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}/BlurHash", imageHandler.GetItemImageBlurHash)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}/{index}", imageHandler.GetItemImageByIndex)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}/Tag/{tag}", imageHandler.GetItemImageTag)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}/Crop", imageHandler.GetItemImageCrop)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}/Resize", imageHandler.GetItemImageResize)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Images/{type}/Rotate", imageHandler.GetItemImageRotation)
}

// registerMediaRoutes registers media-related routes.
func (r *Router) registerMediaRoutes(router *chi.Mux) {
	mediaHandler := handlers.NewMediaHandler(r.mediaSvc)

	r.registerRoute(router, http.MethodGet, "/Items/{id}", mediaHandler.GetItem)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Stream", mediaHandler.GetStream)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Subtitles", mediaHandler.GetSubtitles)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Subtitles/{index}/Stream", mediaHandler.GetSubtitleStream)
	r.registerRoute(router, http.MethodGet, "/Items/{id}/Audio", mediaHandler.GetAudioStream)
}

// registerNotificationRoutes registers notification-related routes.
func (r *Router) registerNotificationRoutes(router *chi.Mux) {
	notificationHandler := handlers.NewNotificationHandler(r.notificationSvc)

	r.registerRoute(router, http.MethodGet, "/Notifications", notificationHandler.GetNotifications)
	r.registerRoute(router, http.MethodGet, "/Notifications/Unread", notificationHandler.GetUnreadNotifications)
	r.registerRoute(router, http.MethodPost, "/Notifications/{id}/MarkRead", notificationHandler.MarkAsRead)
	r.registerRoute(router, http.MethodPost, "/Notifications/MarkAllRead", notificationHandler.MarkAllAsRead)
	r.registerRoute(router, http.MethodDelete, "/Notifications/{id}", notificationHandler.DeleteNotification)
	r.registerRoute(router, http.MethodGet, "/Notifications/Count", notificationHandler.GetNotificationCount)
	r.registerRoute(router, http.MethodGet, "/Notifications/UnreadCount", notificationHandler.GetUnreadNotificationCount)
}

// registerScheduledTaskRoutes registers scheduled task-related routes.
func (r *Router) registerScheduledTaskRoutes(router *chi.Mux) {
	scheduledHandler := handlers.NewScheduledTaskHandler(r.scheduledSvc)

	r.registerRoute(router, http.MethodGet, "/ScheduledTasks", scheduledHandler.GetAllTasks)
	r.registerRoute(router, http.MethodGet, "/ScheduledTasks/Running", scheduledHandler.GetRunningTasks)
	r.registerRoute(router, http.MethodGet, "/ScheduledTasks/{id}", scheduledHandler.GetTask)
	r.registerRoute(router, http.MethodPost, "/ScheduledTasks/{id}/Execute", scheduledHandler.ExecuteTask)
	r.registerRoute(router, http.MethodPost, "/ScheduledTasks/{id}/Cancel", scheduledHandler.CancelTask)
	r.registerRoute(router, http.MethodGet, "/ScheduledTasks/Count", scheduledHandler.GetTaskCount)
	r.registerRoute(router, http.MethodGet, "/ScheduledTasks/RunningCount", scheduledHandler.GetRunningTaskCount)
}

// registerTranscodingRoutes registers transcoding-related routes.
func (r *Router) registerTranscodingRoutes(router *chi.Mux) {
	transcodingHandler := handlers.NewTranscodingHandler(r.transcodingSvc)

	r.registerRoute(router, http.MethodGet, "/TranscodingProfiles", transcodingHandler.GetTranscodingProfiles)
	r.registerRoute(router, http.MethodGet, "/TranscodingProfiles/{id}", transcodingHandler.GetTranscodingProfile)
	r.registerRoute(router, http.MethodGet, "/ActiveTranscodes", transcodingHandler.GetActiveTranscodes)
	r.registerRoute(router, http.MethodGet, "/ActiveTranscodes/{id}", transcodingHandler.GetActiveTranscode)
	r.registerRoute(router, http.MethodPost, "/ActiveTranscodes/{id}/Stop", transcodingHandler.StopTranscode)
}

// registerRoute registers a single route on the chi router.
func (r *Router) registerRoute(router *chi.Mux, method, path string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		router.Get(path, handler)
	case http.MethodPost:
		router.Post(path, handler)
	case http.MethodPut:
		router.Put(path, handler)
	case http.MethodDelete:
		router.Delete(path, handler)
	case http.MethodPatch:
		router.Patch(path, handler)
	}
}
