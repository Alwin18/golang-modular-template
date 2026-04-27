package middleware

import (
	"fmt"
	"time"

	"github.com/Alwin18/golang-module-template/internal/shared/cache"
	"github.com/Alwin18/golang-module-template/internal/shared/response"
	"github.com/gofiber/fiber/v2"
)

// RateLimit returns a Redis-based rate limiter middleware.
// max: maximum requests allowed, window: time window.
func RateLimit(cacheClient *cache.Cache, max int64, window time.Duration) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		key := cache.RateLimitKey(fmt.Sprintf("%s:%s", ctx.Path(), ctx.IP()))

		count, err := cacheClient.Increment(ctx.Context(), key, window)
		if err != nil {
			// Fail open on Redis errors
			return ctx.Next()
		}

		if count > max {
			return ctx.Status(fiber.StatusTooManyRequests).JSON(response.NewErrorResponse(response.ResponseError{
				Message: "rate limit exceeded, please try again later",
				Code:    fiber.StatusTooManyRequests,
			}))
		}

		return ctx.Next()
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
