package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nayan9229/go-ad-services/shared/utilities"
	"github.com/nayan9229/go-ad-services/ssai/config"
	"github.com/nayan9229/go-ad-services/ssai/server"
	"github.com/rs/zerolog/log"
)

var appname = "ssai"
var release string

func main() {
	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load configuration")
	}
	cfg.Version = release
	cfg.Appname = appname

	utilities.LogSetup(appname, cfg.DevMode)

	// Initialize logger
	log.Info().Str("version", cfg.Version).Msg("Starting service")

	// Connect to database
	// db, err := repositories.NewPostgresDB(cfg.Database)
	// if err != nil {
	// 	log.Fatal("Failed to connect to database", "error", err)
	// }
	// defer db.Close()

	// Initialize repositories
	// userRepo := repositories.NewUserRepository(db)

	// // Initialize services
	// userService := services.NewUserService(userRepo)

	// Setup HTTP server
	server := server.NewServer(cfg)

	// Start server in goroutine
	go func() {
		log.Info().Str("address", cfg.Server.Address).Msg("Starting HTTP server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Failed to start server")
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exited properly")
}
