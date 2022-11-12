package v1

import (
	"database/sql"
	"main/internal/domain/books"

	"github.com/gofiber/fiber/v2"
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
	books, err := r.service.GetAll(c.UserContext(), &books.QueryFilterDTO{})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(books)
}

func (r *BookController) GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	book, err := r.service.Find(c.UserContext(), int64(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	if book.ID == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "book does not exist"})
	}
	return c.JSON(book)
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
