package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/safe/utils"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("supersecretkey"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

// AuthMiddleware checks if the request is authenticated
func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Get(utils.AUTHORIZATION)
	if len(token) > 7 && token[:7] == utils.BEARER {
		return c.Next()

	}
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		utils.STATUS:  fiber.StatusUnauthorized,
		utils.MESSAGE: utils.TOKEN_INVALID,
	})

}

// LoggingMiddleware logs the incoming requests
func LoggingMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	log.Printf("Started %s %s", c.Method(), c.Path())
	err := c.Next()
	log.Printf("Completed in %v", time.Since(start))
	return err
}
