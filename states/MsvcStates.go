package states

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
func MsvcStates(c *fiber.Ctx) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	url := os.Getenv("MSVC_STATES_URL")
	MSVC_STATES_BY_CITY_URL := os.Getenv("MSVC_STATES_BY_CITY_URL")

	id := c.Params(utils.ID)
	cityId := c.Params(utils.CITY_ID)

	if len(id) != 0 && url != "" {
		url = url + "/" + id
	}
	if len(cityId) != 0 && url != "" {
		url = MSVC_STATES_BY_CITY_URL + "/" + cityId
	}

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
	return c.Send(respBody)
}
