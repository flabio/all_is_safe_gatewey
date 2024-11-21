package course

import (
	"github.com/gofiber/fiber/v2"
)

func NewTypeCourseRouter(app *fiber.App) {
	// app.Use(middleware.LoggingMiddleware)
	// app.Use(middleware.AuthMiddleware)

	//app.Use("/course", middleware.Protected())
	app.Get("/type_course", func(c *fiber.Ctx) error {
		return MsvcTypeCourse(c)
	})
	app.Get("/type_course/:id", func(c *fiber.Ctx) error {
		return MsvcTypeCourse(c)
	})
	app.Post("/type_course", func(c *fiber.Ctx) error {
		return MsvcTypeCourse(c)
	})
	app.Put("/type_course/:id", func(c *fiber.Ctx) error {
		return MsvcTypeCourse(c)
	})
	app.Delete("/type_course/:id", func(c *fiber.Ctx) error {
		return MsvcTypeCourse(c)
	})

}
