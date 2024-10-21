package school

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewSchoolRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	//app.Use("/user", middleware.Protected())
	app.Get("/school/", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Get("/school/:id", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Post("/school", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Put("/school/:id", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
	app.Delete("/school/:id", func(c *fiber.Ctx) error {
		return MsvcSchool(c)
	})
}
