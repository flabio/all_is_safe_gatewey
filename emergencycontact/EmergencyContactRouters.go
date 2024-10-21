package emergencycontact

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewEmergencyContactRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)
	app.Use("/emergency", middleware.Protected())
	app.Get("/emergency/", func(c *fiber.Ctx) error {
		return EmergencyContact("", c)
	})
	app.Get("/emergency/:id", func(c *fiber.Ctx) error {
		return EmergencyContact("", c)
	})
	app.Get("/emergency/user/:id", func(c *fiber.Ctx) error {
		return EmergencyContact("user", c)
	})
	app.Post("/emergency", func(c *fiber.Ctx) error {
		return EmergencyContact("", c)
	})
	app.Put("/emergency/:id", func(c *fiber.Ctx) error {
		return EmergencyContact("", c)
	})
	app.Delete("/emergency/:id", func(c *fiber.Ctx) error {
		return EmergencyContact("", c)
	})
}
