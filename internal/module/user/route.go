package user

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(r fiber.Router, s Service) {
	h := NewHandler(s)

	r.Get("/users", h.GetUsers)
}
