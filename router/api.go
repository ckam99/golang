package router

import (
	"example/fiber/handler"
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

	userHandler := &handler.UserHanlder{Repo: repository.UserRepository{Query: w.DB}}
	userRoute := api.Group("/users")
	userRoute.Get("/", userHandler.GetUsersHandler)
	userRoute.Post("/", userHandler.CreateUserHandler)
	userRoute.Get("/:id", userHandler.GetUserHandler)
	userRoute.Put("/:id", userHandler.UpdateUserHandler)
	userRoute.Delete("/:id", userHandler.DeleteUserHandler)
}
