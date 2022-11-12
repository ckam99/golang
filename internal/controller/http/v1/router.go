package v1

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	db *sql.DB
}

func NewRouter(db *sql.DB) *Router {
	return &Router{
		db: db,
	}
}

func (r *Router) Routes(router fiber.Router) {

	// books endpoints
	book := NewBookController(r.db)
	router.Get("/books/health", book.HealthBook)
	router.Get("/books", book.GetBooks)
	router.Get("/books/:id", book.GetBook)
	router.Post("/books", book.PostBook)
	router.Put("/books/:id", book.PutBook)
	router.Patch("/books/:id", book.PatchBook)
	router.Delete("/books/:id", book.DeleteBook)

	// authors endpoints
	author := NewAuthorController(r.db)
	router.Get("/authors/health", author.HealthAuthor)
	router.Get("/authors", author.GetAuthors)
	router.Get("/authors/:id", author.GetAuthor)
	router.Post("/authors", author.PostAuthor)
	router.Put("/authors/:id", author.PutAuthor)

}
