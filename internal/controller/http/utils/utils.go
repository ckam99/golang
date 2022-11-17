package utils

import "github.com/gofiber/fiber/v2"

func HTTPMessage(msg string) fiber.Map {
	return fiber.Map{
		"message": msg,
	}
}

func HTTPError(c *fiber.Ctx, msg string, status int) error {
	return c.Status(status).JSON(HTTPMessage(msg))
}
