package noteHandler

import (
	"github.com/Zulhaditya/restful-note/database"
	"github.com/Zulhaditya/restful-note/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note

	// find all notes in the database
	db.Find(&notes)

	// if no note is present returns an error
	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes present", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes found", "data": notes})
}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	// store the body in the note and return error if encountered
	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// add uuid to the note
	note.ID = uuid.New()

	// create the note
	err = db.Create(&note).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "failed to create a note", "data": err})
	}

	// return the created note
	return c.JSON(fiber.Map{"status": "success", "message": "Created note", "data": note})
}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// read the param noteId
	id := c.Params("noteId")

	// find the note with the given id
	db.Find(&note, "id = ?", id)

	// if no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// return note with the id
	return c.JSON(fiber.Map{"status": "success", "message": "Notes found", "data": note})
}

func UpdateNote(c *fiber.Ctx) error {
	type updateNote struct {
		Title    string `json:"title"`
		Subtitle string `json:"subtitle"`
		Text     string `json:"text"`
	}

	db := database.DB
	var note model.Note

	// read the param noteId
	id := c.Params("noteId")

	// find the note with given id
	db.Find(&note, "id = ?", id)

	// if no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// store the body containing the updated data
	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// edit the note
	note.Title = updateNoteData.Title
	note.Subtitle = updateNoteData.Subtitle
	note.Text = updateNoteData.Text

	// save the changes to database
	db.Save(&note)

	// return the updated note
	return c.JSON(fiber.Map{"status": "success", "message": "Notes found", "data": note})
}

func DeleteNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	// read the param noteId
	id := c.Params("noteId")

	// find the not with given id
	db.Find(&note, "id = ?", id)

	// if no such note present return an error
	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No note present", "data": nil})
	}

	// delete the note and return error if encountered
	err := db.Delete(&note, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete note", "data": nil})
	}

	// return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted note"})
}
