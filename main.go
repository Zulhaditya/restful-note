package main

import (
	"github.com/Zulhaditya/restful-note/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// start a new fiber app
	app := fiber.New()

	// connect to database
	database.ConnectDB()

	// send a string for get calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("APIs is running...")
		return err
	})

	// listen on port 3000
	app.Listen(":3000")
}
