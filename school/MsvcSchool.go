package school

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/safe/utils"
)

func MsvcSchool(c *fiber.Ctx) error {
	pageParam := c.Query(utils.PAGE)
	id := c.Params(utils.ID)
	url := utils.MSVC_SCHOOL_URL

	if len(id) != 0 && url != "" {
		url = utils.MSVC_SCHOOL_URL + "/" + id
	}
	if len(pageParam) > 0 && url != "" {
		url = url + "?page=" + pageParam
	}
	if url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("URL vacía")
	}

	// Crear un nuevo buffer para almacenar el multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Procesar los campos del formulario
	formFields := []string{"name", "email", "address", "phone", "zip_code", "provider_number", "state_id"}
	for _, field := range formFields {
		value := c.FormValue(field)
		if value != "" {
			err := writer.WriteField(field, value)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).SendString("Error al escribir campo de formulario")
			}
		}
	}

	// Procesar el archivo subido
	fileHeader, err := c.FormFile("file")

	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al abrir el archivo")
		}
		defer file.Close()

		part, err := writer.CreateFormFile("file", fileHeader.Filename)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al crear archivo en el multipart")
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error al copiar el archivo")
		}

	}

	// Cerrar el writer
	err = writer.Close()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al cerrar el multipart writer")
	}

	// Crear la solicitud al servicio externo
	req, err := http.NewRequest(c.Method(), url, body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al crear la solicitud")
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString("El servicio no está disponible")
	}
	defer resp.Body.Close()

	// Leer la respuesta
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error al leer la respuesta del servicio")
	}

	return c.Send(respBody)
}

/*func MsvcSchool(c *fiber.Ctx) error {
	id := c.Params(utils.ID)
	url := utils.MSVC_SCHOOL_URL

	if len(id) != 0 && url != "" {
		url = utils.MSVC_SCHOOL_URL + "/" + id
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
}*/
