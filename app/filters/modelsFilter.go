package filters

import (
	"ddd2/domain/extras"
	"ddd2/domain/services"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetModels(c *fiber.Ctx) error {
	c.Set("Context-Type", "applicaction/json")
	objs, err := services.ListModels()
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	c.Status(fiber.StatusOK).JSON(objs)
	return nil
}

func CreateModel(c *fiber.Ctx) error {
	obj := fiber.Map{}
	c.BodyParser(&obj)
	objJSON, _ := json.Marshal(obj)
	services.CreateModel(objJSON)
	return nil
}
