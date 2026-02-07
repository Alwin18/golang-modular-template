package app

import (
	"github.com/Alwin18/golang-modular-template/config"
	"github.com/Alwin18/golang-modular-template/internal/module/user"
	"github.com/Alwin18/golang-modular-template/internal/shared/db"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"github.com/Alwin18/golang-modular-template/internal/shared/redis"
)

type Container struct {
	DB     *db.DB
	Redis  *redis.Client
	Logger logger.Logger

	UserService user.Service
}

func NewContainer(cfg *config.Config) (*Container, error) {
	log := logger.New()

	database, err := db.NewPostgres(cfg, log)
	if err != nil {
		return nil, err
	}

	redisClient := redis.New()

	userRepo := user.NewGormRepository(database.Gorm)
	userService := user.NewService(userRepo, log)

	return &Container{
		DB:          database,
		Redis:       redisClient,
		Logger:      log,
		UserService: userService,
	}, nil
}

// Cleanup closes all resources (database, redis, etc.)
func (c *Container) Cleanup() {
	c.Logger.Info("Starting cleanup...")

	// Close database connection
	if err := c.DB.Close(c.Logger); err != nil {
		c.Logger.Error("Error closing database:", err)
	}

	// Close redis connection
	if err := c.Redis.Close(); err != nil {
		c.Logger.Error("Error closing redis:", err)
	} else {
		c.Logger.Info("Redis connection closed successfully")
	}

	c.Logger.Info("Cleanup completed")
}
