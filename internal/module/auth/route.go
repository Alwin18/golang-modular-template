package auth

import (
	"github.com/Alwin18/golang-module-template/internal/http/middleware"
	"github.com/Alwin18/golang-module-template/internal/shared/cache"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all auth endpoints.
func RegisterRoutes(router fiber.Router, h *Handler, cacheClient *cache.Cache, authMW fiber.Handler) {
	auth := router.Group("/auth")

	auth.Post("/login", middleware.LoginRateLimit(cacheClient), h.Login)

}
