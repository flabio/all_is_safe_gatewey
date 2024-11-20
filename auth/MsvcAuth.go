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

		return c.Status(fiber.StatusBadRequest).Send([]byte(respBody))
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(respBody, &dataMap)
	if err != nil {

		return c.Status(fiber.StatusBadRequest).Send([]byte(respBody))
	}

	t, err := GenerateToken(dataMap)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	log.Println(string(t))
	return c.Send([]byte(t))
}

func GenerateToken(dataMap map[string]interface{}) (string, error) {
	roleMap := dataMap["Role"].(map[string]interface{})
	idFloat64 := roleMap["id"].(float64)
	roleId := int(idFloat64)
	roleName := roleMap["name"].(string)
	roleModuleListInterface := roleMap["role_module"]
	roleModuleList := roleModuleListInterface.([]interface{})
	var roleModules []dto.RoleModule
	for _, roleModuleItem := range roleModuleList {
		roleModuleMap, ok := roleModuleItem.(map[string]interface{})
		if !ok {
			continue
		}
		roleModuleIdFloat64, ok := roleModuleMap["id"].(float64)
		if !ok {
			continue
		}
		roleModuleId := int(roleModuleIdFloat64)
		roleModuleActive, ok := roleModuleMap["active"].(bool)
		if !ok {
			continue
		}

		moduleInterface := roleModuleMap["module"]
		moduleMap := moduleInterface.(map[string]interface{})
		moduleIdFloat64 := moduleMap["id"].(float64)
		moduleId := uint(moduleIdFloat64)

		moduleName, ok := moduleMap["name"].(string)
		moduleIcon, ok := moduleMap["icon"].(string)
		moduleOrderFloat64, ok := moduleMap["order"].(float64)
		moduleOrder := int(moduleOrderFloat64)

		moduleActive, ok := moduleMap["active"].(bool)
		moduleRole := moduleMap["module_role"]
		module := dto.Module{
			Id:         moduleId,
			Name:       moduleName,
			Icon:       moduleIcon,
			Order:      moduleOrder,
			Active:     moduleActive,
			ModuleRole: moduleRole,
		}

		// Crear el objeto `RoleModule`
		roleModule := dto.RoleModule{
			Id:     roleModuleId,
			Active: roleModuleActive,
			Module: module,
		}

		// Agregar el `roleModule` al slice
		roleModules = append(roleModules, roleModule)
	}

	log.Println(roleModules)
	role := dto.Role{
		Id:          roleId,
		Name:        roleName,
		RoleModules: roleModules,
	}
	userData := dto.UserDTO{
		Id:        int(dataMap["Id"].(float64)),
		Avatar:    dataMap["Avatar"].(string),
		FirstName: dataMap["FirstName"].(string),
		LastName:  dataMap["LastName"].(string),
		Email:     dataMap["Email"].(string),
		Role:      role,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["full_name"] = userData.FirstName + " " + userData.LastName
	claims["email"] = userData.Email
	claims["avatar"] = userData.Avatar
	claims["rol"] = userData.Role
	claims["id"] = userData.Id
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	t, err := token.SignedString([]byte("supersecretkey"))
	return t, err

}
