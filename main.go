package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"example/fiber/http/middleware"
	"example/fiber/routes"
	"fmt"
	"log"
	"os"

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
	app := fiber.New(fiber.Config{
		Views: conf.Server.HtmlEngine,
		// ViewsLayout: "layouts/base",
	})

	db := database.Init(conf.Database, true) // true for migration database

	routes.SetupWebRoutes(app, db)
	routes.SetupAPIRoutes(app, db)

	// middleware
	app.Use(middleware.TestMiddleware)
	app.Use(middleware.CorsMiddleware())
	//app.Use(middleware.RouteMiddleware)

	fmt.Println(os.Getenv("APP_ENV"))

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
