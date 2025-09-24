package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	RedisAddr   string
	RedisPass   string
}

// Load reads config from env
func Load() Config {
	return Config{
		Port:        getEnv("REDIRECT_PORT", "8083"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@db:5432/shortener"),
		RedisAddr:   getEnv("REDIS_ADDR", "redis:6379"),
		RedisPass:   getEnv("REDIS_PASS", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
