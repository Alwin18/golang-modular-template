package http

import (
	"github.com/Alwin18/golang-modular-template/internal/module/user"
	"github.com/gofiber/fiber/v2"
)

type Deps struct {
	UserService user.Service
}

func RegisterRoutes(app *fiber.App, d Deps) {
	api := app.Group("/api")

	user.RegisterRoutes(api, d.UserService)
}
