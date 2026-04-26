package user

import "github.com/gofiber/fiber/v2"

// RegisterRoutes registers all user endpoints.
func RegisterRoutes(router fiber.Router, h *Handler, authMW fiber.Handler) {
	users := router.Group("/users", authMW)

	users.Get("/", h.List)
	users.Get("/:id", h.GetByID)
	users.Put("/:id", h.Update)
	users.Delete("/:id", h.Delete)
}
