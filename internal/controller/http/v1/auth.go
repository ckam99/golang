package v1

import (
	"context"
	"main/internal/domain/auth"
	"main/pkg/clients/postgresql"
	"main/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gopkg.in/dealancer/validate.v2"
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
	var payload auth.LoginDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Validate(&payload); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).
			JSON(fiber.Map{"message": err.Error()})
	}
	token, err := c.service.Login(context.Background(), payload)
	if err != nil {
		if err == utils.ErrNoEntity || err == utils.ErrInvalidCredentials {
			return ctx.Status(404).JSON(fiber.Map{
				"message": "email/phone or password is invalid",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(token)
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {
	var payload auth.RegisterDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"message": err.Error()})
	}
	if err := validate.Validate(&payload); err != nil {
		return ctx.Status(422).JSON(fiber.Map{"message": err.Error()})
	}
	user, err := c.service.Register(context.Background(), payload)
	if err != nil {
		if err == utils.ErrUniqueField {
			return ctx.Status(404).JSON(fiber.Map{
				"message": "email or phone is not available",
			})
		}
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(user)
}
