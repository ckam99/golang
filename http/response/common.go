package response

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Error struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Value     string `json:"value"`
}

func HttpResponseError(c *fiber.Ctx, status int, message string) error {
	log.Println(message)
	return c.Status(status).JSON(fiber.Map{
		"message": message,
	})
}
