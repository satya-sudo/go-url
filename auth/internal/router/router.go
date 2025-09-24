package router

import (
	"auth/internal/handlers"

	"github.com/gofiber/fiber/v2"
)

// Setup initializes all routes for the Auth service
func Setup(app *fiber.App) {
	// Public endpoints
	app.Post("/auth/signup", handlers.Signup)
	app.Post("/auth/login", handlers.Login)

	// Protected endpoints (uses JWT header directly in handlers for now)
	app.Get("/auth/me", handlers.Me)
}
