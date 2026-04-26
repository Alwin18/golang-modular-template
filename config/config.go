package config

// Config holds all application configuration.
type Config struct {
	AppPort string
	AppEnv  string

	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	RedisHost string
	RedisPort string

	JWTSecret          string
	JWTAccessTokenTTL  int // minutes
	JWTRefreshTokenTTL int // days
}

// IsDevelopment returns true if running in development mode.
func (c *Config) IsDevelopment() bool {
	return c.AppEnv == "local" || c.AppEnv == "development"
}

// IsProduction returns true if running in production mode.
func (c *Config) IsProduction() bool {
	return c.AppEnv == "production"
}
