package noteRoutes

import (
	noteHandler "github.com/Zulhaditya/restful-note/internal/handlers/note"
	"github.com/gofiber/fiber/v2"
)

func SetupNoteRoutes(router fiber.Router) {
	note := router.Group("/note")

	// create a note
	note.Post("/", noteHandler.CreateNotes)

	// read all notes
	note.Get("/", noteHandler.GetNotes)

	// read one note by id
	note.Get("/:noteId", noteHandler.GetNote)

	// update one note
	note.Put("/:noteId", noteHandler.UpdateNote)

	// delete one note
	note.Delete("/:noteId", noteHandler.DeleteNote)

}
