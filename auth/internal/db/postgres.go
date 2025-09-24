package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Connect initializes a pgxpool.Pool with sane defaults
func connectInternal(dbURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	// tune defaults (optional, adjust for your load)
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute

	// initialize pool
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// test connection
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}

var pool *pgxpool.Pool

func Connect(dbURL string) (*pgxpool.Pool, error) {
	var err error
	pool, err = connectInternal(dbURL)
	return pool, err
}

func GetPool() *pgxpool.Pool {
	return pool
}
