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

// @Summary     Sign In
// @Security  ApiKeyAuth
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.LoginRequest true "Credential"
// @Success      201  {object}  response.AccessToken
// @Failure      404,422,400  {object}   response.ErrorResponse
// @Router       /auth/signin [post]
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

// @Summary     Sign Up
// @Security  ApiKeyAuth
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.RegisterRequest true "Credential"
// @Success      201  {object}  response.UserResponse
// @Failure      404,422,400  {object}   response.ErrorResponse
// @Router       /auth/signup [post]
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
