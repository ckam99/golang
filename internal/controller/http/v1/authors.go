package v1

import (
  "github.com/gofiber/fiber/v2"
)

import (
  "github.com/gofiber/fiber/v2"
  "main/internal/domain/books"
  "database/sql"
)

type AuthorController struct {
  // service authors.Service,
  *fiber.App
}

func RegisterAuthorRouter(app *fiber.App,db *sql.DB) *AuthorController{
  r := &AuthorController{
    App: app,
    // service: authors.NewService(db),
  }

  r.Get("/authors", r.GetAuthors)
  r.Get("/authors/:id", r.GetAuthor)
  r.Post("/authors", r.PostAuthor)
  r.Put("/authors/:id", r.PutAuthor)
  
  return r
}

func(r *AuthorController) GetAuthers(c *fiber.Ctx) error{
   return c.String("list authors")
}

func(r *AuthorController) GetAuthor(c *fiber.Ctx) error{
   return c.String("get author by id")
}

func(r *AuthorController) PostAuthor(c *fiber.Ctx) error{
   return c.String("create author")
}

func(r *AuthorController) PutAuthor(c *fiber.Ctx) error{
   return c.String("update author")
}
