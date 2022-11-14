package v1

import (
	"main/internal/domain/auth"
	"main/pkg/clients/postgresql"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service auth.Service
}

func NewAuthController(db postgresql.Client) *AuthController {
	r := &AuthController{
		service: auth.NewService(db),
	}
	return r
}

func (r *AuthController) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "ok"})
}

func (r *AuthController) GetCurrentUser(c *fiber.Ctx) error {
  user, err := c.service.GetAuthUser(ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
      "message": err.Error()
    })
	}
  return ctx.JSON(user)
}

func (r *AuthController) SignIn(c *fiber.Ctx) error {
	return c.SendString("get author by id")
}

func (r *AuthController) SignUp(c *fiber.Ctx) error {
	return c.SendString("create author")
}

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
  bearToken := ctx.Get("Authorization ")
	user, err := c.service.GetAuthUser(bearToken)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
      "message": err.Error()
    })
	}
 token, err := c.service.RefreshAccessToken(&user,bearToken)
  if err != nil{
    return ctx.Status(401).JSON(fiber.Map{
      "message": err.Error(),
    })
  }
	return ctx.JSON(token)
}
