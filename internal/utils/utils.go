package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HttpError(msg string) error {
	return echo.NewHTTPError(http.StatusBadRequest, msg)
}
