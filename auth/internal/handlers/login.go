package handlers

import (
	"context"
	"github.com/satya-sudo/go-url/auth/internal/db"
	"github.com/satya-sudo/go-url/auth/internal/utils"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Login handles user login and returns a JWT
func Login(c *fiber.Ctx) error {
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

	// fetch user by email
	var userId string
	var hash string
	var role string

	err := db.GetPool().QueryRow(context.Background(),
		`SELECT id, password_hash, role FROM users WHERE email=$1`, body.Email,
	).Scan(&userId, &hash, &role)

	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "invalid email or password")
	}

	// check password
	if !utils.CheckPassword(hash, body.Password) {
		return fiber.NewError(http.StatusUnauthorized, "invalid email or password")
	}

	// parse expiry (from env, default 15m)
	expiryStr := os.Getenv("JWT_EXPIRES")
	if expiryStr == "" {
		expiryStr = "15m"
	}
	expiry, _ := time.ParseDuration(expiryStr)

	// issue JWT
	token, err := utils.GenerateJWT(os.Getenv("JWT_SECRET"), userId, role, expiry)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to generate token")
	}

	return c.JSON(fiber.Map{
		"accessToken": token,
		"expiresIn":   int(expiry.Seconds()),
		"userId":      userId,
		"role":        role,
	})
}
