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

	db := database.Init(conf.Database, true) // true for migration database
	web := routes.WebRoute{DB: db, App: app}
	api := routes.APIRoute{DB: db, App: app}
	web.SetupWebRoutes()
	api.SetupAPIRoutes()

	log.Fatal(app.Listen(":8000"))
}
