package module

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

func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
func MsModule(c *fiber.Ctx) error {
	EnvLoad()
	MSVC_MODULE_URL := os.Getenv("MSVC_MODULE_URL")
	id := c.Params(utils.ID)

	if len(id) != 0 && MSVC_MODULE_URL != "" {
		MSVC_MODULE_URL = MSVC_MODULE_URL + "/" + id
	}
	req, err := http.NewRequest(c.Method(), MSVC_MODULE_URL, bytes.NewBuffer(c.Body()))
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

func MsModuleRole(c *fiber.Ctx) error {
	EnvLoad()
	MSVC_MODULE_ROLE_URL := os.Getenv("MSVC_MODULE_ROLE_URL")
	id := c.Params(utils.ID)

	if len(id) != 0 && MSVC_MODULE_ROLE_URL != "" {
		MSVC_MODULE_ROLE_URL = MSVC_MODULE_ROLE_URL + "/" + id
	}
	req, err := http.NewRequest(c.Method(), MSVC_MODULE_ROLE_URL, bytes.NewBuffer(c.Body()))
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
