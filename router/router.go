package router

import (
	noteRoutes "github.com/Zulhaditya/restful-note/internal/routes/note"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	// setup note routes
	noteRoutes.SetupNoteRoutes(api)
}
