package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func BearerAuthMiddleware() func(*fiber.Ctx) error {
	config := jwtware.Config{
		SigningKey: []byte(os.Getenv("SECRET_KEY")),
		//AuthScheme: "Bearer", by default is "Bearer"
		ContextKey: "x-agent-user", // used in private routes
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
