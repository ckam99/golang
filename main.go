package main

import (
	"log"

	"app/middleware"
	"app/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()
	// Middleware
	app.Use(recover.New())
	app.Use(middleware.CorsMiddleware())
	// routes
	routes.SwaggerRoutes(app)
	routes.TodoRoutes(app)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
