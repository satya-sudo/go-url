package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

// connectInternal builds a pgxpool.Pool with defaults
func connectInternal(dbURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}

	// tune connection pool settings
	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	p, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	// test connection
	if err := p.Ping(ctx); err != nil {
		return nil, err
	}

	return p, nil
}

// Connect initializes the global pool
func Connect(dbURL string) (*pgxpool.Pool, error) {
	var err error
	pool, err = connectInternal(dbURL)
	return pool, err
}

// GetPool returns the global pool
func GetPool() *pgxpool.Pool {
	return pool
}
