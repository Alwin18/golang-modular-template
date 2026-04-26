package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Load reads environment variables and returns a Config.
func Load() (*Config, error) {
	// Load .env file if present (ignored in production)
	_ = godotenv.Load()

	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		AppEnv:  getEnv("APP_ENV", "local"),

		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", "secret"),
		DBName: getEnv("DB_NAME", "starter"),

		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),

		JWTSecret:          getEnv("JWT_SECRET", "supersecret"),
		JWTAccessTokenTTL:  getEnvInt("JWT_ACCESS_TOKEN_TTL", 15),
		JWTRefreshTokenTTL: getEnvInt("JWT_REFRESH_TOKEN_TTL", 7),
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	if c.JWTSecret == "" || c.JWTSecret == "supersecret" {
		if c.IsProduction() {
			return fmt.Errorf("JWT_SECRET must be set in production")
		}
	}
	return nil
}

func getEnv(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

func getEnvInt(key string, defaultVal int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return defaultVal
}
