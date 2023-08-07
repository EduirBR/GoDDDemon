package filters

import (
	"ddd2/app/extras"
	"ddd2/domain/services"
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMovies(c *fiber.Ctx) error {
	c.Set("Context-Type", "applicaction/json")
	objs, err := services.ListMovies()
	if err != nil {
		extras.Errors(extras.GetFunctionName(), err)
	}
	c.Status(fiber.StatusOK).JSON(objs)
	return nil
}

func CreateMovie(c *fiber.Ctx) error {
	obj := fiber.Map{}
	c.BodyParser(&obj)
	objJSON, _ := json.Marshal(obj)
	err := services.CreateMovie(objJSON)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	c.Status(fiber.StatusOK).SendString("Obj Created")
	return nil
}
func UpdateMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	obj := fiber.Map{}
	c.BodyParser(&obj)
	objJSON, _ := json.Marshal(obj)
	err = services.EditMovie(objJSON, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	c.Status(fiber.StatusOK).SendString("Obj Updated")
	return nil
}

func RetreiveMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	objs, _ := services.RetreiveMovie(id)
	c.Status(fiber.StatusOK).JSON(objs)
	return nil
}

func DestroyMovie(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}
	err = services.DeleteMovie(id)
	if err != nil {
		c.Status(fiber.StatusOK).SendString(err.Error())
	}
	c.Status(fiber.StatusOK).SendString("Obj destroyed")
	return nil
}
