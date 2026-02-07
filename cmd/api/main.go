package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Alwin18/golang-modular-template/config"
	"github.com/Alwin18/golang-modular-template/internal/app"
)

func main() {
	cfg := config.LoadConfig()

	container, err := app.NewContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	fiberApp := app.NewApp(container)

	// Channel to listen for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := fiberApp.Listen(":" + cfg.AppPort); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	log.Printf("Server started on port %s. Press Ctrl+C to shutdown.", cfg.AppPort)

	// Wait for interrupt signal
	<-quit
	log.Println("Shutting down server...")

	// Create context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown Fiber server gracefully
	if err := fiberApp.ShutdownWithContext(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}

	// Cleanup resources (database, redis, etc.)
	container.Cleanup()

	log.Println("Server shutdown complete")
}
