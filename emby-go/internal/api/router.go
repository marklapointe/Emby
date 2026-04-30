package api

import (
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/api/handlers"
	"github.com/emby/emby-go/internal/service/device"
	"github.com/emby/emby-go/internal/service/image"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/media"
	"github.com/emby/emby-go/internal/service/notification"
	"github.com/emby/emby-go/internal/service/scheduled"
	"github.com/emby/emby-go/internal/service/session"
	"github.com/emby/emby-go/internal/service/transcoding"
	"github.com/emby/emby-go/internal/service/user"
	"go.uber.org/zap"
)

// Router manages API routes and handlers.
type Router struct {
	config   *config.Config
	logger   *zap.Logger
	dbManager *database.Manager
	userSvc  *user.Manager
	librarySvc *library.Manager
	mediaSvc *media.Manager
	sessionSvc *session.Manager
	deviceSvc *device.Manager
	imageSvc *image.Manager
	notificationSvc *notification.Manager
	scheduledSvc *scheduled.Manager
	transcodingSvc *transcoding.Manager
}

// NewRouter creates a new API router.
func NewRouter(cfg *config.Config, logger *zap.Logger, dbManager *database.Manager) *Router {
	return &Router{
		config:    cfg,
		logger:    logger,
		dbManager: dbManager,
	}
}

// RegisterAll registers all API routes.
func (r *Router) RegisterAll() {
	// Initialize services
	r.userSvc = user.NewManager(r.dbManager, r.logger)
	r.librarySvc = library.NewManager(r.config, r.logger, r.dbManager)
	r.mediaSvc = media.NewManager(r.config, r.logger)
	r.sessionSvc = session.NewManager(r.config, r.logger)
	r.deviceSvc = device.NewManager(r.logger)
	r.imageSvc = image.NewManager(r.logger)
	r.notificationSvc = notification.NewManager(r.logger)
	r.scheduledSvc = scheduled.NewManager(r.logger)
	r.transcodingSvc = transcoding.NewManager(r.config, r.logger)

	// Register routes
	r.registerLibraryRoutes()
	r.registerSessionRoutes()
	r.registerUserRoutes()
	r.registerDeviceRoutes()
	r.registerImageRoutes()
	r.registerMediaRoutes()
	r.registerNotificationRoutes()
	r.registerScheduledTaskRoutes()
	r.registerTranscodingRoutes()
}

// registerLibraryRoutes registers library-related routes.
func (r *Router) registerLibraryRoutes() {
	libHandler := handlers.NewLibraryHandler(library.NewScanner(r.config, r.logger, nil))

	r.registerRoute(http.MethodGet, "/Library/Root", libHandler.GetLibraryRoot)
	r.registerRoute(http.MethodGet, "/Library/Items", libHandler.GetItems)
	r.registerRoute(http.MethodGet, "/Library/MediaFolders", libHandler.GetMediaFolders)
	r.registerRoute(http.MethodPost, "/Library/MediaFolders", libHandler.CreateMediaFolder)
	r.registerRoute(http.MethodGet, "/Library/MediaFolders/{id}", libHandler.GetMediaFolder)
	r.registerRoute(http.MethodDelete, "/Library/MediaFolders/{id}", libHandler.DeleteMediaFolder)
	r.registerRoute(http.MethodGet, "/Library/MediaFolders/{id}/Items", libHandler.GetFolderItems)
	r.registerRoute(http.MethodPost, "/Library/Folders/FullScan", libHandler.ScanLibrary)
	r.registerRoute(http.MethodGet, "/Library/VirtualFolders", libHandler.GetVirtualFolders)
	r.registerRoute(http.MethodGet, "/Library/VirtualFolders/{id}/Items", libHandler.GetVirtualFolderItems)
}

