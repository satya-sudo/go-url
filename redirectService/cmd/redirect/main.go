package main

import (
	"log"

	"redirectService/internal/config"
	"redirectService/internal/db"
	"redirectService/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	// Connect Postgres
	pg, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect Postgres: %v", err)
	}
	defer pg.Close()

	// Connect Redis
	rdb, err := db.ConnectRedis(cfg.RedisAddr, cfg.RedisPass)
	if err != nil {
		log.Fatalf("failed to connect Redis: %v", err)
	}
	defer rdb.Close()

	app := fiber.New()

	// setup routes
	router.Setup(app)

	log.Printf("Redirect service listening on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
