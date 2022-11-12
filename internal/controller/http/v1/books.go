package v1

import (
  "github.com/gofiber/fiber/v2"
  "main/internal/domain/books"
  "database/sql"
)

type BookController struct {
  service books.Service,
  *fiber.App
}

func RegisterBookRouter(app *fiber.App,db *sql.DB) *BookController{
  
  r := &bookRouter{
    App: app,
    service: books.NewService(db)
  }

  r.Get("/books", r.GetBooks)
  r.Get("/books/:id", r.GetBook)
  r.Post("/books", r.PostBook)
  r.Put("/books/:id", r.PutBook)
  r.Patch("/books/:id", r.PatchBook)
  r.Delete("/books/:id", r.DeleteBook)
  return r
}

func(r *BookController) GetBooks(c *fiber.Ctx) error{
   return c.String("list books")
}

func(r *BookController) GetBook(c *fiber.Ctx) error{
  books, err:= r.service.GetAll(&books.QueryFilterDTO{})
  if err != nil{
    return c.Json(fiber.Map{
      "error": err.Error()
  }).Status(fiber.StatusInternalError)
  }
   return c.Json(books)
}

func(r *BookController) PostBook(c *fiber.Ctx) error{
   return c.String("create book")
}

func(r *BookController) PutBook(c *fiber.Ctx) error{
   return c.String("update book by id")
}

func(r *BookController) PatchBook(c *fiber.Ctx) error{
   return c.String("partial update book")
}

func(r *BookController) DeleteBook(c *fiber.Ctx) error{
   return c.String("delete by id")
}
