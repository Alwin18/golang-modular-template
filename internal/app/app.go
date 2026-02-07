package app

import (
	httpDelivery "github.com/Alwin18/golang-modular-template/internal/http"
	"github.com/gofiber/fiber/v2"
)

func NewApp(c *Container) *fiber.App {
	app := fiber.New()

	httpDelivery.RegisterRoutes(app, httpDelivery.Deps{
		UserService: c.UserService,
	})

	return app
}
