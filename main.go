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
	app := fiber.New(fiber.Config{})
	// middleware
	app.Use(middleware.TestMiddleware)

	if conf, err := config.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	} else {
		db := database.Init(conf.Database, true) // true for migration database
		web := routes.WebRoute{DB: db, App: app}
		api := routes.APIRoute{DB: db, App: app}
		web.SetupWebRoutes()
		api.SetupAPIRoutes()

	}

	log.Fatal(app.Listen(":8000"))
}
