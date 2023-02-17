package main

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"id-provider/handler"
	"id-provider/middleware"
	"id-provider/oauth"
	"log"
)

func main() {
	app := fiber.New()
	srv := oauth.New(&oauth.Config{})

	app.Use(adaptor.HTTPMiddleware(middleware.LogMiddleware))

	router := handler.NewRouter(srv, app)
	router.Register()

	log.Fatal(app.Listen(":7000"))
}
