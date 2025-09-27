package handlers

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"github.com/satya-sudo/go-url/crudService/internal/db"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// generateCode creates a random short code (6 chars)
func generateCode() (string, error) {
	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// Create handles POST /shorten
func Create(c *fiber.Ctx) error {
	var body struct {
		LongURL      string  `json:"longUrl"`
		ExpirationAt *string `json:"expirationAt,omitempty"`
	}
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(http.StatusBadRequest, "invalid request body")
	}
	if body.LongURL == "" {
		return fiber.NewError(http.StatusBadRequest, "longUrl is required")
	}

	// user id injected by Gateway
	userIdStr := c.Get("X-User-Id")
	if userIdStr == "" {
		return fiber.NewError(http.StatusUnauthorized, "missing user identity")
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "invalid user identity")
	}

	code, err := generateCode()
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to generate code")
	}

	var expiration *time.Time
	if body.ExpirationAt != nil {
		t, err := time.Parse(time.RFC3339, *body.ExpirationAt)
		if err != nil {
			return fiber.NewError(http.StatusBadRequest, "invalid expirationAt format")
		}
		expiration = &t
	}

	_, err = db.GetPool().Exec(context.Background(),
		`INSERT INTO urls (short_code, long_url, user_id, expiration_at) VALUES ($1,$2,$3,$4)`,
		code, body.LongURL, userId, expiration,
	)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to insert url")
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"shortCode": code,
		"longUrl":   body.LongURL,
	})
}
