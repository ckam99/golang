package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ckam225/golang/fiber/internal/config"
	"github.com/ckam225/golang/fiber/internal/database"
	"github.com/ckam225/golang/fiber/internal/http/middleware"
	"github.com/ckam225/golang/fiber/internal/jobs"
	"github.com/ckam225/golang/fiber/internal/routes"

	"github.com/gofiber/fiber/v2"
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
	// load environment variable
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	// http server
	app := fiber.New(fiber.Config{
		Views: conf.Server.HtmlEngine,
		// ViewsLayout: "layouts/base",
	})
	// Database
	db := database.Init(conf.Database, true) // true for migration database
	// Routes
	routes.SetupWebRoutes(app, db)
	routes.SetupAPIRoutes(app, db)
	// middleware
	app.Use(middleware.TestMiddleware)
	app.Use(middleware.CorsMiddleware())
	//app.Use(middleware.RouteMiddleware)

	// jobs
	jobs.RegisterNotificationChannel()

	defer jobs.UnregisterNotificationChannel()

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
