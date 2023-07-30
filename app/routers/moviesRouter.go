package routers

import (
	"ddd2/app/filters"

	"github.com/gofiber/fiber/v2"
)

func MoviesR(app *fiber.App) {
	group := app.Group("/path")
	group.Get("", filters.GetMovie)
	group.Post("", filters.CreateMovie)
}
