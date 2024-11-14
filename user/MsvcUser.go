package user

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	"net/http"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/utils"
)

// Handler for the  service
func MsvcUser(namePath string, c *fiber.Ctx) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(utils.ENV_ERROR)
	}
	url := os.Getenv("MSVC_USER_URL")
	pageParam := c.Query(utils.PAGE)
	id := c.Params(utils.ID)
	if len(id) != 0 && url != "" {
		url = url + "/" + id
	} else {
		url = os.Getenv("MSVC_USER_URL") + "?page=" + pageParam
	}

	if namePath != "" {
		url = os.Getenv("MSVC_USER_URL") + "/" + namePath + "/" + "?page=" + pageParam
	}
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFields := []string{
		constants.FIRST_NAME,
		constants.LAST_NAME,
		constants.ADDRESS,
		constants.PHONE,
		constants.STATE_ID,
		constants.ROL_ID,
		constants.EMAIL,
		constants.PASSWORD,
		"password_confirmation",
		constants.ACTIVE,
	}
	for _, field := range formFields {
		value := c.FormValue(field)
		if value != "" {
			err := writer.WriteField(field, value)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString(utils.FIELD_FORM)
			}
		}
	}
	fileHeader, err := c.FormFile(utils.FILE)
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_OPEN)
		}
		defer file.Close()
		part, err := writer.CreateFormFile(utils.FILE, fileHeader.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_CREATE)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_COPY_ERROR)
		}
	}
	err = writer.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_WRITER_ERROR)
	}
	req, err := http.NewRequest(c.Method(), url, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_REQUEST)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.FILE_ERROR_SERVICE)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_REQUEST_SERVICE)
	}
	return c.Send(respBody)
}

func MsvcUserAvatar(c *fiber.Ctx) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(utils.ENV_ERROR)
	}
	url := os.Getenv("MSVC_USER_URL")
	id := c.Params(utils.ID)
	url = url + "/avatar/" + id
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileHeader, err := c.FormFile(utils.FILE)
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_OPEN)
		}
		defer file.Close()
		part, err := writer.CreateFormFile(utils.FILE, fileHeader.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_CREATE)
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_COPY_ERROR)
		}
	}
	err = writer.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_WRITER_ERROR)
	}
	req, err := http.NewRequest(c.Method(), url, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_REQUEST)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.FILE_ERROR_SERVICE)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FILE_ERROR_REQUEST_SERVICE)
	}
	return c.Send(respBody)
}

func MsvcUserUpdatePassword(c *fiber.Ctx) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf(utils.ENV_ERROR)
	}
	url := os.Getenv("MSVC_USER_URL")
	id := c.Params(utils.ID)

	url = url + "/password/" + id
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
