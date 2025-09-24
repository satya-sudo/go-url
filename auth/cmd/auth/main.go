package main

import (
	"log"

	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	// init DB connection
	pool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	defer pool.Close()

	app := fiber.New()

	// setup routes
	router.Setup(app)

	log.Printf("Auth service listening on %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
