package handlers

import (
	"context"
	"github.com/satya-sudo/go-url/crudService/internal/db"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Delete handles DELETE /shorten/:id
func Delete(c *fiber.Ctx) error {
	code := c.Params("id")
	if code == "" {
		return fiber.NewError(http.StatusBadRequest, "shortCode is required")
	}

	// ensure user owns this url
	userIdStr := c.Get("X-User-Id")
	if userIdStr == "" {
		return fiber.NewError(http.StatusUnauthorized, "missing user identity")
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		return fiber.NewError(http.StatusUnauthorized, "invalid user identity")
	}

	res, err := db.GetPool().Exec(context.Background(),
		`DELETE FROM urls WHERE short_code=$1 AND user_id=$2`, code, userId,
	)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "failed to delete url")
	}
	if res.RowsAffected() == 0 {
		return fiber.NewError(http.StatusNotFound, "url not found or not owned by user")
	}

	return c.SendStatus(http.StatusNoContent)
}
