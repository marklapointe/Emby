package api

import (
	"net/http"

	"github.com/emby/emby-go/internal/api/handlers"
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/service/channel"
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
	itemRepo    *repository.ItemRepository
	configRepo  *repository.ConfigRepository
	userRepo    *repository.UserRepository
	userSvc     *user.Manager
	librarySvc  *library.Manager
	mediaSvc    *media.Manager
	sessionSvc  *session.Manager
	deviceSvc   *device.Manager
	imageSvc    *image.Manager
	notificationSvc *notification.Manager
	scheduledSvc    *scheduled.Manager
	transcodingSvc   *transcoding.Manager
	channelSvc      *channel.Manager
}

// NewRouter creates a new API router.
func NewRouter(cfg *config.Config, logger *zap.Logger, dbManager *database.Manager) *Router {
	itemRepo := repository.NewItemRepository(dbManager.DB())
	configRepo := repository.NewConfigRepository(dbManager.DB())
	userRepo := repository.NewUserRepository(dbManager.DB())

	// Initialize database schema
	configRepo.CreateConfigTable()
	userRepo.CreateUserTable()
	itemRepo.CreateSchema()

	return &Router{
		config:    cfg,
		logger:    logger,
		dbManager: dbManager,
		itemRepo:  itemRepo,
		configRepo: configRepo,
		userRepo:  userRepo,
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
	r.channelSvc = channel.NewManager(r.logger)

	// Create subrouter for /emby API routes
	embyRouter := chi.NewRouter()
	r.registerLibraryRoutes(embyRouter)
	r.registerSessionRoutes(embyRouter)
	r.registerUserRoutes(embyRouter)
	r.registerDeviceRoutes(embyRouter)
	r.registerImageRoutes(embyRouter)
	r.registerMediaRoutes(embyRouter)
	r.registerNotificationRoutes(embyRouter)
	r.registerScheduledTaskRoutes(embyRouter)
	r.registerTranscodingRoutes(embyRouter)
	r.registerChannelRoutes(embyRouter)
	r.registerLiveTVRoutes(embyRouter)
	r.registerMoviesRoutes(embyRouter)
	r.registerTVShowRoutes(embyRouter)
	r.registerSystemRoutes(embyRouter)
	r.registerPlaylistRoutes(embyRouter)
	r.registerActivityRoutes(embyRouter)
	r.registerBrandingRoutes(embyRouter)
	r.registerConfigRoutes(embyRouter)
	r.registerDisplayPrefsRoutes(embyRouter)
	r.registerEnvironmentRoutes(embyRouter)
	r.registerFilterRoutes(embyRouter)
	r.registerGamesRoutes(embyRouter)
	r.registerLocalizationRoutes(embyRouter)
	r.registerPackageRoutes(embyRouter)
	r.registerPlaybackRoutes(embyRouter)
	r.registerSearchRoutes(embyRouter)
	r.registerStartupRoutes(embyRouter)

	// Mount /emby API routes
	router.Mount("/emby", embyRouter)

	// Register static file serving for web dashboard
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web/", http.StatusFound)
	})

	// Serve static files from /web/ path
	webFS := http.Dir("web")
	router.Handle("/web/*", http.StripPrefix("/web/", http.FileServer(webFS)))
	router.Handle("/web", http.RedirectHandler("/web/", http.StatusFound))

	// Register health endpoint
	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
}

// registerLibraryRoutes registers library-related routes.
func (r *Router) registerLibraryRoutes(router *chi.Mux) {
	libHandler := handlers.NewLibraryHandler(library.NewScanner(r.config, r.logger, r.itemRepo), r.itemRepo)

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

// registerChannelRoutes registers channel-related routes.
func (r *Router) registerChannelRoutes(router *chi.Mux) {
	channelHandler := handlers.NewChannelHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/Channels", channelHandler.GetChannels)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}", channelHandler.GetChannel)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/Folders", channelHandler.GetChannelFolders)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/Items", channelHandler.GetChannelItems)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/Images/{type}", channelHandler.GetChannelImage)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/LogoImage", channelHandler.GetChannelLogoImage)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/BannerImage", channelHandler.GetChannelBannerImage)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/BackdropImage", channelHandler.GetChannelBackdropImage)
	r.registerRoute(router, http.MethodGet, "/Channels/{id}/ThumbImage", channelHandler.GetChannelThumbImage)
}

// registerLiveTVRoutes registers LiveTV-related routes.
func (r *Router) registerLiveTVRoutes(router *chi.Mux) {
	liveTVHandler := handlers.NewLiveTVHandler(r.itemRepo, r.logger)

	r.registerRoute(router, http.MethodGet, "/LiveTv/Channels", liveTVHandler.GetChannels)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Channels/{id}", liveTVHandler.GetChannels)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Programs", liveTVHandler.GetPrograms)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Programs/{id}", liveTVHandler.GetProgram)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Recordings", liveTVHandler.GetRecordings)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Recordings/{id}", liveTVHandler.GetRecording)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Timers", liveTVHandler.GetTimers)
	r.registerRoute(router, http.MethodGet, "/LiveTv/Info", liveTVHandler.GetGuideInfo)
}

