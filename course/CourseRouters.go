package course

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe/middleware"
)

func NewCourseRouter(app *fiber.App) {
	app.Use(middleware.LoggingMiddleware)
	app.Use(middleware.AuthMiddleware)

	//app.Use("/course", middleware.Protected())
	app.Get("/course", func(c *fiber.Ctx) error {
		return MsvcCourse("", c)
	})
	app.Get("/course/:id", func(c *fiber.Ctx) error {
		return MsvcCourse("", c)
	})
	app.Post("/course", func(c *fiber.Ctx) error {
		return MsvcCourse("", c)
	})
	app.Put("/course/:id", func(c *fiber.Ctx) error {
		return MsvcCourse("", c)
	})
	app.Delete("/course/:id", func(c *fiber.Ctx) error {
		return MsvcCourse("", c)
	})
	//course for school
	app.Get("/course/school", func(c *fiber.Ctx) error {
		return MsvcCourse("school", c)
	})
	app.Post("/course/school", func(c *fiber.Ctx) error {
		return MsvcCourse("school", c)
	})
	app.Delete("/course/school/:id", func(c *fiber.Ctx) error {
		return MsvcCourse("school", c)
	})
}
