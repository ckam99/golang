package handler

import (
	"app/pkg/clients/postgresql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Start(host, port string)
	Register()
}

func NewHandler(client postgresql.Client) Handler {
	h := &handle{
		App:    *fiber.New(),
		client: client,
	}
	h.Register()
	return h
}

type handle struct {
	client postgresql.Client
	fiber.App
}

// Register implements Handler
func (h *handle) Register() {
	RegisterAuthorHandler("/authors", h.client, h.App)
}

// Start implements Handler
func (h *handle) Start(host, port string) {
	log.Fatal(h.Listen(host + ":" + port))
}
