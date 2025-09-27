package router

import (
	"github.com/satya-sudo/go-url/crudService/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// Setup initializes all CRUD routes
func Setup(app *fiber.App) {
	// Protected routes (Gateway already validated JWT & injected user headers)
	shorten := app.Group("/shorten")

	shorten.Post("/", handlers.Create)        // POST /shorten
	shorten.Delete("/:id", handlers.Delete)   // DELETE /shorten/:id
	shorten.Get("/:id/stats", handlers.Stats) // GET /shorten/:id/stats
}
