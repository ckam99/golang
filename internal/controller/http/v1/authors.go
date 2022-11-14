package v1

import (
	"main/internal/domain/authors"
	"main/pkg/clients/postgresql"

	"github.com/gofiber/fiber/v2"
)

type AuthorController struct {
	service authors.Service
}

func NewAuthorController(db postgresql.Client) *AuthorController {
	r := &AuthorController{
		service: authors.NewService(db),
	}
	return r
}

func (r *AuthorController) HealthAuthor(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "ok"})
}

func (r *AuthorController) GetAuthors(c *fiber.Ctx) error {
	return c.SendString("list authors")
}

func (r *AuthorController) GetAuthor(c *fiber.Ctx) error {
	return c.SendString("get author by id")
}

func (r *AuthorController) PostAuthor(c *fiber.Ctx) error {
	return c.SendString("create author")
}

func (r *AuthorController) PutAuthor(c *fiber.Ctx) error {
	return c.SendString("update author")
}
