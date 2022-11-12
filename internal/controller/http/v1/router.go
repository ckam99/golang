package v1

import (
  "github.com/gofiber/fiber/v2"
  "database/sql"
  "main/internal/domain/books"
)
type Router struct {
  *fiber.App
}

func Register(app *fiber.App, db *sql.DB){
  router := &Router{
    App: app,
  }
  r := router.Group("/v1")
  // books endpoints 
  RegisterBookRouter(r, db)
  // authors endpoints 
  RegisterAuthorRouter(r, db)
}
