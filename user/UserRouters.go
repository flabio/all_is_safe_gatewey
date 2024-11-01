package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewUserRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	//app.Use("/user", middleware.Protected())

	app.Get("/user/", func(c *fiber.Ctx) error {

		return MsvcUser("", c)
	})
	app.Get("/user/students/", func(c *fiber.Ctx) error {

		return MsvcUser("students", c)
	})
	app.Get("/user/instructor/", func(c *fiber.Ctx) error {

		return MsvcUser("instructor", c)
	})
	app.Get("/user/:id", func(c *fiber.Ctx) error {
		return MsvcUser("", c)
	})
	app.Post("/user", func(c *fiber.Ctx) error {
		return MsvcUser("", c)
	})

	app.Put("/user/:id", func(c *fiber.Ctx) error {
		return MsvcUser("", c)
	})
	app.Put("/user/password/:id", func(c *fiber.Ctx) error {
		return MsvcUserUpdatePassword(c)
	})
	app.Put("/user/avatar/:id", func(c *fiber.Ctx) error {
		return MsvcUserAvatar(c)
	})
	app.Delete("/user/:id", func(c *fiber.Ctx) error {
		return MsvcUser("", c)
	})

}
