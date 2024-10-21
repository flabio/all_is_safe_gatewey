package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewAuthRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Post("/auth", func(c *fiber.Ctx) error {
		return MsvcAuth(c)
	})

}
