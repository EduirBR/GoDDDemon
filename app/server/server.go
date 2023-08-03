package server

import (
	"ddd2/app/routers"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func serverConfig() string {
	port := os.Getenv("PORT_SERVER")
	return fmt.Sprintf(":%v", port)
}

func RunServer() *fiber.App {
	app := fiber.New()
	routers.MoviesR(app)
	routers.ModelsRoutes(app)
	app.Listen(serverConfig())
	return app
}
