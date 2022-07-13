package main

import (
	"fmt"
	"log"
	"project-struct/config"
	"project-struct/internal/database"
	"project-struct/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	app := fiber.New(fiber.Config{
		Views: conf.Server.TemplateEngine,
	})
	db := database.Init(&conf.Db, true)
	routes.SetupWebRoutes(app, db)
	routes.SetupAPIRoutes(app, db)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", conf.Server.Port)))
}
