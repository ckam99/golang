package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

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
