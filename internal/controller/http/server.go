package http

import (
	"main/internal/config"
	routerV1 "main/internal/controller/http/v1"
	"main/pkg/clients/postgresql"

	"github.com/gofiber/fiber/v2"
)

type HttpServer struct {
	*fiber.App
	cfg config.Config
}

func NewHTTP(db postgresql.Client, cfg config.Config) *HttpServer {
	s := &HttpServer{
		App: fiber.New(),
		cfg: cfg,
	}
	s.Route("/v1", routerV1.NewRouter(db).Routes)
	return s
}

func (s *HttpServer) Run() error {
	return s.Listen(":8000")
}
