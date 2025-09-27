package handlers

import (
	"context"
	"github.com/satya-sudo/go-url/auth/internal/db"
	"github.com/satya-sudo/go-url/auth/internal/models"
	"github.com/satya-sudo/go-url/auth/internal/utils"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Signup handles user registration
func Signup(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}

	if body.Email == "" || body.Password == "" {
		return fiber.NewError(http.StatusBadRequest, "email and password are required")
	}

	// hash password
	hash, err := utils.HashPassword(body.Password)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to hash password")
	}

	// create user struct
	user := models.User{
		ID:           uuid.New(),
		Email:        body.Email,
		PasswordHash: hash,
		Role:         "user",
	}

	// insert into DB
	conn := db.GetPool() // we'll add this helper in db/postgres.go
	_, err = conn.Exec(context.Background(),
		`INSERT INTO users (id, email, password_hash, role) VALUES ($1, $2, $3, $4)`,
		user.ID, user.Email, user.PasswordHash, user.Role,
	)
	if err != nil {
		// could be duplicate email
		return fiber.NewError(http.StatusConflict, "email already exists")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"userId": user.ID.String(),
		"email":  user.Email,
		"role":   user.Role,
	})
}
