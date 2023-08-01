package filters

import (
	"ddd2/domain/services"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	c.Set("Context-Type", "applicaction/json")
	objs, err := services.ListMovies()
	if err != nil {

	}
	c.Status(fiber.StatusOK).JSON(objs)
	return nil
}

func CreateMovie(c *fiber.Ctx) error {
	obj := fiber.Map{}
	c.BodyParser(&obj)
	objJSON, _ := json.Marshal(obj)
	fmt.Println(objJSON)
	services.CreateMovie(objJSON)
	return nil
}
