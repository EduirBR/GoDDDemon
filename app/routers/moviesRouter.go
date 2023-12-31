package routers

import (
	"ddd2/app/filters"
	"ddd2/domain/services"

	"github.com/gofiber/fiber/v2"
)

func MoviesR(app *fiber.App) {
	group := app.Group("/path")
	group.Get("", filters.GetMovies)
	group.Post("", filters.CreateMovie)
	group.Get("/create", func(c *fiber.Ctx) error {

		services.CreateTableMovies()

		return nil
	})
}
