package course

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/utils"
)

func MsvcTopic(course string, c *fiber.Ctx) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	id := c.Params(utils.ID)
	courseId := c.Params("course_id")

	url := "http://localhost:3007/api/topic/"

	if len(id) != 0 && url != "" {
		url += id
	}
	if course != "" {
		log.Println(courseId)
		url += "course/" + courseId
		log.Println(url)
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
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to read response body")
	}
	return c.Send(respBody)
}
