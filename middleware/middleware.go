package middleware

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func TestMiddleware(c *fiber.Ctx) error {
	// Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	// Go to next middleware:
	return c.Next()
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("log middleware", r.Method, r.Host, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
