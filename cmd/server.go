package main

import (
	"github.com/ckam225/golang/echo/internal/routes"
	"github.com/ckam225/golang/echo/internal/schemas"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()
	app.Validator = &schemas.CustomValidator{Validator: validator.New()}
	routes.SetupWebRoutes(app)
	routes.SetupApiRoutes(app)
	app.Logger.Fatal(app.Start(":8000"))
}
