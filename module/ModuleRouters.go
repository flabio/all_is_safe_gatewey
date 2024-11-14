package module

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewUserRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)
	//app.Use("/user", middleware.Protected())
	app.Get("/module/", func(c *fiber.Ctx) error {
		return MsModule(c)
	})
	app.Post("/module", func(c *fiber.Ctx) error {
		return MsModule(c)
	})
	app.Put("/module/:id", func(c *fiber.Ctx) error {
		return MsModule(c)
	})
	app.Delete("/module/:id", func(c *fiber.Ctx) error {
		return MsModule(c)
	})
	app.Post("/module/role", func(c *fiber.Ctx) error {
		return MsModuleRole(c)
	})
	app.Delete("/module/role/:id", func(c *fiber.Ctx) error {
		return MsModuleRole(c)
	})
}
