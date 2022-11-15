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

func (c *AuthController) Health(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"message": "ok"})
}

func (c *AuthController) CurrentUser(ctx *fiber.Ctx) error {
	user, err := c.service.GetCurrentUser(ctx.UserContext(), ctx.Get("Authorization"))
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(user)
}

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	bearToken := ctx.Get("Authorization")
	user, err := c.service.GetCurrentUser(ctx.UserContext(), bearToken)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	token, err := c.service.RefreshAccessToken(&user, bearToken)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(token)
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {
	return ctx.SendString("get author by id")
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {
	return ctx.SendString("create author")
}
