package middleware

import (
	"fmt"
	"time"

	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	response "github.com/Alwin18/golang-module-template/internal/shared/responses"
	"github.com/gofiber/fiber/v2"
)

// RateLimit returns a Redis-based rate limiter middleware.
// max: maximum requests allowed, window: time window.
func RateLimit(cacheClient *cache.Cache, max int64, window time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		key := cache.RateLimitKey(fmt.Sprintf("%s:%s", c.Path(), c.IP()))

		count, err := cacheClient.Increment(c.Context(), key, window)
		if err != nil {
			// Fail open on Redis errors
			return c.Next()
		}

		if count > max {
			return response.Error(c, fiber.StatusTooManyRequests, "rate limit exceeded, please try again later", nil)
		}

		return c.Next()
	}
}

// LoginRateLimit limits login attempts to 5/min per IP.
func LoginRateLimit(cacheClient *cache.Cache) fiber.Handler {
	return RateLimit(cacheClient, 5, time.Minute)
}

// RegisterRateLimit limits register attempts to 3/min per IP.
func RegisterRateLimit(cacheClient *cache.Cache) fiber.Handler {
	return RateLimit(cacheClient, 3, time.Minute)
}
