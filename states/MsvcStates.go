package states

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/safe/utils"
)

// Handler for the  service
func MsvcStates(c *fiber.Ctx) error {
	id := c.Params(utils.ID)
	cityId := c.Params(utils.CITY_ID)

	url := utils.MSVC_STATES_URL

	if len(id) != 0 && url != "" {
		url = utils.MSVC_STATES_URL + "/" + id
	}
	if len(cityId) != 0 && url != "" {
		url = utils.MSVC_STATES_BY_CITY_URL + "/" + cityId
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
