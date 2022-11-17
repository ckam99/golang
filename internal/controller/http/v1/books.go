package v1

import (
	"context"
	httpUtils "main/internal/controller/http/utils"
	"main/internal/domain/books"
	"main/pkg/clients/postgresql"
	"main/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
)

type BookController struct {
	service books.Service
}

func NewBookController(db postgresql.Client) *BookController {
	return &BookController{
		service: books.NewService(db),
	}
}
func (r *BookController) HealthBook(c *fiber.Ctx) error {
	return c.JSON(httpUtils.HTTPMessage("ok"))
}

func (r *BookController) GetBooks(c *fiber.Ctx) error {
	books, err := r.service.GetAll(c.UserContext(), &books.QueryFilterDTO{})
	if err != nil {
		return httpUtils.HTTPError(c, err.Error(), 500)
	}
	return c.JSON(books)
}

func (r *BookController) GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return httpUtils.HTTPError(c, err.Error(), 400)
	}
	book, err := r.service.Find(c.UserContext(), int64(id))
	if err != nil {
		if err == utils.ErrNoEntity {
			return httpUtils.HTTPError(c, "book does not exist", 404)
		}
		return httpUtils.HTTPError(c, err.Error(), 500)
	}
	return c.JSON(book)
}

func (r *BookController) PostBook(c *fiber.Ctx) error {
	var payload books.CreateDTO
	if err := c.BodyParser(&payload); err != nil {
		return httpUtils.HTTPError(c, err.Error(), 422)
	}
	if err := validate.Validate(&payload); err != nil {
		return httpUtils.HTTPError(c, err.Error(), 422)
	}
	book, err := r.service.Create(context.Background(), payload)
	if err != nil {
		if err == utils.ErrInvalidForeinKey {
			return httpUtils.HTTPError(c, "author_id does not exists", 400)
		}
		return httpUtils.HTTPError(c, err.Error(), 500)
	}
	return c.Status(201).JSON(book)
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
