package handler

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"id-provider/oauth"
)

type Router struct {
	oauth *oauth.Oauth
	*fiber.App
}

func NewRouter(srv *oauth.Oauth, app *fiber.App) *Router {
	return &Router{
		srv,
		app,
	}
}

func (r *Router) Register() {
	r.Get("/", r.GreetHandler)
	r.Get("/protected", adaptor.HTTPMiddleware(r.oauth.Middleware), r.ProtectedHandler)
	r.Get("/token", adaptor.HTTPHandlerFunc(r.oauth.GetAccessToken))
	r.Post("credentials", r.GenerateCredentialsHandler)
}

func (r *Router) GenerateCredentialsHandler(c *fiber.Ctx) error {
	credential, err := r.oauth.GetCredentials()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(credential)
}

func (r *Router) ProtectedHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "this is a protected",
	})
}

func (r *Router) GreetHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "hello",
	})
}
