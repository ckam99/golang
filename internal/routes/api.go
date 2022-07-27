package routes

import (
	"github.com/ckam225/golang/echo/internal/handlers"
	"github.com/labstack/echo/v4"
)

func SetupApiRoutes(app *echo.Echo) {

	users := app.Group("/users")
	users.GET("/", handlers.GetUsersHandler)
	users.POST("/", handlers.CreateUserHandler)
	users.GET("/:id", handlers.GetUserHandler)
	users.DELETE("/:id", handlers.DeleteUserHandler)
}
