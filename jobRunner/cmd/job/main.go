package main

import (
	"jobRunner/internal/config"
	"jobRunner/internal/db"
	"jobRunner/internal/jobs"
	"log"
	"time"
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

	log.Println("Job Runner started — syncing hits from Redis to Postgres")

	// Run sync every 1 minute
	jobs.StartSync(1 * time.Minute)
}
