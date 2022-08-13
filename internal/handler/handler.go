package handler

import (
	"log"
	"strings"

	"github.com/ckam225/golang/fiber-sqlx/internal/database/postgres/storage"
	"github.com/ckam225/golang/fiber-sqlx/internal/dto"
	"github.com/ckam225/golang/fiber-sqlx/internal/service"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	*fiber.App
	service   *service.Service
	validator *validator.Validate
}

func NewHandler(store storage.Store) *Handler {
	h := &Handler{
		App:       fiber.New(),
		service:   service.NewService(&store),
		validator: validator.New(),
	}
	//	h.Use(middleware.Logger())
	//	h.Use(middleware.Recover())

	h.setupApiRoutes()
	h.setupWebRoutes()

	return h
}

func (h *Handler) setupApiRoutes() {
	users := h.Group("/users")
	users.Get("/", h.GetUsersHandler)
	users.Post("/", h.CreateUserHandler)
	users.Get("/:id", h.GetUserHandler)
	users.Delete("/:id", h.DeleteUserHandler)
}

func (h *Handler) setupWebRoutes() {
	h.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(dto.HttpError{
			Message: "Welcome echo",
		})
	})
}

func (h *Handler) Raise(c *fiber.Ctx, statusCode int, err error) error {
	log.Println(err)
	return c.Status(statusCode).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func (h *Handler) Validate(s interface{}) []*dto.ValidateError {
	var errors []*dto.ValidateError
	err := h.validator.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element dto.ValidateError
			element.Namespace = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Field = strings.ToLower(err.Field())

			errors = append(errors, &element)
		}
	}
	return errors
}
