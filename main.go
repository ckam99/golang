package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"example/fiber/http/middleware"
	"example/fiber/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

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

	// middleware
	app.Use(middleware.TestMiddleware)
	app.Use(middleware.RouteMiddleware)

	db := database.Init(conf.Database, true) // true for migration database

	routes.SetupWebRoutes(app, db)
	routes.SetupAPIRoutes(app, db)

	log.Fatal(app.Listen(":8000"))
}
