package api

import (
	"net/http"

	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/server"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/media"
	"github.com/emby/emby-go/internal/service/session"
	"github.com/emby/emby-go/internal/service/user"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Router handles API route registration.
type Router struct {
	server     *server.HTTPServer
	config     *config.Config
	logger     *zap.Logger
	dbManager  *database.Manager
	userSvc    *user.Manager
	librarySvc *library.Manager
	mediaSvc   *media.Manager
	sessionSvc *session.Manager
}

// NewRouter creates a new API router.
func NewRouter(
	srv *server.HTTPServer,
	cfg *config.Config,
	logger *zap.Logger,
	dbManager *database.Manager,
) *Router {
	return &Router{
		server:    srv,
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

	// Register routes
	r.registerLibraryRoutes()
	r.registerUserRoutes()
	r.registerSessionRoutes()
	r.registerMediaRoutes()
	r.registerDeviceRoutes()
	r.registerPublicRoutes()
}

// registerLibraryRoutes registers library-related API endpoints.
func (r *Router) registerLibraryRoutes() {
	r.server.Router().Mount("/Library", chi.NewRouter())
}

// registerUserRoutes registers user-related API endpoints.
func (r *Router) registerUserRoutes() {
	r.server.Router().Mount("/Users", chi.NewRouter())
}

// registerSessionRoutes registers session-related API endpoints.
func (r *Router) registerSessionRoutes() {
	r.server.Router().Mount("/Sessions", chi.NewRouter())
}

// registerMediaRoutes registers media-related API endpoints.
func (r *Router) registerMediaRoutes() {
	r.server.Router().Mount("/Items", chi.NewRouter())
}

// registerDeviceRoutes registers device-related API endpoints.
func (r *Router) registerDeviceRoutes() {
	r.server.Router().Mount("/Devices", chi.NewRouter())
}

// registerPublicRoutes registers public (no auth required) API endpoints.
func (r *Router) registerPublicRoutes() {
	r.server.Router().Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})
}
