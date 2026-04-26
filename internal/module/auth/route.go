package auth

import (
	"github.com/Alwin18/golang-module-template/internal/http/middleware"
	"github.com/Alwin18/golang-module-template/internal/shared/cache"

	"github.com/gofiber/fiber/v2"
)

// RegisterRoutes registers all auth endpoints.
func RegisterRoutes(router fiber.Router, h *Handler, cacheClient *cache.Cache, authMW fiber.Handler) {
	auth := router.Group("/auth")

	auth.Post("/register", middleware.RegisterRateLimit(cacheClient), h.Register)
	auth.Post("/login", middleware.LoginRateLimit(cacheClient), h.Login)
	auth.Post("/refresh", h.Refresh)
	auth.Post("/logout", authMW, h.Logout)
	auth.Get("/me", authMW, h.Me)
}
