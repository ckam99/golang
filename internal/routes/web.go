package routes

import (
	"net/http"

	"github.com/ckam225/golang/echo/internal/schemas"
	"github.com/labstack/echo/v4"
)

func SetupWebRoutes(app *echo.Echo) {
	app.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, schemas.ErrorResponse{
			Message: "Welcome echo",
		})
	})
}