// registerSessionRoutes registers session-related routes.
func (r *Router) registerSessionRoutes() {
	sessionHandler := handlers.NewSessionHandler(r.sessionSvc)

	r.registerRoute(http.MethodGet, "/Sessions", sessionHandler.GetSessions)
	r.registerRoute(http.MethodGet, "/Sessions/{id}", sessionHandler.GetSession)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/Playing", sessionHandler.StartPlayback)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/Playing/Progress", sessionHandler.PlaybackProgress)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/Playing/Stopped", sessionHandler.StopPlayback)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/Volume", sessionHandler.SetVolume)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/Pause", sessionHandler.PausePlayback)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/Unpause", sessionHandler.UnpausePlayback)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/ToggleFullscreen", sessionHandler.ToggleFullscreen)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/GoTo", sessionHandler.NavigateTo)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/SendKey", sessionHandler.SendKey)
	r.registerRoute(http.MethodPost, "/Sessions/{id}/SendText", sessionHandler.SendText)
	r.registerRoute(http.MethodDelete, "/Sessions/{id}", sessionHandler.CloseSession)
}

// registerUserRoutes registers user-related routes.
func (r *Router) registerUserRoutes() {
	userHandler := handlers.NewUserHandler(r.userSvc)

	r.registerRoute(http.MethodGet, "/Users", userHandler.GetUsers)
	r.registerRoute(http.MethodGet, "/Users/Public", userHandler.GetPublicUsers)
	r.registerRoute(http.MethodPost, "/Users/Login", userHandler.Login)
	r.registerRoute(http.MethodPost, "/Users/Logout", userHandler.Logout)
	r.registerRoute(http.MethodGet, "/Users/{id}", userHandler.GetUser)
	r.registerRoute(http.MethodPut, "/Users/{id}", userHandler.UpdateUser)
	r.registerRoute(http.MethodDelete, "/Users/{id}", userHandler.DeleteUser)
	r.registerRoute(http.MethodPost, "/Users/{id}/Password", userHandler.ChangePassword)
	r.registerRoute(http.MethodGet, "/Users/{id}/Images/{type}", userHandler.GetUserImage)
	r.registerRoute(http.MethodGet, "/Users/{id}/Configuration", userHandler.GetUserConfiguration)
	r.registerRoute(http.MethodPut, "/Users/{id}/Configuration", userHandler.UpdateUserConfiguration)
	r.registerRoute(http.MethodGet, "/Users/{id}/Policy", userHandler.GetUserPolicy)
	r.registerRoute(http.MethodPut, "/Users/{id}/Policy", userHandler.UpdateUserPolicy)
	r.registerRoute(http.MethodGet, "/Users/Device/{deviceId}", userHandler.GetUsersByDevice)
	r.registerRoute(http.MethodGet, "/Users/LibraryFolders/{folderId}", userHandler.GetUsersByLibraryFolder)
}

// registerDeviceRoutes registers device-related routes.
func (r *Router) registerDeviceRoutes() {
	deviceHandler := handlers.NewDeviceHandler(r.deviceSvc)

	r.registerRoute(http.MethodGet, "/Devices", deviceHandler.GetDevices)
	r.registerRoute(http.MethodGet, "/Devices/{id}", deviceHandler.GetDevice)
	r.registerRoute(http.MethodPut, "/Devices/{id}", deviceHandler.UpdateDevice)
	r.registerRoute(http.MethodDelete, "/Devices/{id}", deviceHandler.DeleteDevice)
	r.registerRoute(http.MethodGet, "/Devices/{id}/Icon", deviceHandler.GetDeviceIcon)
	r.registerRoute(http.MethodGet, "/Devices/{id}/Profile", deviceHandler.GetDeviceProfile)
}

// registerImageRoutes registers image-related routes.
func (r *Router) registerImageRoutes() {
	imageHandler := handlers.NewImageHandler(r.imageSvc)

	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}", imageHandler.GetItemImage)
	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}/BlurHash", imageHandler.GetItemImageBlurHash)
	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}/{index}", imageHandler.GetItemImageByIndex)
	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}/Tag/{tag}", imageHandler.GetItemImageTag)
	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}/Crop", imageHandler.GetItemImageCrop)
	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}/Resize", imageHandler.GetItemImageResize)
	r.registerRoute(http.MethodGet, "/Items/{id}/Images/{type}/Rotate", imageHandler.GetItemImageRotation)
}

