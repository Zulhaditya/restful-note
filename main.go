package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// start a new fiber app
	app := fiber.New()

	// send a string for get calls to the endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("APIs is running brother!...")
		return err
	})

	// listen on port 3000
	app.Listen(":3000")
}
