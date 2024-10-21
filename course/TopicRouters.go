package course

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewTopicRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	//app.Use("/topic", middleware.Protected())
	app.Get("/topic/", func(c *fiber.Ctx) error {
		return MsvcTopic("",c)
	})
	app.Get("/topic/course/:course_id", func(c *fiber.Ctx) error {
		return MsvcTopic("course",c)
	})
	app.Get("/topic/:id", func(c *fiber.Ctx) error {
		return MsvcTopic("",c)
	})
	app.Post("/topic", func(c *fiber.Ctx) error {
		return MsvcTopic("",c)
	})
	app.Put("/topic/:id", func(c *fiber.Ctx) error {
		return MsvcTopic("",c)
	})
	app.Delete("/topic/:id", func(c *fiber.Ctx) error {
		return MsvcTopic("",c)
	})
}
