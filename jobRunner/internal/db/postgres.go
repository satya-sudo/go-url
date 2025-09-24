package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pool *pgxpool.Pool

func connectInternal(dbURL string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		return nil, err
	}
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

	if err := p.Ping(ctx); err != nil {
		return nil, err
	}
	return p, nil
}

func Connect(dbURL string) (*pgxpool.Pool, error) {
	var err error
	pool, err = connectInternal(dbURL)
	return pool, err
}

func GetPool() *pgxpool.Pool { return pool }