// registerMediaRoutes registers media-related routes.
func (r *Router) registerMediaRoutes() {
	mediaHandler := handlers.NewMediaHandler(r.mediaSvc)

	r.registerRoute(http.MethodGet, "/Items/{id}", mediaHandler.GetItem)
	r.registerRoute(http.MethodGet, "/Items/{id}/Stream", mediaHandler.GetStream)
	r.registerRoute(http.MethodGet, "/Items/{id}/Subtitles", mediaHandler.GetSubtitles)
	r.registerRoute(http.MethodGet, "/Items/{id}/Subtitles/{index}/Stream", mediaHandler.GetSubtitleStream)
	r.registerRoute(http.MethodGet, "/Items/{id}/Audio", mediaHandler.GetAudioStream)
}

// registerNotificationRoutes registers notification-related routes.
func (r *Router) registerNotificationRoutes() {
	notificationHandler := handlers.NewNotificationHandler(r.notificationSvc)

	r.registerRoute(http.MethodGet, "/Notifications", notificationHandler.GetNotifications)
	r.registerRoute(http.MethodGet, "/Notifications/Unread", notificationHandler.GetUnreadNotifications)
	r.registerRoute(http.MethodPost, "/Notifications/{id}/MarkRead", notificationHandler.MarkAsRead)
	r.registerRoute(http.MethodPost, "/Notifications/MarkAllRead", notificationHandler.MarkAllAsRead)
	r.registerRoute(http.MethodDelete, "/Notifications/{id}", notificationHandler.DeleteNotification)
	r.registerRoute(http.MethodGet, "/Notifications/Count", notificationHandler.GetNotificationCount)
	r.registerRoute(http.MethodGet, "/Notifications/UnreadCount", notificationHandler.GetUnreadNotificationCount)
}

// registerScheduledTaskRoutes registers scheduled task-related routes.
func (r *Router) registerScheduledTaskRoutes() {
	scheduledHandler := handlers.NewScheduledTaskHandler(r.scheduledSvc)

	r.registerRoute(http.MethodGet, "/ScheduledTasks", scheduledHandler.GetAllTasks)
	r.registerRoute(http.MethodGet, "/ScheduledTasks/Running", scheduledHandler.GetRunningTasks)
	r.registerRoute(http.MethodGet, "/ScheduledTasks/{id}", scheduledHandler.GetTask)
	r.registerRoute(http.MethodPost, "/ScheduledTasks/{id}/Execute", scheduledHandler.ExecuteTask)
	r.registerRoute(http.MethodPost, "/ScheduledTasks/{id}/Cancel", scheduledHandler.CancelTask)
	r.registerRoute(http.MethodGet, "/ScheduledTasks/Count", scheduledHandler.GetTaskCount)
	r.registerRoute(http.MethodGet, "/ScheduledTasks/RunningCount", scheduledHandler.GetRunningTaskCount)
}

// registerTranscodingRoutes registers transcoding-related routes.
func (r *Router) registerTranscodingRoutes() {
	transcodingHandler := handlers.NewTranscodingHandler(r.transcodingSvc)

	r.registerRoute(http.MethodGet, "/TranscodingProfiles", transcodingHandler.GetTranscodingProfiles)
	r.registerRoute(http.MethodGet, "/TranscodingProfiles/{id}", transcodingHandler.GetTranscodingProfile)
	r.registerRoute(http.MethodGet, "/ActiveTranscodes", transcodingHandler.GetActiveTranscodes)
	r.registerRoute(http.MethodGet, "/ActiveTranscodes/{id}", transcodingHandler.GetActiveTranscode)
	r.registerRoute(http.MethodPost, "/ActiveTranscodes/{id}/Stop", transcodingHandler.StopTranscode)
}

// registerRoute registers a single route.
func (r *Router) registerRoute(method, path string, handler http.HandlerFunc) {
	routeMap := map[string]map[string]http.HandlerFunc{
		http.MethodGet:    {},
		http.MethodPost:   {},
		http.MethodPut:    {},
		http.MethodDelete: {},
	}

	routeMap[method][path] = handler
}
