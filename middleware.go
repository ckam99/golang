package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func BearerAuthMiddleware() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey: []byte(secretKey),
		//AuthScheme: "Bearer", by default is "Bearer"
		ContextKey: "x-fiber-user", // used in private routes
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": err.Error(),
				})
			}
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	}
	return jwtware.New(config)
}

func RouteMiddleware(c *fiber.Ctx) error {
	// Return HTTP 404 status and JSON response.
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"error": true,
		"msg":   "sorry, endpoint is not found",
	})
}
