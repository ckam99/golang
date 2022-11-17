package v1

import (
	"context"
	httpUtils "main/internal/controller/http/utils"
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
	return ctx.JSON(httpUtils.HTTPMessage("ok"))
}

func (c *AuthController) CurrentUser(ctx *fiber.Ctx) error {
	user, err := c.service.GetCurrentUser(ctx.UserContext(), ctx.Get("Authorization"))
	if err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 401)
	}
	return ctx.JSON(user)
}

func (c *AuthController) RefreshToken(ctx *fiber.Ctx) error {
	bearToken := ctx.Get("Authorization")
	user, err := c.service.GetCurrentUser(ctx.UserContext(), bearToken)
	if err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 401)
	}
	token, err := c.service.RefreshAccessToken(&user, bearToken)
	if err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 401)
	}
	return ctx.JSON(token)
}

func (c *AuthController) SignIn(ctx *fiber.Ctx) error {
	var payload auth.LoginDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 422)
	}
	if err := validate.Validate(&payload); err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 422)
	}
	token, err := c.service.Login(context.Background(), payload)
	if err != nil {
		if err == utils.ErrNoEntity || err == utils.ErrInvalidCredentials {
			return httpUtils.HTTPError(ctx, "email/phone or password is invalid", 404)
		}
		return httpUtils.HTTPError(ctx, err.Error(), 500)
	}
	return ctx.JSON(token)
}

func (c *AuthController) SignUp(ctx *fiber.Ctx) error {
	var payload auth.RegisterDTO
	if err := ctx.BodyParser(&payload); err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 400)
	}
	if err := validate.Validate(&payload); err != nil {
		return httpUtils.HTTPError(ctx, err.Error(), 422)
	}
	user, err := c.service.Register(context.Background(), payload)
	if err != nil {
		if err == utils.ErrUniqueField {
			return httpUtils.HTTPError(ctx, "email or phone is not available", 404)
		}
		return httpUtils.HTTPError(ctx, err.Error(), 500)
	}
	return ctx.Status(201).JSON(user)
}
