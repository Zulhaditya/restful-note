package main

import (
	"github.com/Zulhaditya/restful-note/database"
	"github.com/Zulhaditya/restful-note/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// start a new fiber app
	app := fiber.New()

	// connect to database
	database.ConnectDB()

	// setup the router
	router.SetupRoutes(app)

	// listen on port 3000
	app.Listen(":3000")
}
