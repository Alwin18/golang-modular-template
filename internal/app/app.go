package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Alwin18/golang-module-template/config"
	apphttp "github.com/Alwin18/golang-module-template/internal/http"
	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	"github.com/Alwin18/golang-module-template/internal/shared/db"

	rdb "github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// App is the root application struct.
type App struct {
	container *Container
}

// New initializes the application and all its dependencies.
func New(cfg *config.Config) (*App, error) {
	container, err := NewContainer(cfg)
	if err != nil {
		return nil, fmt.Errorf("container init: %w", err)
	}

	return &App{container: container}, nil
}

// Run starts the HTTP server and blocks until shutdown signal.
func (a *App) Run() error {
	// Pass a flat Deps struct — avoids importing internal/app inside internal/http
	router := apphttp.NewRouter(&apphttp.Deps{
		Config:    a.container.Config,
		DB:        a.container.DB,
		Redis:     a.container.Redis,
		Cache:     a.container.Cache,
		Logger:    a.container.Logger,
		Validator: a.container.Validator,
	})
	router.RegisterRoutes()

	fiberApp := router.App()

	// Start workers
	go a.container.startWorkers()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		addr := ":" + a.container.Config.AppPort
		a.container.Logger.Info("server starting", zap.String("addr", addr))
		if err := fiberApp.Listen(addr); err != nil {
			a.container.Logger.Error("server error", zap.Error(err))
		}
	}()

	<-quit
	a.container.Logger.Info("shutting down server...")

	if err := fiberApp.Shutdown(); err != nil {
		a.container.Logger.Error("server shutdown error", zap.Error(err))
	}

	a.container.AsynqServer.Shutdown()
	a.container.Close()

	a.container.Logger.Info("server stopped")
	return nil
}

// initDB initializes the postgres connection.
func initDB(cfg *config.Config, log *zap.Logger) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)
	database, err := db.NewPostgres(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres init: %w", err)
	}
	log.Info("postgres connected")
	return database, nil
}

// initRedis initializes the redis connection.
func initRedis(cfg *config.Config, log *zap.Logger) (*rdb.Client, error) {
	client, err := cache.NewRedis(cfg.RedisHost + ":" + cfg.RedisPort)
	if err != nil {
		return nil, fmt.Errorf("redis init: %w", err)
	}
	log.Info("redis connected")
	return client, nil
}
