package router

import (
	"github.com/satya-sudo/go-url/redirectService/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// Setup initializes all routes for the Redirect service
func Setup(app *fiber.App) {
	// Public route: redirect by short code
	app.Get("/:code", handlers.Redirect)
}
