package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
}

// Load reads environment variables with sensible defaults
func Load() Config {
	return Config{
		Port:        getEnv("CRUD_PORT", "8082"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@db:5432/shortener"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
