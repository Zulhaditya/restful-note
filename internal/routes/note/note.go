package noteRoutes

import "github.com/gofiber/fiber/v2"

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	// create a note
	note.Post("/", func(c *fiber.Ctx) error {})

	// read all notes
	note.Get("/", func(c *fiber.Ctx) error {})

	// read one note
	note.Get("/:noteId", func(c *fiber.Ctx) error {})

	// update one note
	note.Put("/:noteId", func(c *fiber.Ctx) error {})

	// delete one note
	note.Delete("/:noteId", func(c *fiber.Ctx) error {})
}
