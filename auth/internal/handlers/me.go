package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/satya-sudo/go-url/auth/internal/db"
	"net/http"
)

// Me returns info about the current authenticated user
func Me(c *fiber.Ctx) error {
	userId := c.Get("X-User-Id")
	role := c.Get("X-User-Role")

	if userId == "" {
		return fiber.NewError(http.StatusUnauthorized, "missing user identity")
	}

	// optionally fetch email from DB
	var email string
	err := db.GetPool().QueryRow(context.Background(),
		`SELECT email FROM users WHERE id=$1`, userId,
	).Scan(&email)
	if err != nil {
		return c.JSON(fiber.Map{
			"userId": userId,
			"role":   role,
		})
	}

	return c.JSON(fiber.Map{
		"userId": userId,
		"email":  email,
		"role":   role,
	})
}
