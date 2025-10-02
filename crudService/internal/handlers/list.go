package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/satya-sudo/go-url/crudService/internal/db"
	"github.com/satya-sudo/go-url/crudService/internal/models"
	"net/http"
)

// List handles GET /shorten/all
func List(c *fiber.Ctx) error {
	// user id injected by Gateway
	userIdStr := c.Get("X-User-Id")
	if userIdStr == "" {
		return fiber.NewError(http.StatusUnauthorized, "missing user identity")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "invalid user identity")
	}

	rows, err := db.GetPool().Query(
		context.Background(),
		`SELECT short_code, long_url, user_id, created_at, expiration_at, hits 
         FROM urls
         WHERE user_id=$1
         ORDER BY created_at DESC`,
		userId,
	)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to fetch urls")
	}
	defer rows.Close()

	var urls []models.URL
	for rows.Next() {
		var u models.URL
		if err := rows.Scan(
			&u.ShortCode,
			&u.LongURL,
			&u.UserID,
			&u.CreatedAt,
			&u.ExpirationAt,
			&u.Hits,
		); err != nil {
			return fiber.NewError(http.StatusInternalServerError, "failed to parse url row")
		}
		urls = append(urls, u)
	}

	return c.Status(http.StatusOK).JSON(urls)
}
