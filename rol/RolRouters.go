package rol

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewRolRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	// app.Use("/rol", middleware.Protected())
	app.Get("/rol/", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Get("/rol/:id", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Post("/rol", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Put("/rol/:id", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
	app.Delete("/rol/:id", func(c *fiber.Ctx) error {
		return MsvcRol(c)
	})
}
