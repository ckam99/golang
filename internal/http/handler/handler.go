package handler

import (
	"os"

	"github.com/ckam225/golang/fiber/internal/config"
	"github.com/ckam225/golang/fiber/internal/http/middleware"
	"github.com/ckam225/golang/fiber/internal/repository"
	"github.com/ckam225/golang/fiber/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

type Handler struct {
	*fiber.App
	repo *repository.Repository
}

func NewHandler(cfg *config.Configuration) *Handler {

	ctrl := &Handler{
		App: fiber.New(fiber.Config{
			Views: cfg.Server.HtmlEngine,
			// ViewsLayout: "layouts/base",
		}),
		repo: repository.NewRepositoy(*cfg.Database),
	}
	ctrl.setupWebRoutes()
	ctrl.setupAPIRoutes()

	// middleware
	ctrl.Use(middleware.TestMiddleware)
	ctrl.Use(middleware.CorsMiddleware())
	//ctrl.Use(middleware.RouteMiddleware)

	return ctrl
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
