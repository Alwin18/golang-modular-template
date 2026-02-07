package config

import (
	"os"
)

type Config struct {
	AppName string
	AppPort string

	DBHost string
	DBUser string
	DBPass string
	DBName string
	DBPort string

	RedisHost string
	RedisPort string
}

func LoadConfig() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "go-modular"),
		AppPort: getEnv("APP_PORT", "3000"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASSWORD", "postgres"),
		DBName: getEnv("DB_NAME", "app_db"),
		DBPort: getEnv("DB_PORT", "5432"),

		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
