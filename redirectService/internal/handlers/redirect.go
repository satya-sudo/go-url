package handlers

import (
	"context"
	"fmt"
	"github.com/satya-sudo/go-url/redirectService/internal/db"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// Redirect handles GET /:code
func Redirect(c *fiber.Ctx) error {
	code := c.Params("code")
	if code == "" {
		return fiber.NewError(http.StatusBadRequest, "shortCode required")
	}

	ctx := context.Background()
	rdb := db.GetRedis()
	pg := db.GetPool()

	// 🔹 1. Try Redis cache
	cacheKey := fmt.Sprintf("url:%s", code)
	longURL, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// increment counter in Redis (async)
		rdb.Incr(ctx, fmt.Sprintf("hits:%s", code))
		return c.Redirect(longURL, http.StatusFound)
	}

	if err != redis.Nil {
		// real Redis error (not just key missing)
		return fiber.NewError(http.StatusInternalServerError, "redis error")
	}

	// 🔹 2. Fallback to Postgres
	var url string
	err = pg.QueryRow(ctx,
		`UPDATE urls 
		 SET hits = hits + 1 
		 WHERE short_code = $1 
		 RETURNING long_url`, code,
	).Scan(&url)

	if err != nil {
		return fiber.NewError(http.StatusNotFound, "url not found")
	}

	// 🔹 3. Cache in Redis
	rdb.Set(ctx, cacheKey, url, 24*time.Hour) // expire in 1 day
	rdb.Incr(ctx, fmt.Sprintf("hits:%s", code))

	return c.Redirect(url, http.StatusFound)
}
