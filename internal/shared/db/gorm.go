package db

import (
	"fmt"
	"os"
	"time"

	"github.com/Alwin18/golang-modular-template/config"
	"github.com/Alwin18/golang-modular-template/internal/shared/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Gorm *gorm.DB
}

func NewPostgres(cfg *config.Config, log logger.Logger) (*DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.New(
		// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
		// 	logger.Config{
		// 		SlowThreshold: time.Second, // Query di atas 1 detik dianggap lambat
		// 		LogLevel:      logger.Info, // Bisa diganti ke Warn atau Error jika terlalu verbose
		// 		Colorful:      true,
		// 	},
		// ),
	})
	if err != nil {
		log.Error("Failed to connect to database:", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Error("failed to connect database: %v", err)
	}

	if err = connection.Ping(); err != nil {
		log.Error("Failed to ping database:", err)
	}

	log.Info("Connected and pinged PostgreSQL database successfully")

	// Konfigurasi koneksi pool
	maxIdleConns := 30
	// if cfg.SetMaxIdleConns != "" {
	// 	maxIdleConns, _ = strconv.Atoi(cfg.SetMaxIdleConns)
	// }

	maxOpenConns := 100
	// if cfg.SetMaxOpenConns != "" {
	// 	maxOpenConns, _ = strconv.Atoi(cfg.SetMaxOpenConns)
	// }

	maxLifeTimeConnection := 300
	// if cfg.SetMaxLifeTime != "" {
	// 	maxLifeTimeConnection, _ = strconv.Atoi(cfg.SetMaxLifeTime)
	// }

	connection.SetMaxIdleConns(maxIdleConns)
	connection.SetMaxOpenConns(maxOpenConns)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	if os.Getenv("ENV") == "development" {
		db = db.Debug()
	}

	return &DB{Gorm: db}, nil
}
