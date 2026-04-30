package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/emby/emby-go/internal/api/middleware"
	"github.com/emby/emby-go/internal/config"
	cmid "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// HTTPServer represents the main HTTP server.
type HTTPServer struct {
	config *config.Config
	logger *zap.Logger
	router *chi.Mux
	server *http.Server
}

// NewHTTPServer creates a new HTTP server.
func NewHTTPServer(cfg *config.Config, logger *zap.Logger) *HTTPServer {
	return &HTTPServer{
		config: cfg,
		logger: logger,
	}
}

// Start starts the HTTP server.
func (s *HTTPServer) Start() error {
	// Create router
	s.router = chi.NewRouter()

	// Add chi middleware
	s.router.Use(cmid.RequestID)
	s.router.Use(cmid.RealIP)
	s.router.Use(cmid.Recoverer)
	s.router.Use(cmid.Timeout(60 * time.Second))
	s.router.Use(cmid.Logger)
	s.router.Use(cmid.AllowContentType("application/json"))

	// Add custom middleware
	s.router.Use(middleware.CORSMiddleware())
	s.router.Use(middleware.RequestLogger(s.logger))

	// Create HTTP server
	addr := fmt.Sprintf("%s:%d", s.config.Server.Host, s.config.Server.Port)
	s.server = &http.Server{
		Addr:         addr,
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s.logger.Info("HTTP server starting", zap.String("addr", addr))

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server failed: %w", err)
	}

	return nil
}

// Shutdown gracefully shuts down the HTTP server.
func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.logger.Info("HTTP server shutting down")
	return s.server.Shutdown(ctx)
}

// Router returns the Chi router.
func (s *HTTPServer) Router() *chi.Mux {
	return s.router
}

// GetConfig returns the server configuration.
func (s *HTTPServer) GetConfig() *config.Config {
	return s.config
}

// GetLogger returns the server logger.
func (s *HTTPServer) GetLogger() *zap.Logger {
	return s.logger
}
