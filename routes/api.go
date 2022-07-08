package routes

import (
	"example/fiber/http/controller"
	"example/fiber/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupAPIRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	// user controller
	userController := &controller.UserController{Repo: repository.UserRepository{Query: db}}
	userRoute := api.Group("/users")
	userRoute.Get("/", userController.GetUsersHandler)
	userRoute.Post("/", userController.CreateUserHandler)
	userRoute.Get("/:id", userController.GetUserHandler)
	userRoute.Put("/:id", userController.UpdateUserHandler)
	userRoute.Delete("/:id", userController.DeleteUserHandler)
}
