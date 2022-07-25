package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

func RouteMiddleware(c *fiber.Ctx) error {
	// Return HTTP 404 status and JSON response.
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "sorry, endpoint is not found",
	})
}

func CorsMiddleware() func(ctx *fiber.Ctx) error {
	config := cors.Config{
		//AllowOrigins: "https://gofiber.io, https://gofiber.net",
		// AllowHeaders: "Origin, Content-Type, Accept",
		Next:             nil,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	}
	return cors.New(config)
}
