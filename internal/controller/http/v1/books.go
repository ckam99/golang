package v1

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"main/internal/domain/books"
)

type BookController struct {
	service books.Service
}

func NewBookController(db *sql.DB) *BookController {
	return &BookController{
		service: books.NewService(db),
	}
}
func (r *BookController) HealthBook(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "ok"})
}

func (r *BookController) GetBooks(c *fiber.Ctx) error {
	return c.SendString("list books")
}

func (r *BookController) GetBook(c *fiber.Ctx) error {
	books, err := r.service.GetAll(c.UserContext(), &books.QueryFilterDTO{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(books)
}

func (r *BookController) PostBook(c *fiber.Ctx) error {
	return c.SendString("create book")
}

func (r *BookController) PutBook(c *fiber.Ctx) error {
	return c.SendString("update book by id")
}

func (r *BookController) PatchBook(c *fiber.Ctx) error {
	return c.SendString("partial update book")
}

func (r *BookController) DeleteBook(c *fiber.Ctx) error {
	return c.SendString("delete by id")
}
