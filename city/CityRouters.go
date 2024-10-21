package city

import "github.com/gofiber/fiber/v2"

func NewCityRouter(app *fiber.App) {
	app.Get("/cities/", func(c *fiber.Ctx) error {
		return MsvcCity(c)
	})
	app.Get("/cities/:id", func(c *fiber.Ctx) error {
		return MsvcCity(c)
	})
	app.Post("/cities", func(c *fiber.Ctx) error {
		return MsvcCity(c)
	})
	app.Put("/cities/:id", func(c *fiber.Ctx) error {
		return MsvcCity(c)
	})
	app.Delete("/cities/:id", func(c *fiber.Ctx) error {
		return MsvcCity(c)
	})
}
