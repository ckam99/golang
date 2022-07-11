package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// JWT Middleware

	app.Get("/", WelcomeHandler)

	app.Post("/login", LoginHandler)

	app.Post("/auth/token", TestTokenHandler)

	app.Get("/auth/me", BearerAuthMiddleware(), CurrentUserHandler)
	app.Get("/auth/user", BearerAuthMiddleware(), CurrentUserHandler2)

	app.Get("/protected", BearerAuthMiddleware())
	app.Use(RouteMiddleware)

	log.Fatal(app.Listen(":8000"))

}
