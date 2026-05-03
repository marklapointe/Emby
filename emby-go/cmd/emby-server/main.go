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
	"github.com/emby/emby-go/internal/api/handlers"
	"github.com/emby/emby-go/internal/config"
	"github.com/emby/emby-go/internal/database"
	"github.com/emby/emby-go/internal/logging"
	"github.com/emby/emby-go/internal/repository"
	"github.com/emby/emby-go/internal/server"
	"github.com/emby/emby-go/internal/service/device"
	"github.com/emby/emby-go/internal/service/image"
	"github.com/emby/emby-go/internal/service/library"
	"github.com/emby/emby-go/internal/service/media"
	"github.com/emby/emby-go/internal/service/metadata"
	"github.com/emby/emby-go/internal/service/notification"
	"github.com/emby/emby-go/internal/service/scheduled"
	"github.com/emby/emby-go/internal/service/session"
	"github.com/emby/emby-go/internal/service/transcoding"
	"github.com/emby/emby-go/internal/service/user"
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

	// Initialize services
	_ = user.NewManager(dbManager, logger)
	_ = library.NewManager(cfg, logger, dbManager)
	_ = media.NewManager(cfg, logger)
	_ = session.NewManager(cfg, logger)
	_ = device.NewManager(logger)
	_ = image.NewManager(logger)
	_ = metadata.NewManager(logger)
	_ = notification.NewManager()
	scheduledSvc := scheduled.NewManager(logger)
	_ = transcoding.NewManager(cfg, logger)
	_ = media.NewStreamManager(cfg.Stream.MaxConcurrentStreams, logger)

	// Initialize library scanner
	scanner := library.NewScanner(cfg, logger, itemRepo)
	scanner.StartScheduledScans(scheduledSvc)

	// Initialize HTTP server
	httpServer := server.NewHTTPServer(cfg, logger)

	// Initialize API router
	apiRouter := api.NewRouter(cfg, logger, dbManager)
	apiRouter.RegisterAll()

	// Register handlers
	libHandler := handlers.NewLibraryHandler(library.NewScanner(cfg, logger, itemRepo))
	mediaHandler := handlers.NewMediaHandler(media.NewManager(cfg, logger))
	sessionHandler := handlers.NewSessionHandler(session.NewManager(cfg, logger))
	userHandler := handlers.NewUserHandler(user.NewManager(dbManager, logger))

	// Register routes
	r := httpServer.Router()
	r.Get("/Library/Root", libHandler.GetLibraryRoot)
	r.Get("/Library/Items", libHandler.GetItems)
	r.Post("/Library/Root/Scan", libHandler.ScanLibrary)
	r.Get("/Items/{id}", mediaHandler.GetItem)
	r.Get("/Items/{id}/Subtitles", mediaHandler.GetSubtitles)
	r.Get("/Sessions", sessionHandler.GetSessions)
	r.Get("/Sessions/{id}", sessionHandler.GetSession)
	r.Post("/Sessions/{id}/SendKey", sessionHandler.SendKey)
	r.Get("/Users", userHandler.GetUsers)
	r.Get("/Users/{id}", userHandler.GetUser)
	r.Post("/Users/Login", userHandler.Login)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

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
