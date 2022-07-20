package routes

import (
	"example/fiber/http/controller"
	"example/fiber/repository"
	"example/fiber/utils"
	"os"

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
	userRoute.Get("/", userCtr.GetUsersHandler)
	userRoute.Post("/", userCtr.CreateUserHandler)
	userRoute.Get("/:id", userCtr.GetUserHandler)
	userRoute.Put("/:id", userCtr.UpdateUserHandler)
	userRoute.Delete("/:id", userCtr.DeleteUserHandler)
}

func AuthControllerRoutes(app fiber.Router, db *gorm.DB) {
	authCtr := &controller.AuthController{Repo: repository.AuthRepository{Query: db}}
	authRoute := app.Group("/auth")
	authRoute.Post("/signin", authCtr.SignInHandler)
	authRoute.Post("/signup", authCtr.SignUpHandler)
}

func SwaggerRoutes(app *fiber.App) {
	utils.SetSwaggerInfos()
	app.Get("/docs/*", swagger.HandlerDefault)
}
