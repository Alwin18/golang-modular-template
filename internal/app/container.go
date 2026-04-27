package app

import (
	"github.com/Alwin18/golang-module-template/config"
	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	"github.com/Alwin18/golang-module-template/internal/shared/logger"
	validation "github.com/Alwin18/golang-module-template/internal/shared/validations"
	"github.com/go-playground/validator/v10"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Container holds all shared dependencies for the application.
type Container struct {
	Config      *config.Config
	DB          *gorm.DB
	Redis       *redis.Client
	Cache       *cache.Cache
	Logger      *zap.Logger
	AsynqClient *asynq.Client
	AsynqServer *asynq.Server
	Validator   *validator.Validate
}

// NewContainer creates and wires all dependencies.
func NewContainer(cfg *config.Config) (*Container, error) {
	// Init logger
	log, err := logger.New(cfg.IsDevelopment())
	if err != nil {
		return nil, err
	}

	// Init DB
	db, err := initDB(cfg, log)
	if err != nil {
		return nil, err
	}

	// Init Redis
	rdb, err := initRedis(cfg, log)
	if err != nil {
		return nil, err
	}

	// Init cache wrapper
	cacheClient := cache.New(rdb)

	// validator
	validator := validation.NewValidator()

	// Init Asynq
	redisOpt := asynq.RedisClientOpt{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
	}
	asynqClient := asynq.NewClient(redisOpt)
	asynqServer := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
	})

	return &Container{
		Config:      cfg,
		DB:          db,
		Redis:       rdb,
		Cache:       cacheClient,
		Logger:      log,
		AsynqClient: asynqClient,
		AsynqServer: asynqServer,
		Validator:   validator,
	}, nil
}

// Close gracefully shuts down all connections.
func (c *Container) Close() {
	if c.AsynqClient != nil {
		_ = c.AsynqClient.Close()
	}
	if c.Redis != nil {
		_ = c.Redis.Close()
	}
	if sqlDB, err := c.DB.DB(); err == nil {
		_ = sqlDB.Close()
	}
	_ = c.Logger.Sync()
}
