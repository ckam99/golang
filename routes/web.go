package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type WebRoute struct {
	DB  *gorm.DB
	App *fiber.App
}

func (w *WebRoute) SetupWebRoutes() {
	w.App.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome", fiber.Map{
			"Title": "Hello, <b>World</b>!",
		})
	})

	w.App.Get("about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{})
	})
}
