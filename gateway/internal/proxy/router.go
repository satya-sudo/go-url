package proxy

import (
	"github.com/gin-contrib/cors"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/satya-sudo/go-url/gateway/internal/config"
)

func SetupRouter(cfg config.Config) *gin.Engine {
	r := gin.Default()
	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // for dev
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // for dev
		MaxAge:           12 * time.Hour,
	}))

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
		protected.GET("/list/all", reverseProxy(cfg.CrudService))
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
