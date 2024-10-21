package course

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewLanguageRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	//app.Use("/language", middleware.Protected())
	app.Get("/language/", func(c *fiber.Ctx) error {
		return MsvcLanguage(c)
	})
	app.Get("/language/:id", func(c *fiber.Ctx) error {
		return MsvcLanguage(c)
	})
	app.Post("/language", func(c *fiber.Ctx) error {
		return MsvcLanguage(c)
	})
	app.Put("/language/:id", func(c *fiber.Ctx) error {
		return MsvcLanguage(c)
	})
	app.Delete("/language/:id", func(c *fiber.Ctx) error {
		return MsvcLanguage(c)
	})
}
