package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupWebRoutes(app *fiber.App, db *gorm.DB) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome", fiber.Map{
			"Title": "Hello, <b>World</b>!",
		})
	})

	app.Get("about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{})
	})
}
