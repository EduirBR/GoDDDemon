package main

import (
	"ddd2/app/routers"
	"ddd2/app/server"
)

func main() {

	app := server.Server()
	routers.MoviesR(app)
	routers.Routs(app)
	app.Listen(":8888")

}
