package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	JWTSecret   string
	JWTExpires  string // e.g. "15m"
}

// Load reads config from env with fallbacks
func Load() Config {
	return Config{
		Port:        getEnv("AUTH_PORT", "8081"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@db:5432/shortener"),
		JWTSecret:   getEnv("JWT_SECRET", "supersecret"),
		JWTExpires:  getEnv("JWT_EXPIRES", "15m"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
