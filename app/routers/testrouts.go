package routers

import (
	"ddd2/domain/services"

	"github.com/gofiber/fiber/v2"
)

func Routs(app *fiber.App) {
	app.Get("/create", func(c *fiber.Ctx) error {

		services.CreateTableMovies()

		return nil
	})
}
