package handler

import (
	"app/internal/domain/author"
	"app/pkg/clients/postgresql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type AuthorHandler struct {
	service author.Service
}

func RegisterAuthorHandler(prefix string, client postgresql.Client, app fiber.App) *AuthorHandler {
	h := &AuthorHandler{
		service: author.NewService(client),
	}
	router := app.Group(prefix)
	router.Get("/", h.GetAllAuthorsHandler)
	router.Get("/{id}", h.GetAuthorHandler)
	return h
}

func (ahandler *AuthorHandler) GetAllAuthorsHandler(ctx *fiber.Ctx) error {
	var queryParams author.FilterParamsDTO
	if err := ctx.ParamsParser(&queryParams); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	authors, err := ahandler.service.GetAuthors(ctx.Context(), queryParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(authors)
}

func (ahandler *AuthorHandler) GetAuthorHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	author, err := ahandler.service.FindAuthor(ctx.Context(), id)
	if err != nil {
		log.Println(err.Error())
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(author)
}
