package handlers

import (
	"context"
	"github.com/satya-sudo/go-url/crudService/internal/db"
	"github.com/satya-sudo/go-url/crudService/internal/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// Stats handles GET /shorten/:id/stats
func Stats(c *fiber.Ctx) error {
	code := c.Params("id")
	if code == "" {
		return fiber.NewError(http.StatusBadRequest, "shortCode is required")
	}

	var url models.URL
	err := db.GetPool().QueryRow(context.Background(),
		`SELECT short_code, long_url, user_id, created_at, expiration_at, hits 
		 FROM urls WHERE short_code=$1`, code,
	).Scan(&url.ShortCode, &url.LongURL, &url.UserID, &url.CreatedAt, &url.ExpirationAt, &url.Hits)

	if err != nil {
		return fiber.NewError(http.StatusNotFound, "url not found")
	}

	return c.JSON(url)
}
