package routes

import (
	"app/handler"
	"app/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoutes(app *fiber.App) {
	helpers.SetSwaggerInfos()
	app.Get("/docs/*", swagger.HandlerDefault)

	// app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
	// 	URL:         "http://localhost:8000/docs/doc.json",
	// 	DeepLinking: false,
	// 	// Expand ("list") or Collapse ("none") tag groups by default
	// 	DocExpansion: "none",
	// 	// Prefill OAuth ClientId on Authorize popup
	// 	OAuth: &swagger.OAuthConfig{
	// 		AppName:  "OAuth Provider",
	// 		ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
	// 	},
	// 	// Ability to change OAuth2 redirect uri location
	// 	OAuth2RedirectUrl: "http://localhost:8080/swagger2/oauth2-redirect.html",
	// }))
}

func TodoRoutes(app *fiber.App) {
	app.Get("/", handler.HealthCheck)
	app.Get("/todos", handler.GetTodosHandler)
	app.Post("/todos", handler.CreateTodoHandler)
	app.Get("/todos/:todo_id", handler.GetTodoHandler)
}
