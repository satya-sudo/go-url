package db

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

// ConnectRedis initializes the global Redis client
func ConnectRedis(addr, pass string) (*redis.Client, error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0, // default DB
	})

	// test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}

// GetRedis returns the global Redis client
func GetRedis() *redis.Client {
	return rdb
}
