package handler

import (
	"log"
	"net/http"

	"github.com/ckam225/golang/echo/internal/database/postgres/storage"
	"github.com/ckam225/golang/echo/internal/dto"
	"github.com/ckam225/golang/echo/internal/service"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handler struct {
	*echo.Echo
	service *service.Service
}

func NewHandler(store storage.Store) *Handler {
	h := &Handler{
		Echo:    echo.New(),
		service: service.NewService(&store),
	}
	h.Validator = &dto.CustomValidator{Validator: validator.New()}
	h.Use(middleware.Logger())
	h.Use(middleware.Recover())

	h.setupApiRoutes()
	h.setupWebRoutes()

	return h
}

func (h *Handler) setupApiRoutes() {

	users := h.Group("/users")
	users.GET("/", h.GetUsersHandler)
	users.POST("/", h.CreateUserHandler)
	users.GET("/:id", h.GetUserHandler)
	users.DELETE("/:id", h.DeleteUserHandler)
}

func (h *Handler) setupWebRoutes() {
	h.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dto.ErrorResponse{
			Message: "Welcome echo",
		})
	})
}

func (h *Handler) fatal(statusCode int, err error) error {
	log.Println(err)
	return echo.NewHTTPError(statusCode, &dto.ErrorResponse{
		Message: err.Error(),
	})
}
