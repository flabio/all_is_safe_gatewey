package emergencycontact

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/utils"
)

// Handler for the  service
func EmergencyContact(path string, c *fiber.Ctx) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	MSVC_EMERGENCY_CONTACT_URL := os.Getenv("MSVC_EMERGENCY_CONTACT_URL")
	id := c.Params(utils.ID)
	url := MSVC_EMERGENCY_CONTACT_URL

	if len(id) != 0 && url != "" {
		url = MSVC_EMERGENCY_CONTACT_URL + "/" + id
	}
	if path != "" {
		url = MSVC_EMERGENCY_CONTACT_URL + "/" + path + "/" + id
	}
	req, err := http.NewRequest(c.Method(), url, bytes.NewBuffer(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FAILED_CREATE)
	}
	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.SERVICE_NOT_AVAILALE)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	return c.Send(respBody)
}
