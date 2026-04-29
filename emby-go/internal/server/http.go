package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/emby/emby-go/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// HTTPServer wraps the HTTP server with chi router.
type HTTPServer struct {
	config *config.Config
	router *chi.Mux
	server *http.Server
	logger *zap.Logger
}

// NewHTTPServer creates a new HTTP server instance.
func NewHTTPServer(cfg *config.Config, logger *zap.Logger) *HTTPServer {
	router := chi.NewRouter()

	// Core middleware
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))
	router.Use(loggerMiddleware(logger))

	return &HTTPServer{
		config: cfg,
		router: router,
		logger: logger,
	}
}

// loggerMiddleware logs HTTP requests.
func loggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			logger.Info("request completed",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Duration("duration", time.Since(start)),
			)
		})
	}
}

// Start begins listening for HTTP requests.
func (s *HTTPServer) Start() error {
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port),
		Handler:      s.router,
		MaxHeaderBytes: s.config.Server.MaxHeaderBytes,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
	}

	s.logger.Info("starting HTTP server",
		zap.String("host", s.config.Server.Host),
		zap.Int("port", s.config.Server.Port),
	)

	return s.server.ListenAndServe()
}

// Shutdown gracefully shuts down the HTTP server.
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.logger.Info("shutting down HTTP server")
	return s.server.Shutdown(ctx)
}

// Router returns the chi router for route registration.
func (s *HTTPServer) Router() *chi.Mux {
	return s.router
}
