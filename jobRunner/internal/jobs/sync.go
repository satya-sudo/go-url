package jobs

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/satya-sudo/go-url/jobRunner/internal/db"
)

// SyncHits moves counters from Redis -> Postgres
func SyncHits() {
	ctx := context.Background()
	rdb := db.GetRedis()
	pg := db.GetPool()

	iter := rdb.Scan(ctx, 0, "hits:*", 0).Iterator()
	for iter.Next(ctx) {
		key := iter.Val()
		code := strings.TrimPrefix(key, "hits:")

		val, err := rdb.Get(ctx, key).Result()
		if err != nil {
			continue
		}
		count, _ := strconv.ParseInt(val, 10, 64)

		// Update DB
		_, err = pg.Exec(ctx,
			`UPDATE urls SET hits = hits + $1 WHERE short_code = $2`,
			count, code,
		)
		if err != nil {
			fmt.Printf("failed to update %s: %v\n", code, err)
			continue
		}

		// Reset counter
		rdb.Del(ctx, key)
	}

	if err := iter.Err(); err != nil {
		fmt.Printf("redis scan error: %v\n", err)
	}
}

// StartSync runs SyncHits on an interval
func StartSync(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		SyncHits()
	}
}
