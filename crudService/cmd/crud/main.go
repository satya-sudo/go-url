package main

import (
	"github.com/satya-sudo/go-url/crudService/internal/config"
	"github.com/satya-sudo/go-url/crudService/internal/db"
	"github.com/satya-sudo/go-url/crudService/internal/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.Load()

	// connect DB
	pool, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect DB: %v", err)
	}
	defer pool.Close()

	app := fiber.New()

	// setup routes
	router.Setup(app)

	log.Printf("CRUD service listening on :%s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
