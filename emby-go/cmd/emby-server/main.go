package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emby/emby-go/internal/api"
	"github.com/emby/emby-go/internal/api/middleware"
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/logging"
	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/server"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/scheduled"
	cmid "github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	logger, err := logging.NewLogger(cfg.Logging.Level, cfg.Logging.Format)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	logger.Info("Emby Server Go starting",
		zap.String("version", "0.1.0"),
		zap.String("host", cfg.Server.Host),
		zap.Int("port", cfg.Server.Port),
	)

	// Initialize database
	dbManager, err := database.NewManager(&cfg.Database)
	if err != nil {
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer dbManager.Close()

	logger.Info("database connected", zap.String("path", cfg.Database.Path))

	// Initialize repositories
	itemRepo := repository.NewItemRepository(dbManager.DB())
	if err := itemRepo.CreateSchema(); err != nil {
		logger.Fatal("Failed to create schema", zap.Error(err))
	}

	// Initialize HTTP server first
	httpServer := server.NewHTTPServer(cfg, logger)

	// Add middleware to the HTTP server's chi router
	httpServer.Router().Use(cmid.RequestID)
	httpServer.Router().Use(cmid.RealIP)
	httpServer.Router().Use(cmid.Recoverer)
	httpServer.Router().Use(cmid.Timeout(60 * time.Second))
	httpServer.Router().Use(cmid.Logger)
	httpServer.Router().Use(cmid.AllowContentType("application/json"))
	httpServer.Router().Use(middleware.CORSMiddleware())
	httpServer.Router().Use(middleware.RequestLogger(logger))

	// Initialize API router and register routes on the HTTP server's chi router
	apiRouter := api.NewRouter(cfg, logger, dbManager)
	apiRouter.RegisterRoutes(httpServer.Router())

	// Initialize library scanner with scheduled tasks
	scheduledSvc := scheduled.NewManager(logger)
	scanner := library.NewScanner(cfg, logger, itemRepo)
	scanner.StartScheduledScans(scheduledSvc)

	// Setup graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		logger.Info("received shutdown signal")
		cancel()
	}()

	// Start HTTP server in a goroutine
	go func() {
		if err := httpServer.Start(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("HTTP server error", zap.Error(err))
		}
	}()

	logger.Info("Emby Server Go started",
		zap.String("host", cfg.Server.Host),
		zap.Int("port", cfg.Server.Port),
	)

	// Wait for shutdown signal
	<-ctx.Done()
	logger.Info("shutting down...")

	// Gracefully shutdown HTTP server
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Error("HTTP server shutdown error", zap.Error(err))
	}

	logger.Info("Emby Server Go stopped")
}
