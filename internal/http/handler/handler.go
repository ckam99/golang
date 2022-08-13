package handler

import (
	"log"
	"os"
	"strings"

	"github.com/ckam225/golang/fiber/internal/config"
	"github.com/ckam225/golang/fiber/internal/http/middleware"
	"github.com/ckam225/golang/fiber/internal/http/response"
	"github.com/ckam225/golang/fiber/internal/service"
	"github.com/ckam225/golang/fiber/internal/utils"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Handler struct {
	*fiber.App
	service   *service.Service
	validator *validator.Validate
}

func NewHandler(cfg *config.Configuration) *Handler {

	h := &Handler{
		App: fiber.New(fiber.Config{
			Views: cfg.Server.HtmlEngine,
			// ViewsLayout: "layouts/base",
		}),
		service:   service.NewService(*cfg),
		validator: validator.New(),
	}
	h.setupWebRoutes()
	h.setupAPIRoutes()

	// middleware
	h.Use(middleware.TestMiddleware)
	h.Use(middleware.CorsMiddleware())
	//h.Use(middleware.RouteMiddleware)

	return h
}

func (c *Handler) setupAPIRoutes() {
	router := c.Group("/api")

	if os.Getenv("APP_ENV") != "production" {
		utils.SetSwaggerInfos()
		c.Get("/docs/*", swagger.HandlerDefault)
	}

	authRouter := router.Group("/auth")
	authRouter.Post("/signup", c.SignUpHandler)
	authRouter.Post("/token", c.TokenAuthenticationHandler)
	authRouter.Get("/token/refresh", middleware.BearerAuthMiddleware(), c.RefreshTokenHandler)
	authRouter.Post("/email/confirm", c.EmailConfirmationHandler)
	authRouter.Post("/phone/confirm", c.EmailConfirmationHandler)
	authRouter.Get("/user", middleware.BearerAuthMiddleware(), c.CurrentUserHandler)

	userRouter := router.Group("/users")
	userRouter.Get("/", middleware.BearerAuthMiddleware(), c.GetUsersHandler)
	userRouter.Post("/", middleware.BearerAuthMiddleware(), c.CreateUserHandler)
	userRouter.Get("/:id", middleware.BearerAuthMiddleware(), c.GetUserHandler)
	userRouter.Put("/:id", middleware.BearerAuthMiddleware(), c.UpdateUserHandler)
	userRouter.Delete("/:id", middleware.BearerAuthMiddleware(), c.DeleteUserHandler)

}

func (c *Handler) setupWebRoutes() {
	c.Get("/", func(c *fiber.Ctx) error {
		return c.Render("welcome", fiber.Map{
			"Title": "Hello, <b>World</b>!",
		})
	})
	c.Get("about", func(c *fiber.Ctx) error {
		return c.Render("about", fiber.Map{})
	})
}

func (h *Handler) Raise(c *fiber.Ctx, statusCode int, err error) error {
	log.Println(err)
	return c.Status(statusCode).JSON(fiber.Map{
		"message": err.Error(),
	})
}

func (h *Handler) Validate(s interface{}) []*response.ValidationError {
	var errors []*response.ValidationError
	err := h.validator.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ValidationError
			element.Namespace = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Field = strings.ToLower(err.Field())

			errors = append(errors, &element)
		}
	}
	return errors
}
