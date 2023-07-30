package server

import (
	"github.com/gofiber/fiber/v2"
)

func Server() *fiber.App {
	app := fiber.New()
	return app
}
