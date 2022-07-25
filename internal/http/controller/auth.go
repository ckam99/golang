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

// @Summary     Sign In
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.LoginRequest true "Credential"
// @Success      200  {object}  response.AccessToken
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
	claims, err := security.ExtractJWT(ctx)
	if err != nil {
		return response.HttpResponseError(ctx, fiber.StatusBadRequest, err.Error())
	}
	user := security.GetUserFromClaim(*claims)
	if err = c.Repo.Query.Find(&user).Error; err != nil {
		return response.HttpResponseError(ctx, fiber.StatusUnauthorized, "Bad credentials")
	}
	return ctx.JSON(response.ParseUserEntity(user))
}

// @Summary     Email|Phone confirmation
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.EmailConfirmRequest true "Credential"
// @Success      204
// @Failure      400,401,500  {object}   response.ErrorResponse
// @Router       /auth/confirm/email [post]
// @Router       /auth/confirm/phone [post]
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