// registerMoviesRoutes registers movies-related routes.
func (r *Router) registerMoviesRoutes(router *chi.Mux) {
	moviesHandler := handlers.NewMoviesHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/Movies", moviesHandler.GetMovies)
	r.registerRoute(router, http.MethodGet, "/Movies/{id}", moviesHandler.GetMovie)
	r.registerRoute(router, http.MethodGet, "/Movies/{id}/Similar", moviesHandler.GetSimilar)
	r.registerRoute(router, http.MethodGet, "/Movies/Recommendations", moviesHandler.GetRecommendations)
}

// registerTVShowRoutes registers TV show-related routes.
func (r *Router) registerTVShowRoutes(router *chi.Mux) {
	tvShowsHandler := handlers.NewTVShowsHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/TvShows", tvShowsHandler.GetTVShows)
	r.registerRoute(router, http.MethodGet, "/TvShows/{id}", tvShowsHandler.GetTVShow)
	r.registerRoute(router, http.MethodGet, "/TvShows/{id}/Seasons", tvShowsHandler.GetSeasons)
	r.registerRoute(router, http.MethodGet, "/TvShows/{id}/Episodes", tvShowsHandler.GetEpisodes)
	r.registerRoute(router, http.MethodGet, "/TvShows/{id}/Similar", tvShowsHandler.GetSimilar)
}

// registerSystemRoutes registers system-related routes.
func (r *Router) registerSystemRoutes(router *chi.Mux) {
	systemHandler := handlers.NewSystemHandler(r.config, r.logger)

	r.registerRoute(router, http.MethodGet, "/System/Info", systemHandler.GetInfo)
	r.registerRoute(router, http.MethodGet, "/System/Info/Public", systemHandler.GetPublicSystemInfo)
	r.registerRoute(router, http.MethodGet, "/System/Logs", systemHandler.GetLogs)
	r.registerRoute(router, http.MethodGet, "/System/Logs/Log", systemHandler.GetLog)
	r.registerRoute(router, http.MethodGet, "/System/Configuration", systemHandler.GetConfiguration)
	r.registerRoute(router, http.MethodGet, "/System/Ping", systemHandler.Ping)
	r.registerRoute(router, http.MethodPost, "/System/Shutdown", systemHandler.Shutdown)
}

// registerPlaylistRoutes registers playlist-related routes.
func (r *Router) registerPlaylistRoutes(router *chi.Mux) {
	playlistHandler := handlers.NewPlaylistHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/Playlists", playlistHandler.GetPlaylists)
	r.registerRoute(router, http.MethodGet, "/Playlists/{id}", playlistHandler.GetPlaylist)
	r.registerRoute(router, http.MethodGet, "/Playlists/{id}/Items", playlistHandler.GetPlaylistItems)
	r.registerRoute(router, http.MethodPost, "/Playlists", playlistHandler.CreatePlaylist)
	r.registerRoute(router, http.MethodPut, "/Playlists/{id}", playlistHandler.UpdatePlaylist)
	r.registerRoute(router, http.MethodDelete, "/Playlists/{id}", playlistHandler.DeletePlaylist)
}

// registerActivityRoutes registers activity-related routes.
func (r *Router) registerActivityRoutes(router *chi.Mux) {
	activityHandler := handlers.NewActivityHandler()

	r.registerRoute(router, http.MethodGet, "/Activities", activityHandler.GetActivities)
	r.registerRoute(router, http.MethodGet, "/Activities/{id}", activityHandler.GetActivity)
}

// registerBrandingRoutes registers branding-related routes.
func (r *Router) registerBrandingRoutes(router *chi.Mux) {
	brandingHandler := handlers.NewBrandingHandler(r.config)

	r.registerRoute(router, http.MethodGet, "/Branding/Css", brandingHandler.GetCss)
	r.registerRoute(router, http.MethodGet, "/Branding/Json", brandingHandler.GetJson)
	r.registerRoute(router, http.MethodGet, "/Branding/Images/{name}", brandingHandler.GetImage)
}

// registerConfigRoutes registers configuration-related routes.
func (r *Router) registerConfigRoutes(router *chi.Mux) {
	configHandler := handlers.NewConfigHandler(r.config, r.logger)

	r.registerRoute(router, http.MethodGet, "/Configuration", configHandler.GetConfiguration)
	r.registerRoute(router, http.MethodPut, "/Configuration", configHandler.UpdateConfiguration)
	r.registerRoute(router, http.MethodGet, "/Configuration/{name}", configHandler.GetNamedConfiguration)
}

