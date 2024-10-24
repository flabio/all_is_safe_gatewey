package user

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/safe/utils"
)

// Handler for the  service
func MsvcUser(namePath string, c *fiber.Ctx) error {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	url := os.Getenv("MSVC_USER_URL")
	pageParam := c.Query(utils.PAGE)

	id := c.Params(utils.ID)

	if len(id) != 0 && url != "" {
		url = url + "/" + id
	} else {
		url = url + "?page=" + pageParam
	}
	if namePath != "" {
		url = url + "/" + namePath + "/" + "?page=" + pageParam
	}
	// Crear un nuevo buffer para almacenar el multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	// Procesar los campos del formulario
	formFields := []string{"first_name", "first_sur_name", "secon_sur_name", "address",
		"phone",
		"zip_code",
		"state_id",
		"rol_id",
		"email",
		"password",
		"active"}
	for _, field := range formFields {
		value := c.FormValue(field)
		log.Println(value)
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

// 	req, err := http.NewRequest(c.Method(), url, bytes.NewBuffer(c.Body()))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(utils.FAILED_CREATE)
// 	}
// 	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.SERVICE_NOT_AVAILALE)
// 	}
// 	defer resp.Body.Close()
// 	respBody, _ := ioutil.ReadAll(resp.Body)

// 	return c.Send(respBody)
// }
