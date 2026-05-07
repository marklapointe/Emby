package api

import (
	"net/http"
	"strings"

	"github.com/emby/emby-go/internal/version"
	"github.com/emby/emby-go/internal/api/handlers"
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/service/channel"
	"github.com/emby/emby-go/internal/service/device"
	"github.com/emby/emby-go/internal/service/dlna"
	"github.com/emby/emby-go/internal/service/image"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/media"
	"github.com/emby/emby-go/internal/service/notification"
	"github.com/emby/emby-go/internal/service/scheduled"
	"github.com/emby/emby-go/internal/service/session"
	"github.com/emby/emby-go/internal/service/sync"
	"github.com/emby/emby-go/internal/service/transcoding"
	"github.com/emby/emby-go/internal/service/user"
	"github.com/emby/emby-go/internal/server"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

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
	dlnaSvc         *dlna.Manager
	syncSvc         *sync.Manager
}

func NewRouter(cfg *config.Config, logger *zap.Logger, dbManager *database.Manager) *Router {
	itemRepo := repository.NewItemRepository(dbManager.DB())
	configRepo := repository.NewConfigRepository(dbManager.DB())
	userRepo := repository.NewUserRepository(dbManager.DB())

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

func (r *Router) RegisterRoutes(router *chi.Mux) {
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
	r.dlnaSvc = dlna.NewManager(r.logger)
	r.syncSvc = sync.NewManager(r.logger)

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
	r.registerDLNARoutes(embyRouter)
	r.registerSyncRoutes(embyRouter)
	r.registerPluginRoutes(embyRouter)
	r.registerCollectionRoutes(embyRouter)
	r.registerAuthRoutes(embyRouter)

	// Add missing routes
	// /emby redirects to /emby/web/
	embyRouter.Get("/emby", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/emby/web/", http.StatusFound)
	})
	embyRouter.Get("/emby/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/emby/web/", http.StatusFound)
	})
	// /emby/health returns server health info
	embyRouter.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"Status":"Healthy","Version":"4.8.1.0"}`))
	})
	// /emby/web redirects to /web/
	embyRouter.Get("/web", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web/", http.StatusFound)
	})
	embyRouter.Get("/web/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web/", http.StatusFound)
	})

	router.Mount("/emby", embyRouter)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/web/", http.StatusFound)
	})

	staticHandler := server.NewStaticHandler("web", version.Version, func() bool {
		cfg, err := r.configRepo.GetConfig()
		if err != nil {
			return false
		}
		return cfg.IsStartupWizardCompleted
	})
	router.Handle("/web/*", http.StripPrefix("/web/", staticHandler))
	router.Handle("/web", http.RedirectHandler("/web/", http.StatusFound))

	router.Get("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
}

func (r *Router) registerLibraryRoutes(router *chi.Mux) {
	libHandler := handlers.NewLibraryHandler(library.NewScanner(r.config, r.logger, r.itemRepo), r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/library/root", libHandler.GetLibraryRoot)
	r.registerRoute(router, http.MethodGet, "/library/items", libHandler.GetItems)
	r.registerRoute(router, http.MethodGet, "/library/mediafolders", libHandler.GetMediaFolders)
	r.registerRoute(router, http.MethodPost, "/library/mediafolders", libHandler.CreateMediaFolder)
	r.registerRoute(router, http.MethodGet, "/library/mediafolders/{id}", libHandler.GetMediaFolder)
	r.registerRoute(router, http.MethodDelete, "/library/mediafolders/{id}", libHandler.DeleteMediaFolder)
	r.registerRoute(router, http.MethodGet, "/library/mediafolders/{id}/items", libHandler.GetFolderItems)
	r.registerRoute(router, http.MethodPost, "/library/folders/fullscan", libHandler.ScanLibrary)
	r.registerRoute(router, http.MethodGet, "/library/virtualfolders", libHandler.GetVirtualFolders)
	router.Get("/Library/VirtualFolders", libHandler.GetVirtualFolders)
	r.registerRoute(router, http.MethodPost, "/library/virtualfolders", libHandler.AddVirtualFolder)
	router.Post("/Library/VirtualFolders", libHandler.AddVirtualFolder)
	r.registerRoute(router, http.MethodPost, "/library/virtualfolders/name", libHandler.RenameVirtualFolder)
	router.Post("/Library/VirtualFolders/Name", libHandler.RenameVirtualFolder)
	r.registerRoute(router, http.MethodPost, "/library/virtualfolders/paths", libHandler.AddMediaPath)
	router.Post("/Library/VirtualFolders/Paths", libHandler.AddMediaPath)
	r.registerRoute(router, http.MethodPost, "/library/virtualfolders/libraryoptions", libHandler.UpdateVirtualFolderOptions)
	router.Post("/Library/VirtualFolders/LibraryOptions", libHandler.UpdateVirtualFolderOptions)
	r.registerRoute(router, http.MethodGet, "/library/virtualfolders/{id}/items", libHandler.GetVirtualFolderItems)
	router.Get("/Libraries/AvailableOptions", libHandler.GetAvailableOptions)
}

func (r *Router) registerSessionRoutes(router *chi.Mux) {
	sessionHandler := handlers.NewSessionHandler(r.sessionSvc)

	r.registerRoute(router, http.MethodGet, "/sessions", sessionHandler.GetSessions)
	r.registerRoute(router, http.MethodGet, "/sessions/{id}", sessionHandler.GetSession)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/playing", sessionHandler.StartPlayback)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/playing/progress", sessionHandler.PlaybackProgress)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/playing/stopped", sessionHandler.StopPlayback)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/volume", sessionHandler.SetVolume)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/pause", sessionHandler.PausePlayback)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/unpause", sessionHandler.UnpausePlayback)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/togglefullscreen", sessionHandler.ToggleFullscreen)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/goto", sessionHandler.NavigateTo)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/sendkey", sessionHandler.SendKey)
	r.registerRoute(router, http.MethodPost, "/sessions/{id}/sendtext", sessionHandler.SendText)
	r.registerRoute(router, http.MethodDelete, "/sessions/{id}", sessionHandler.CloseSession)
}

func (r *Router) registerUserRoutes(router *chi.Mux) {
	userHandler := handlers.NewUserHandler(r.userSvc)

	r.registerRoute(router, http.MethodGet, "/users", userHandler.GetUsers)
	r.registerRoute(router, http.MethodGet, "/users/public", userHandler.GetPublicUsers)
	r.registerRoute(router, http.MethodPost, "/users/login", userHandler.Login)
	r.registerRoute(router, http.MethodPost, "/users/logout", userHandler.Logout)
	r.registerRoute(router, http.MethodGet, "/users/{id}", userHandler.GetUser)
	r.registerRoute(router, http.MethodPut, "/users/{id}", userHandler.UpdateUser)
	r.registerRoute(router, http.MethodDelete, "/users/{id}", userHandler.DeleteUser)
	r.registerRoute(router, http.MethodPost, "/users/{id}/password", userHandler.ChangePassword)
	r.registerRoute(router, http.MethodGet, "/users/{id}/images/{type}", userHandler.GetUserImage)
	r.registerRoute(router, http.MethodGet, "/users/{id}/configuration", userHandler.GetUserConfiguration)
	r.registerRoute(router, http.MethodPut, "/users/{id}/configuration", userHandler.UpdateUserConfiguration)
	r.registerRoute(router, http.MethodGet, "/users/{id}/policy", userHandler.GetUserPolicy)
	r.registerRoute(router, http.MethodPut, "/users/{id}/policy", userHandler.UpdateUserPolicy)
	r.registerRoute(router, http.MethodGet, "/users/device/{deviceId}", userHandler.GetUsersByDevice)
	r.registerRoute(router, http.MethodGet, "/users/libraryfolders/{folderId}", userHandler.GetUsersByLibraryFolder)
}

func (r *Router) registerDeviceRoutes(router *chi.Mux) {
	deviceHandler := handlers.NewDeviceHandler(r.deviceSvc)

	r.registerRoute(router, http.MethodGet, "/devices", deviceHandler.GetDevices)
	r.registerRoute(router, http.MethodGet, "/devices/{id}", deviceHandler.GetDevice)
	r.registerRoute(router, http.MethodPut, "/devices/{id}", deviceHandler.UpdateDevice)
	r.registerRoute(router, http.MethodDelete, "/devices/{id}", deviceHandler.DeleteDevice)
	r.registerRoute(router, http.MethodGet, "/devices/{id}/icon", deviceHandler.GetDeviceIcon)
	r.registerRoute(router, http.MethodGet, "/devices/{id}/profile", deviceHandler.GetDeviceProfile)
}

func (r *Router) registerImageRoutes(router *chi.Mux) {
	imageHandler := handlers.NewImageHandler(r.imageSvc)

	r.registerRoute(router, http.MethodGet, "/items/{id}/images", imageHandler.GetItemImageInfos)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}", imageHandler.GetItemImage)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}/blurhash", imageHandler.GetItemImageBlurHash)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}/{index}", imageHandler.GetItemImageByIndex)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}/tag/{tag}", imageHandler.GetItemImageTag)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}/crop", imageHandler.GetItemImageCrop)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}/resize", imageHandler.GetItemImageResize)
	r.registerRoute(router, http.MethodGet, "/items/{id}/images/{type}/rotate", imageHandler.GetItemImageRotation)
	r.registerRoute(router, http.MethodGet, "/items/{id}/remoteimages/providers", imageHandler.GetRemoteImageProviders)
}

func (r *Router) registerMediaRoutes(router *chi.Mux) {
	mediaHandler := handlers.NewMediaHandler(r.mediaSvc)

	r.registerRoute(router, http.MethodGet, "/items/{id}", mediaHandler.GetItem)
	r.registerRoute(router, http.MethodGet, "/items/{id}/stream", mediaHandler.GetStream)
	r.registerRoute(router, http.MethodGet, "/items/{id}/subtitles", mediaHandler.GetSubtitles)
	r.registerRoute(router, http.MethodGet, "/items/{id}/subtitles/{index}/stream", mediaHandler.GetSubtitleStream)
	r.registerRoute(router, http.MethodGet, "/items/{id}/audio", mediaHandler.GetAudioStream)
}

func (r *Router) registerNotificationRoutes(router *chi.Mux) {
	notificationHandler := handlers.NewNotificationHandler(r.notificationSvc)

	r.registerRoute(router, http.MethodGet, "/notifications", notificationHandler.GetNotifications)
	r.registerRoute(router, http.MethodGet, "/notifications/unread", notificationHandler.GetUnreadNotifications)
	r.registerRoute(router, http.MethodPost, "/notifications/{id}/markread", notificationHandler.MarkAsRead)
	r.registerRoute(router, http.MethodPost, "/notifications/markallread", notificationHandler.MarkAllAsRead)
	r.registerRoute(router, http.MethodDelete, "/notifications/{id}", notificationHandler.DeleteNotification)
	r.registerRoute(router, http.MethodGet, "/notifications/count", notificationHandler.GetNotificationCount)
	r.registerRoute(router, http.MethodGet, "/notifications/unreadcount", notificationHandler.GetUnreadNotificationCount)
}

func (r *Router) registerScheduledTaskRoutes(router *chi.Mux) {
	scheduledHandler := handlers.NewScheduledTaskHandler(r.scheduledSvc)

	r.registerRoute(router, http.MethodGet, "/scheduledtasks", scheduledHandler.GetAllTasks)
	r.registerRoute(router, http.MethodGet, "/scheduledtasks/running", scheduledHandler.GetRunningTasks)
	r.registerRoute(router, http.MethodGet, "/scheduledtasks/{id}", scheduledHandler.GetTask)
	r.registerRoute(router, http.MethodPost, "/scheduledtasks/{id}/execute", scheduledHandler.ExecuteTask)
	r.registerRoute(router, http.MethodPost, "/scheduledtasks/{id}/cancel", scheduledHandler.CancelTask)
	r.registerRoute(router, http.MethodGet, "/scheduledtasks/count", scheduledHandler.GetTaskCount)
	r.registerRoute(router, http.MethodGet, "/scheduledtasks/runningcount", scheduledHandler.GetRunningTaskCount)
}

func (r *Router) registerTranscodingRoutes(router *chi.Mux) {
	transcodingHandler := handlers.NewTranscodingHandler(r.transcodingSvc)

	r.registerRoute(router, http.MethodGet, "/transcodingprofiles", transcodingHandler.GetTranscodingProfiles)
	r.registerRoute(router, http.MethodGet, "/transcodingprofiles/{id}", transcodingHandler.GetTranscodingProfile)
	r.registerRoute(router, http.MethodGet, "/activetranscodes", transcodingHandler.GetActiveTranscodes)
	r.registerRoute(router, http.MethodGet, "/activetranscodes/{id}", transcodingHandler.GetActiveTranscode)
	r.registerRoute(router, http.MethodPost, "/activetranscodes/{id}/stop", transcodingHandler.StopTranscode)
}

func (r *Router) registerChannelRoutes(router *chi.Mux) {
	channelHandler := handlers.NewChannelHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/channels", channelHandler.GetChannels)
	r.registerRoute(router, http.MethodGet, "/channels/{id}", channelHandler.GetChannel)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/folders", channelHandler.GetChannelFolders)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/items", channelHandler.GetChannelItems)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/images/{type}", channelHandler.GetChannelImage)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/logoimage", channelHandler.GetChannelLogoImage)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/bannerimage", channelHandler.GetChannelBannerImage)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/backdropimage", channelHandler.GetChannelBackdropImage)
	r.registerRoute(router, http.MethodGet, "/channels/{id}/thumbimage", channelHandler.GetChannelThumbImage)
}

func (r *Router) registerLiveTVRoutes(router *chi.Mux) {
	liveTVHandler := handlers.NewLiveTVHandler(r.itemRepo, r.logger)

	r.registerRoute(router, http.MethodGet, "/livetv/channels", liveTVHandler.GetChannels)
	r.registerRoute(router, http.MethodGet, "/livetv/channels/{id}", liveTVHandler.GetChannels)
	r.registerRoute(router, http.MethodGet, "/livetv/programs", liveTVHandler.GetPrograms)
	r.registerRoute(router, http.MethodGet, "/livetv/programs/{id}", liveTVHandler.GetProgram)
	r.registerRoute(router, http.MethodGet, "/livetv/recordings", liveTVHandler.GetRecordings)
	r.registerRoute(router, http.MethodGet, "/livetv/recordings/{id}", liveTVHandler.GetRecording)
	r.registerRoute(router, http.MethodGet, "/livetv/timers", liveTVHandler.GetTimers)
	r.registerRoute(router, http.MethodGet, "/livetv/info", liveTVHandler.GetGuideInfo)
	r.registerRoute(router, http.MethodGet, "/livetv/recommendedprograms", liveTVHandler.GetRecommendedPrograms)
	r.registerRoute(router, http.MethodGet, "/livetv/seriestimers", liveTVHandler.GetSeriesTimers)
	r.registerRoute(router, http.MethodGet, "/livetv/timerproviders", liveTVHandler.GetTimerProviders)
	r.registerRoute(router, http.MethodGet, "/livetv/tunerhosts", liveTVHandler.GetTunerHosts)
	r.registerRoute(router, http.MethodGet, "/livetv/tunerhosts/{id}", liveTVHandler.GetTunerHost)
	r.registerRoute(router, http.MethodPost, "/livetv/tunerhosts", liveTVHandler.CreateTunerHost)
	r.registerRoute(router, http.MethodDelete, "/livetv/tunerhosts/{id}", liveTVHandler.DeleteTunerHost)
	r.registerRoute(router, http.MethodGet, "/livetv/tunerhosts/types", liveTVHandler.GetTunerHostTypes)
	r.registerRoute(router, http.MethodGet, "/livetv/listingproviders", liveTVHandler.GetListingProviders)
	r.registerRoute(router, http.MethodPost, "/livetv/listingproviders", liveTVHandler.CreateListingProvider)
	r.registerRoute(router, http.MethodGet, "/livetv/listingproviders/default", liveTVHandler.GetDefaultListingProvider)
	r.registerRoute(router, http.MethodGet, "/livetv/listingproviders/schedulesdirect/countries", liveTVHandler.GetSchedulesDirectCountries)
	r.registerRoute(router, http.MethodPost, "/livetv/channelmappings", liveTVHandler.CreateChannelMapping)
	r.registerRoute(router, http.MethodGet, "/livetv/channelmappingoptions", liveTVHandler.GetChannelMappingOptions)
}

func (r *Router) registerMoviesRoutes(router *chi.Mux) {
	moviesHandler := handlers.NewMoviesHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/movies", moviesHandler.GetMovies)
	r.registerRoute(router, http.MethodGet, "/movies/{id}", moviesHandler.GetMovie)
	r.registerRoute(router, http.MethodGet, "/movies/{id}/similar", moviesHandler.GetSimilar)
	r.registerRoute(router, http.MethodGet, "/movies/recommendations", moviesHandler.GetRecommendations)
}

func (r *Router) registerTVShowRoutes(router *chi.Mux) {
	tvShowsHandler := handlers.NewTVShowsHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/tvshows", tvShowsHandler.GetTVShows)
	r.registerRoute(router, http.MethodGet, "/tvshows/{id}", tvShowsHandler.GetTVShow)
	r.registerRoute(router, http.MethodGet, "/tvshows/{id}/seasons", tvShowsHandler.GetSeasons)
	r.registerRoute(router, http.MethodGet, "/tvshows/{id}/episodes", tvShowsHandler.GetEpisodes)
	r.registerRoute(router, http.MethodGet, "/tvshows/{id}/similar", tvShowsHandler.GetSimilar)
}

func (r *Router) registerSystemRoutes(router *chi.Mux) {
	systemHandler := handlers.NewSystemHandler(r.config, r.configRepo, r.logger)

	r.registerRoute(router, http.MethodGet, "/system/info", systemHandler.GetInfo)
	r.registerRoute(router, http.MethodGet, "/system/info/public", systemHandler.GetPublicSystemInfo)
	r.registerRoute(router, http.MethodGet, "/system/logs", systemHandler.GetLogs)
	r.registerRoute(router, http.MethodGet, "/system/logs/log", systemHandler.GetLog)
	r.registerRoute(router, http.MethodGet, "/system/configuration", systemHandler.GetConfiguration)
	r.registerRoute(router, http.MethodGet, "/system/ping", systemHandler.Ping)
	r.registerRoute(router, http.MethodPost, "/system/shutdown", systemHandler.Shutdown)
}

func (r *Router) registerPlaylistRoutes(router *chi.Mux) {
	playlistHandler := handlers.NewPlaylistHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/playlists", playlistHandler.GetPlaylists)
	r.registerRoute(router, http.MethodGet, "/playlists/{id}", playlistHandler.GetPlaylist)
	r.registerRoute(router, http.MethodGet, "/playlists/{id}/items", playlistHandler.GetPlaylistItems)
	r.registerRoute(router, http.MethodPost, "/playlists", playlistHandler.CreatePlaylist)
	r.registerRoute(router, http.MethodPut, "/playlists/{id}", playlistHandler.UpdatePlaylist)
	r.registerRoute(router, http.MethodDelete, "/playlists/{id}", playlistHandler.DeletePlaylist)
}

func (r *Router) registerActivityRoutes(router *chi.Mux) {
	activityHandler := handlers.NewActivityHandler()

	r.registerRoute(router, http.MethodGet, "/activities", activityHandler.GetActivities)
	r.registerRoute(router, http.MethodGet, "/activities/{id}", activityHandler.GetActivity)
}

func (r *Router) registerBrandingRoutes(router *chi.Mux) {
	brandingHandler := handlers.NewBrandingHandler(r.config)

	r.registerRoute(router, http.MethodGet, "/branding/css", brandingHandler.GetCss)
	r.registerRoute(router, http.MethodGet, "/branding/css.css", brandingHandler.GetCss)
	r.registerRoute(router, http.MethodGet, "/branding/json", brandingHandler.GetJson)
	r.registerRoute(router, http.MethodGet, "/branding/images/{name}", brandingHandler.GetImage)
	r.registerRoute(router, http.MethodGet, "/branding/configuration", brandingHandler.GetConfiguration)
	r.registerRoute(router, http.MethodGet, "/branding/options", brandingHandler.GetBrandingOptions)
}

func (r *Router) registerConfigRoutes(router *chi.Mux) {
	configHandler := handlers.NewConfigHandler(r.config, r.logger)

	r.registerRoute(router, http.MethodGet, "/configuration", configHandler.GetConfiguration)
	r.registerRoute(router, http.MethodPut, "/configuration", configHandler.UpdateConfiguration)
	r.registerRoute(router, http.MethodGet, "/configuration/{name}", configHandler.GetNamedConfiguration)
}

func (r *Router) registerDisplayPrefsRoutes(router *chi.Mux) {
	displayPrefsHandler := handlers.NewDisplayPrefsHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/displaypreferences/{id}", displayPrefsHandler.GetDisplayPreferences)
	r.registerRoute(router, http.MethodPost, "/displaypreferences/{id}", displayPrefsHandler.UpdateDisplayPreferences)
}

func (r *Router) registerEnvironmentRoutes(router *chi.Mux) {
	environmentHandler := handlers.NewEnvironmentHandler()

	r.registerRoute(router, http.MethodGet, "/environment/drives", environmentHandler.GetDrives)
	r.registerRoute(router, http.MethodGet, "/environment/networkshares", environmentHandler.GetNetworkShares)
	r.registerRoute(router, http.MethodGet, "/environment/parentpath", environmentHandler.GetParentPath)
	r.registerRoute(router, http.MethodGet, "/environment/directorycontents", environmentHandler.GetDirectoryContents)
	r.registerRoute(router, http.MethodPost, "/environment/validatepath", environmentHandler.ValidatePath)
	router.Get("/Environment/DefaultDirectoryBrowser", environmentHandler.GetDefaultDirectoryBrowser)
	router.Get("/Environment/ParentPath", environmentHandler.GetParentPath)
	router.Get("/Environment/DirectoryContents", environmentHandler.GetDirectoryContents)
	router.Post("/Environment/ValidatePath", environmentHandler.ValidatePath)
}

func (r *Router) registerFilterRoutes(router *chi.Mux) {
	filterHandler := handlers.NewFilterHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/genres", filterHandler.GetGenres)
	r.registerRoute(router, http.MethodGet, "/studios", filterHandler.GetStudios)
	r.registerRoute(router, http.MethodGet, "/years", filterHandler.GetYears)
	r.registerRoute(router, http.MethodGet, "/cultures", filterHandler.GetCultures)
	r.registerRoute(router, http.MethodGet, "/countries", filterHandler.GetCountries)
	r.registerRoute(router, http.MethodGet, "/musicgenres", filterHandler.GetMusicGenres)
	r.registerRoute(router, http.MethodGet, "/artists", filterHandler.GetArtists)
	r.registerRoute(router, http.MethodGet, "/albumartists", filterHandler.GetAlbumArtists)
}

func (r *Router) registerGamesRoutes(router *chi.Mux) {
	gamesHandler := handlers.NewGamesHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/games", gamesHandler.GetGames)
	r.registerRoute(router, http.MethodGet, "/games/{id}", gamesHandler.GetGame)
}

func (r *Router) registerLocalizationRoutes(router *chi.Mux) {
	localizationHandler := handlers.NewLocalizationHandler(r.logger)

	r.registerRoute(router, http.MethodGet, "/localization/cultures", localizationHandler.GetCultures)
	r.registerRoute(router, http.MethodGet, "/localization/countries", localizationHandler.GetCountries)
	r.registerRoute(router, http.MethodGet, "/localization/parentalratings", localizationHandler.GetParentalRatings)
	r.registerRoute(router, http.MethodGet, "/localization/options", localizationHandler.GetOptions)
	r.registerRoute(router, http.MethodGet, "/localization/{culture}", localizationHandler.GetLocalization)
}

func (r *Router) registerPackageRoutes(router *chi.Mux) {
	packageHandler := handlers.NewPackageHandler()

	r.registerRoute(router, http.MethodGet, "/packages", packageHandler.GetPackages)
	r.registerRoute(router, http.MethodGet, "/packages/{name}", packageHandler.GetPackage)
	r.registerRoute(router, http.MethodPost, "/packages/install", packageHandler.Install)
	r.registerRoute(router, http.MethodPost, "/packages/{name}/uninstall", packageHandler.Uninstall)
}

func (r *Router) registerPlaybackRoutes(router *chi.Mux) {
	playbackHandler := handlers.NewPlaybackHandler(r.sessionSvc)

	r.registerRoute(router, http.MethodPost, "/playback/{type}/selected", playbackHandler.SelectPlayback)
	r.registerRoute(router, http.MethodGet, "/playback/{type}/formats", playbackHandler.GetFormats)
}

func (r *Router) registerSearchRoutes(router *chi.Mux) {
	searchHandler := handlers.NewSearchHandler(r.itemRepo)

	r.registerRoute(router, http.MethodGet, "/search/hints", searchHandler.GetHints)
	r.registerRoute(router, http.MethodGet, "/search/items", searchHandler.SearchItemsByTerm)
}

func (r *Router) registerStartupRoutes(router *chi.Mux) {
	startupHandler := handlers.NewStartupHandler(r.configRepo, r.userRepo, r.logger)

	r.registerRoute(router, http.MethodGet, "/startup/first", startupHandler.IsFirstRun)
	r.registerRoute(router, http.MethodGet, "/startup/options", startupHandler.GetOptions)
	r.registerRoute(router, http.MethodPost, "/startup/complete", startupHandler.Complete)
	r.registerRoute(router, http.MethodGet, "/startup/configuration", startupHandler.GetStartupConfig)
	r.registerRoute(router, http.MethodPost, "/startup/configuration", startupHandler.PostStartupConfig)
	r.registerRoute(router, http.MethodGet, "/startup/user", startupHandler.GetStartupUser)
	r.registerRoute(router, http.MethodPost, "/startup/user", startupHandler.PostUser)
	r.registerRoute(router, http.MethodGet, "/startup/remoteaccess", startupHandler.GetStartupRemoteAccess)
	r.registerRoute(router, http.MethodPost, "/startup/remoteaccess", startupHandler.PostStartupRemoteAccess)
	r.registerRoute(router, http.MethodGet, "/startup/dashboard", startupHandler.GetStartupDashboardInfo)
	r.registerRoute(router, http.MethodGet, "/localization/options", startupHandler.GetStartupLanguage)
	r.registerRoute(router, http.MethodGet, "/wizardsettings", startupHandler.GetWizardSettings)
}

func (r *Router) registerDLNARoutes(router *chi.Mux) {
	dlnaHandler := handlers.NewDLNAHandler(r.dlnaSvc)

	r.registerRoute(router, http.MethodGet, "/dlna/profiles", dlnaHandler.GetProfiles)
	r.registerRoute(router, http.MethodPost, "/dlna/profiles", dlnaHandler.GetProfiles)
	r.registerRoute(router, http.MethodGet, "/dlna/profiles/{id}", dlnaHandler.GetProfile)
	r.registerRoute(router, http.MethodPost, "/dlna/profiles/{id}", dlnaHandler.GetProfile)
	r.registerRoute(router, http.MethodDelete, "/dlna/profiles/{id}", dlnaHandler.GetProfile)
	r.registerRoute(router, http.MethodGet, "/dlna/profiles/default", dlnaHandler.GetDefaultProfile)
	r.registerRoute(router, http.MethodGet, "/dlna/profileinfos", dlnaHandler.GetProfileInfos)
}

func (r *Router) registerSyncRoutes(router *chi.Mux) {
	syncHandler := handlers.NewSyncHandler(r.syncSvc)

	r.registerRoute(router, http.MethodGet, "/sync/jobs", syncHandler.GetJobs)
	r.registerRoute(router, http.MethodPost, "/sync/jobs", syncHandler.CreateJob)
	r.registerRoute(router, http.MethodGet, "/sync/jobs/{id}", syncHandler.GetJob)
	r.registerRoute(router, http.MethodDelete, "/sync/jobs/{id}", syncHandler.DeleteJob)
	r.registerRoute(router, http.MethodPost, "/sync/jobs/{id}/items/{itemId}", syncHandler.AddItemToJob)
}

func (r *Router) registerPluginRoutes(router *chi.Mux) {
	pluginHandler := handlers.NewPluginHandler()

	r.registerRoute(router, http.MethodGet, "/plugins", pluginHandler.GetPlugins)
	r.registerRoute(router, http.MethodDelete, "/plugins/{id}", pluginHandler.DeletePlugin)
	r.registerRoute(router, http.MethodGet, "/plugins/{id}/configuration", pluginHandler.GetPluginConfiguration)
	r.registerRoute(router, http.MethodGet, "/plugins/securityinfo", pluginHandler.GetSecurityInfo)
	r.registerRoute(router, http.MethodGet, "/plugins/released", pluginHandler.GetReleased)
}

func (r *Router) registerCollectionRoutes(router *chi.Mux) {
	collectionHandler := handlers.NewCollectionHandler(r.itemRepo)

	r.registerRoute(router, http.MethodPost, "/collections", collectionHandler.CreateCollection)
	r.registerRoute(router, http.MethodPost, "/collections/{id}/items", collectionHandler.AddToCollection)
}

func (r *Router) registerAuthRoutes(router *chi.Mux) {
	authHandler := handlers.NewAuthHandler()

	r.registerRoute(router, http.MethodGet, "/auth/providers", authHandler.GetAuthProviders)
}

func toTitleCase(s string) string {
	specialCases := map[string]string{
		"livetv":                      "LiveTv",
		"dlna":                        "Dlna",
		"scheduledtasks":              "ScheduledTasks",
		"recommendedprograms":         "RecommendedPrograms",
		"seriestimers":               "SeriesTimers",
		"timerproviders":              "TimerProviders",
		"tunerhosts":                  "TunerHosts",
		"tunerhosts/types":            "TunerHosts/Types",
		"listingproviders":           "ListingProviders",
		"listingproviders/default":   "ListingProviders/Default",
		"schedulesdirect":            "SchedulesDirect",
		"schedulesdirect/countries":  "SchedulesDirect/Countries",
		"channelmappings":             "ChannelMappings",
		"channelmappingoptions":      "ChannelMappingOptions",
		"profileinfos":                "ProfileInfos",
		"usersettings":               "UserSettings",
		"displaypreferences":         "DisplayPreferences",
		"musicgenres":                 "MusicGenres",
		"localization":                "Localization",
		"wizardsettings":             "WizardSettings",
		"mediafolders":               "MediaFolders",
		"virtualfolders":             "VirtualFolders",
		"availableoptions":            "AvailableOptions",
		"defaultdirectorybrowser":     "DefaultDirectoryBrowser",
		"directorycontents":          "DirectoryContents",
		"albumartists":               "AlbumArtists",
		"artists":                    "Artists",
	}

	parts := strings.Split(s, "/")
	for i, part := range parts {
		if len(part) > 0 {
			if mapped, ok := specialCases[part]; ok {
				parts[i] = mapped
			} else if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
				continue
			} else if strings.Contains(part, "{") {
				parts[i] = toTitleCasePreserveVariables(part)
			} else {
				parts[i] = strings.ToUpper(part[:1]) + strings.ToLower(part[1:])
			}
		}
	}
	return strings.Join(parts, "/")
}

func toTitleCasePreserveVariables(s string) string {
	var result strings.Builder
	var varStart int

	for i := 0; i < len(s); i++ {
		if s[i] == '{' {
			if varStart < i {
				result.WriteString(strings.ToUpper(s[varStart:i][:1]) + strings.ToLower(s[varStart:i][1:]))
			}
			varStart = i
		} else if s[i] == '}' {
			result.WriteString(s[varStart : i+1])
			varStart = i + 1
		}
	}

	if varStart < len(s) {
		result.WriteString(strings.ToUpper(s[varStart:varStart+1]) + strings.ToLower(s[varStart+1:]))
	}

	return result.String()
}

func (r *Router) registerRoute(router *chi.Mux, method, path string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		router.Get(path, handler)
		router.Get(toTitleCase(path), handler)
	case http.MethodPost:
		router.Post(path, handler)
		router.Post(toTitleCase(path), handler)
	case http.MethodPut:
		router.Put(path, handler)
		router.Put(toTitleCase(path), handler)
	case http.MethodDelete:
		router.Delete(path, handler)
		router.Delete(toTitleCase(path), handler)
	case http.MethodPatch:
		router.Patch(path, handler)
		router.Patch(toTitleCase(path), handler)
	}
}