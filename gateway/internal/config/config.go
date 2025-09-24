package config

import "os"

type Config struct {
	Port            string
	AuthService     string
	CrudService     string
	JWTSecret       string
	RedirectService string
}

func Load() Config {
	return Config{
		Port:            getEnv("GATEWAY_PORT", "8080"),
		AuthService:     getEnv("AUTH_SERVICE", "http://auth:8081"),
		CrudService:     getEnv("CRUD_SERVICE", "http://crudService:8082"),
		JWTSecret:       getEnv("JWT_SECRET", "supersecret"),
		RedirectService: getEnv("REDIRECT_SERVICE", "http://redirectService:8083"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
