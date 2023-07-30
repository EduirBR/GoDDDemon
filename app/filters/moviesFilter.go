package filters

import (
	"ddd2/domain/services"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func GetMovie(c *fiber.Ctx) error {
	return nil
}

func CreateMovie(c *fiber.Ctx) error {
	obj := fiber.Map{}
	c.BodyParser(&obj)
	objJSON, _ := json.Marshal(obj)
	services.CreateMovie(objJSON)
	return nil
}
