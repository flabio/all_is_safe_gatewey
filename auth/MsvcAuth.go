package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/dto"
	"github.com/safe/utils"
)

func MsvcAuth(c *fiber.Ctx) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	url := os.Getenv("MSVC_AUTH_URL")
	dataAuth := c.Body()
	var dataMapAuth map[string]interface{}
	json.Unmarshal(dataAuth, &dataMapAuth)

	req, err := http.NewRequest(c.Method(), url, bytes.NewBuffer(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FAILED_CREATE)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.SERVICE_NOT_AVAILALE)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Send([]byte(err.Error()))
	}
	var dataMap map[string]interface{}
	err = json.Unmarshal(respBody, &dataMap)
	if err != nil {
		return c.Send([]byte(respBody))
	}
	t, err := GenerateToken(dataMap)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Send([]byte(t))
}

func GenerateToken(dataMap map[string]interface{}) (string, error) {
	userData := dto.UserDTO{
		Id:        int(dataMap["Id"].(float64)),
		FirstName: dataMap["FirstName"].(string),
		LastName:  dataMap["FirstSurName"].(string),
		Email:     dataMap["Email"].(string),
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["full_name"] = userData.FirstName + " " + userData.LastName
	claims["email"] = userData.Email
	claims["id"] = userData.Id
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	t, err := token.SignedString([]byte("supersecretkey"))
	return t, err

}
