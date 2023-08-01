package server

import (
	"ddd2/app/routers"

	"github.com/gofiber/fiber/v2"
)

func Server() {
	app := fiber.New()
	routers.MoviesR(app)
	routers.Routs(app)
	app.Listen(":8888")
}
