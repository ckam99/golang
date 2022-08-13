package handler

import (
	"errors"
	"log"
	"time"

	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/http/response"
	"github.com/ckam225/golang/fiber/internal/security"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Sign Up
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input  body request.RegisterRequest true "Credential"
// @Success      201  {object}  response.UserResponse
// @Failure      400,422,500  {object}   response.ErrorResponse
// @Router       /auth/signup [post]
func (h *Handler) SignUpHandler(ctx *fiber.Ctx) error {
	var body request.RegisterRequest
	if err := ctx.BodyParser(&body); err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	if err := h.Validate(body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	user, err := h.service.SignUp(body)
	if err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	go h.service.SendConfirmationEmail(user)
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
func (h *Handler) CurrentUserHandler(ctx *fiber.Ctx) error {
	user, err := security.GetAuthUser(h.service.GetDB(), ctx)

	if err != nil {
		return h.Raise(ctx, fiber.StatusUnauthorized, err)
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
func (h *Handler) EmailConfirmationHandler(ctx *fiber.Ctx) error {
	var req request.EmailConfirmRequest
	if err := ctx.BodyParser(&req); err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	if err := h.Validate(req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	verifyCode, err := h.service.VerifyConfirmationCode(&req)
	if err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	user, err := h.service.GetUserByEmail(req.Email)
	if err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	if err = h.service.GetDB().Model(&user).UpdateColumn("email_confirmed_at", time.Now()).Error; err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	go func() {
		if err = h.service.GetDB().Unscoped().Delete(&verifyCode).Error; err != nil {
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
func (h *Handler) TokenAuthenticationHandler(ctx *fiber.Ctx) error {
	body := request.LoginRequest{}

	if err := ctx.BodyParser(&body); err != nil {
		return h.Raise(ctx, fiber.StatusBadRequest, err)
	}
	if err := h.Validate(&body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}
	user, err := h.service.SignIn(&body)
	if err != nil {
		return h.Raise(ctx, fiber.StatusUnauthorized, errors.New("bad credentials"))
	}
	token, err := security.CreateAccessToken(user)
	if err != nil {
		return h.Raise(ctx, fiber.StatusUnauthorized, err)
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
func (h *Handler) RefreshTokenHandler(ctx *fiber.Ctx) error {
	user, err := security.GetAuthUser(h.service.GetDB(), ctx)
	if err != nil {
		return h.Raise(ctx, fiber.StatusUnauthorized, err)
	}
	currentToken, err := security.DecodeJWT(ctx.Get("Authorization"))
	if err != nil {
		return h.Raise(ctx, fiber.StatusUnauthorized, err)
	}
	newToken, err := security.RefreshAccessToken(user, currentToken)
	return ctx.JSON(response.Token{
		ID:           user.ID,
		Email:        user.Email,
		AccessToken:  currentToken,
		RefreshToken: newToken,
	})
}
