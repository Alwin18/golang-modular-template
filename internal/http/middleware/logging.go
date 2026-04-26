package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// Logger returns a structured request-logging middleware.
func Logger(log *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		latency := time.Since(start)

		fields := []zap.Field{
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("latency", latency),
			zap.String("ip", c.IP()),
		}

		// Attach user ID if present from JWT middleware
		if userID, ok := c.Locals("user_id").(uint); ok && userID != 0 {
			fields = append(fields, zap.Uint("user_id", userID))
		}

		if err != nil {
			fields = append(fields, zap.Error(err))
			log.Error("request error", fields...)
		} else {
			log.Info("request", fields...)
		}

		return err
	}
}