// registerDisplayPrefsRoutes registers display preferences routes.
func (r *Router) registerDisplayPrefsRoutes(router *chi.Mux) {
	displayPrefsHandler := handlers.NewDisplayPrefsHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/DisplayPreferences/{id}", displayPrefsHandler.GetDisplayPreferences)
	r.registerRoute(router, http.MethodPost, "/DisplayPreferences/{id}", displayPrefsHandler.UpdateDisplayPreferences)
}

// registerEnvironmentRoutes registers environment-related routes.
func (r *Router) registerEnvironmentRoutes(router *chi.Mux) {
	environmentHandler := handlers.NewEnvironmentHandler()

	r.registerRoute(router, http.MethodGet, "/Environment/Drives", environmentHandler.GetDrives)
	r.registerRoute(router, http.MethodGet, "/Environment/NetworkShares", environmentHandler.GetNetworkShares)
	r.registerRoute(router, http.MethodGet, "/Environment/ParentPath", environmentHandler.GetParentPath)
}

// registerFilterRoutes registers filter-related routes.
func (r *Router) registerFilterRoutes(router *chi.Mux) {
	filterHandler := handlers.NewFilterHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/Genres", filterHandler.GetGenres)
	r.registerRoute(router, http.MethodGet, "/Studios", filterHandler.GetStudios)
	r.registerRoute(router, http.MethodGet, "/Years", filterHandler.GetYears)
}

// registerGamesRoutes registers games-related routes.
func (r *Router) registerGamesRoutes(router *chi.Mux) {
	gamesHandler := handlers.NewGamesHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/Games", gamesHandler.GetGames)
	r.registerRoute(router, http.MethodGet, "/Games/{id}", gamesHandler.GetGame)
}

// registerLocalizationRoutes registers localization-related routes.
func (r *Router) registerLocalizationRoutes(router *chi.Mux) {
	localizationHandler := handlers.NewLocalizationHandler(r.logger)

	r.registerRoute(router, http.MethodGet, "/Localization/{culture}", localizationHandler.GetLocalization)
	r.registerRoute(router, http.MethodGet, "/Localization/Options", localizationHandler.GetOptions)
}

// registerPackageRoutes registers package-related routes.
func (r *Router) registerPackageRoutes(router *chi.Mux) {
	packageHandler := handlers.NewPackageHandler()

	r.registerRoute(router, http.MethodGet, "/Packages", packageHandler.GetPackages)
	r.registerRoute(router, http.MethodGet, "/Packages/{name}", packageHandler.GetPackage)
	r.registerRoute(router, http.MethodPost, "/Packages/Install", packageHandler.Install)
	r.registerRoute(router, http.MethodPost, "/Packages/{name}/Uninstall", packageHandler.Uninstall)
}

// registerPlaybackRoutes registers playback-related routes.
func (r *Router) registerPlaybackRoutes(router *chi.Mux) {
	playbackHandler := handlers.NewPlaybackHandler(r.sessionSvc)

	r.registerRoute(router, http.MethodPost, "/Playback/{type}/Selected", playbackHandler.SelectPlayback)
	r.registerRoute(router, http.MethodGet, "/Playback/{type}/Formats", playbackHandler.GetFormats)
}

// registerSearchRoutes registers search-related routes.
func (r *Router) registerSearchRoutes(router *chi.Mux) {
	searchHandler := handlers.NewSearchHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/Search/Hints", searchHandler.GetHints)
	r.registerRoute(router, http.MethodGet, "/Search/Items", searchHandler.SearchItemsByTerm)
}

// registerStartupRoutes registers startup-related routes.
func (r *Router) registerStartupRoutes(router *chi.Mux) {
	startupHandler := handlers.NewStartupHandler(r.configRepo, r.userRepo, r.logger)

	// Startup wizard endpoints (must match C# exactly for external app compatibility)
	r.registerRoute(router, http.MethodGet, "/Startup/First", startupHandler.IsFirstRun)
	r.registerRoute(router, http.MethodGet, "/Startup/Options", startupHandler.GetOptions)
	r.registerRoute(router, http.MethodPost, "/Startup/Complete", startupHandler.Complete)
	r.registerRoute(router, http.MethodGet, "/Startup/Configuration", startupHandler.GetStartupConfig)
	r.registerRoute(router, http.MethodPost, "/Startup/Configuration", startupHandler.PostStartupConfig)
	r.registerRoute(router, http.MethodGet, "/Startup/User", startupHandler.GetStartupUser)
	r.registerRoute(router, http.MethodPost, "/Startup/User", startupHandler.PostUser)
	r.registerRoute(router, http.MethodGet, "/Startup/RemoteAccess", startupHandler.GetStartupRemoteAccess)
	r.registerRoute(router, http.MethodPost, "/Startup/RemoteAccess", startupHandler.PostStartupRemoteAccess)
	r.registerRoute(router, http.MethodGet, "/Startup/Dashboard", startupHandler.GetStartupDashboardInfo)

	// Localization endpoints (used by wizard)
	r.registerRoute(router, http.MethodGet, "/Localization/Options", startupHandler.GetStartupLanguage)
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
