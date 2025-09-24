package proxy

import (
	"net/http/httputil"
	"net/url"

	"gateway/internal/config"
	"github.com/gin-gonic/gin"
)

func SetupRouter(cfg config.Config) *gin.Engine {
	r := gin.Default()

	// 🔓 Public endpoints
	r.Any("/auth/signup", reverseProxy(cfg.AuthService))
	r.Any("/auth/login", reverseProxy(cfg.AuthService))

	// 🔒 Protected: /auth/me (gateway enforces JWT, forwards user info)
	authGroup := r.Group("/auth")
	authGroup.Use(AuthMiddleware(cfg.JWTSecret))
	{
		authGroup.GET("/me", reverseProxy(cfg.AuthService))
	}

	// 🔓 Public redirect service
	r.GET("/:code", reverseProxy(cfg.RedirectService))

	// 🔒 Protected CRUD service
	protected := r.Group("/shorten")
	protected.Use(AuthMiddleware(cfg.JWTSecret))
	{
		protected.POST("/", reverseProxy(cfg.CrudService))
		protected.DELETE("/:id", reverseProxy(cfg.CrudService))
		protected.GET("/:id/stats", reverseProxy(cfg.CrudService))
	}

	return r
}

func reverseProxy(target string) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := url.Parse(target)
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
