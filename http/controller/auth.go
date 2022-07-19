package controller

import (
	"example/fiber/http/request"
	"example/fiber/http/response"
	"example/fiber/repository"
	"example/fiber/security"
	"example/fiber/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Repo repository.AuthRepository
}

func (c *AuthController) SignInHandler(ctx *fiber.Ctx) error {
	body := request.LoginRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	if errors := utils.ValidateCredentials(&body); errors != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	user, err := c.Repo.SignIn(&body)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusUnauthorized, "Bad credentials")
	}
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusUnauthorized, err.Error())
	}
	return ctx.JSON(response.AccessToken{
		ID:          user.ID,
		Email:       user.Email,
		AccessToken: token,
	})
}

func (c *AuthController) SignUpHandler(ctx *fiber.Ctx) error {
	var body request.RegisterRequest
	if err := ctx.BodyParser(&body); err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	if errors := utils.ValidateCredentials(body); errors != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	user, err := c.Repo.SignUp(body)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	return ctx.Status(fiber.StatusCreated).JSON(user)
}
