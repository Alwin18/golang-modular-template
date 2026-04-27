package health

import (
	"context"
	"time"

	"github.com/Alwin18/golang-module-template/internal/shared/response"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// Handler handles health check requests.
type Handler struct {
	db    *gorm.DB
	redis *redis.Client
}

// NewHandler creates a new health Handler.
func NewHandler(db *gorm.DB, redis *redis.Client) *Handler {
	return &Handler{db: db, redis: redis}
}

// Check performs a liveness check on DB and Redis.
func (h *Handler) Check(ctx *fiber.Ctx) error {
	// Use context from Fiber request
	dbCtx, cancel := context.WithTimeout(ctx.UserContext(), 3*time.Second)
	defer cancel()

	dbStatus := "connected"
	if sqlDB, err := h.db.DB(); err != nil || sqlDB.PingContext(dbCtx) != nil {
		dbStatus = "disconnected"
	}

	redisStatus := "connected"
	if err := h.redis.Ping(dbCtx).Err(); err != nil {
		redisStatus = "disconnected"
	}

	status := "ok"
	if dbStatus != "connected" || redisStatus != "connected" {
		status = "degraded"
	}

	return ctx.Status(fiber.StatusOK).JSON(response.NewResponse(fiber.Map{
		"status": status,
		"db":     dbStatus,
		"redis":  redisStatus,
	}, "success login", fiber.StatusOK))
}

// RegisterRoutes registers health routes.
func RegisterRoutes(router fiber.Router, h *Handler) {
	router.Get("/health", h.Check)
}
