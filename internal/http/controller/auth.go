package controller

import (
	"log"
	"time"

	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/http/response"
	"github.com/ckam225/golang/fiber/internal/repository"
	"github.com/ckam225/golang/fiber/internal/security"
	"github.com/ckam225/golang/fiber/internal/service"
	"github.com/ckam225/golang/fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	Repo repository.AuthRepository
}

// @Summary     Sign Up
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.RegisterRequest true "Credential"
// @Success      201  {object}  response.UserResponse
// @Failure      400,422,500  {object}   response.ErrorResponse
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
	go service.SendConfirmationEmail(c.Repo.Query, user)
	return ctx.Status(fiber.StatusCreated).JSON(response.ParseUserEntity(user))
}

// @Summary     Current user
// @Security  ApiKeyAuth
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.UserResponse
// @Failure      400,401,500  {object}   response.ErrorResponse
// @Router       /auth/user [get]
func (c *AuthController) CurrentUserHandler(ctx *fiber.Ctx) error {
	user, err := security.GetAuthUser(c.Repo.Query, ctx)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusUnauthorized, err.Error())
	}
	return ctx.JSON(response.ParseUserEntity(user))
}

// @Summary     Email/Phone confirmation
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.EmailConfirmRequest true "Credential"
// @Success      204
// @Failure      400,401,500  {object}   response.ErrorResponse
// @Router       /auth/email/confirm [post]
// @Router       /auth/phone/confirm [post]
func (c *AuthController) EmailConfirmationHandler(ctx *fiber.Ctx) error {
	var req request.EmailConfirmRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	if errors := utils.ValidateCredentials(req); errors != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	verifyCode, err := service.VerifyConfirmationCode(c.Repo.Query, &req)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	user, err := service.GetUserByEmail(c.Repo.Query, req.Email)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	if err = c.Repo.Query.Model(&user).UpdateColumn("email_confirmed_at", time.Now()).Error; err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	go func() {
		if err = c.Repo.Query.Unscoped().Delete(&verifyCode).Error; err != nil {
			log.Println(err.Error())
		}
	}()

	return ctx.Status(fiber.StatusNoContent).JSON(verifyCode)
}

// @Summary     Token Authentication
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.LoginRequest true "Credential"
// @Success      200  {object}  response.Token
// @Failure      404,422,400  {object}   response.ErrorResponse
// @Router       /auth/token [post]
func (c *AuthController) TokenAuthenticationHandler(ctx *fiber.Ctx) error {
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
	return ctx.JSON(response.Token{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  token,
		RefreshToken: token,
	})
}

// @Summary     Refresh Token
// @Security  ApiKeyAuth
// @Tags         auth
// @Accept       json
// @Produce      json
// @Success      200  {object}  response.Token
// @Failure      400,401,500  {object}   response.ErrorResponse
// @Router       /auth/token/refresh [get]
func (c *AuthController) RefreshTokenHandler(ctx *fiber.Ctx) error {
	user, err := security.GetAuthUser(c.Repo.Query, ctx)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusUnauthorized, err.Error())
	}
	currentToken, err := security.DecodeJWT(ctx.Get("Authorization"))
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusUnauthorized, err.Error())
	}
	newToken, err := security.RefreshAccessToken(user, currentToken)
	return ctx.JSON(response.Token{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  currentToken,
		RefreshToken: newToken,
	})
}
