package main

import (
	"log"

	"app/handler"
	_ "app/handler"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main2() {
	app := fiber.New()
	// Middleware
	app.Use(recover.New())
	app.Use(middleware.CorsMiddleware())
	// routes
	app.Get("/", handler.HealthCheck)
	app.Get("/todos", handler.GetTodosHandler)
	app.Get("/todos/:todo_id", handler.GetTodoHandler)

	// Default method with annotation
	app.Get("/docs/*", swagger.HandlerDefault)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
