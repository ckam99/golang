package http

import (
  "github.com/gofiber/fiber/v2"
  "main/internal/controller/v1"
  "database/sql"
)
type HttpServer struct {
  *fiber.App
}

func NewHTTP(db *sql.DB) *HttpServer{
  s := &HttpServer{
    App: fiber.New(),
    db: db,
  }
  v1.RegisterRouter(s.App,db)
  return s
}

func(s *HttpServer) Run() error {
  return s.Listen(":8000")
}
