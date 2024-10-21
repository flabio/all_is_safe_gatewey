package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/safe/dto"
)

// Secret key for JWT signing
var jwtSecret = []byte("supersecretkey")

func Login(c *fiber.Ctx) error {
	var input dto.LoginDTO

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
	}
	log.Println(input)
	// user, err = h.userService.GetUserByUsername(input.Username)
	// if err != nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "errors": err.Error()})
	// }
	// ud = UserData{
	// 	ID:       user.ID,
	// 	Username: user.Username,
	// 	Email:    user.Email,
	// 	Password: user.Password,
	// }
	// if email == nil && user == nil {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "errors": err.Error()})
	// }

	// if !utils.CheckPasswordHash(pass, ud.Password) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	// }

	// token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = ud.Username
	// claims["user_id"] = ud.ID
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// t, err := token.SignedString([]byte(configs.EnvConfigs.JwtSecret))
	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": "ok"})
}
