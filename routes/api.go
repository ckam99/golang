package routes

import (
	"example/fiber/http/controller"
	"example/fiber/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIRoute struct {
	DB  *gorm.DB
	App *fiber.App
}

func (w *APIRoute) SetupAPIRoutes() {
	api := w.App.Group("/api")

	userController := &controller.UserController{Repo: repository.UserRepository{Query: w.DB}}
	userRoute := api.Group("/users")
	userRoute.Get("/", userController.GetUsersHandler)
	userRoute.Post("/", userController.CreateUserHandler)
	userRoute.Get("/:id", userController.GetUserHandler)
	userRoute.Put("/:id", userController.UpdateUserHandler)
	userRoute.Delete("/:id", userController.DeleteUserHandler)
}
