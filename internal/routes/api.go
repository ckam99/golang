package routes

import (
	"os"

	"github.com/ckam225/golang/fiber/internal/http/controller"
	"github.com/ckam225/golang/fiber/internal/http/middleware"
	"github.com/ckam225/golang/fiber/internal/repository"
	"github.com/ckam225/golang/fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func SetupAPIRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")
	if os.Getenv("APP_ENV") != "production" {
		SwaggerRoutes(app)
	}
	AuthControllerRoutes(api, db)
	UserControllerRoutes(api, db)
}

func UserControllerRoutes(app fiber.Router, db *gorm.DB) {
	userCtr := &controller.UserController{Repo: repository.UserRepository{Query: db}}
	userRoute := app.Group("/users")
	userRoute.Get("/", middleware.BearerAuthMiddleware(), userCtr.GetUsersHandler)
	userRoute.Post("/", middleware.BearerAuthMiddleware(), userCtr.CreateUserHandler)
	userRoute.Get("/:id", middleware.BearerAuthMiddleware(), userCtr.GetUserHandler)
	userRoute.Put("/:id", middleware.BearerAuthMiddleware(), userCtr.UpdateUserHandler)
	userRoute.Delete("/:id", middleware.BearerAuthMiddleware(), userCtr.DeleteUserHandler)
}

func AuthControllerRoutes(app fiber.Router, db *gorm.DB) {
	authCtr := &controller.AuthController{Repo: repository.AuthRepository{Query: db}}
	authRoute := app.Group("/auth")
	authRoute.Post("/signup", authCtr.SignUpHandler)
	authRoute.Post("/token", authCtr.TokenAuthenticationHandler)
	authRoute.Get("/token/refresh", middleware.BearerAuthMiddleware(), authCtr.RefreshTokenHandler)
	authRoute.Post("/email/confirm", authCtr.EmailConfirmationHandler)
	authRoute.Post("/phone/confirm", authCtr.EmailConfirmationHandler)
	authRoute.Get("/user", middleware.BearerAuthMiddleware(), authCtr.CurrentUserHandler)
}

func SwaggerRoutes(app *fiber.App) {
	utils.SetSwaggerInfos()
	app.Get("/docs/*", swagger.HandlerDefault)
}
