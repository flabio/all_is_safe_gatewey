package states

import "github.com/gofiber/fiber/v2"

func NewStatesRouter(app *fiber.App) {
	app.Get("/states/", func(c *fiber.Ctx) error {
		return MsvcStates(c)
	})
	app.Get("/states/:id", func(c *fiber.Ctx) error {
		return MsvcStates(c)
	})
	app.Get("/states/city/:city_id", func(c *fiber.Ctx) error {
		return MsvcStates(c)
	})
	app.Post("/states", func(c *fiber.Ctx) error {
		return MsvcStates(c)
	})
	app.Put("/states/:id", func(c *fiber.Ctx) error {
		return MsvcStates(c)
	})
	app.Delete("/states/:id", func(c *fiber.Ctx) error {
		return MsvcStates(c)
	})
}
