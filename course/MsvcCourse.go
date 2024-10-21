package course

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/safe/utils"
)

// Handler for the  service
func MsvcCourse(namePath string, c *fiber.Ctx) error {
	id := c.Params(utils.ID)
	url := utils.MSVC_COURSE_URL

	if len(id) != 0 && url != "" {
		url = utils.MSVC_COURSE_URL + "/" + id
	}

	if namePath != "" {
		url = utils.MSVC_COURSE_URL + "/" + namePath
	}
	if namePath != "" && id != "" {
		url = utils.MSVC_COURSE_URL + "/" + namePath + "/" + id
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
