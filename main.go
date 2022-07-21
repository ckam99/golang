package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"example/fiber/entity"
	"example/fiber/http/middleware"
	"example/fiber/routes"
	"example/fiber/service"
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

	users := []entity.User{
		{
			Email: "admin@example.com",
			Name:  "Claver Amon",
		},
		{
			Email: "admin@example.com",
			Name:  "Claver Amon",
		},
	}
	data := map[string]string{
		"name":  "PUSH",
		"email": "BNBB",
	}
	if err := service.NotifyUsers(&users, "Welcome", data, "mail/register.tmpl"); err != nil {
		fmt.Println(err.Error())
	}

	log.Fatal(app.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
