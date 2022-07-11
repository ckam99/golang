package main

import (

	// "github.com/dgrijalva/jwt-go/v4"

	"github.com/gofiber/fiber/v2"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var body LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if body.Email != "test@example.com" || body.Password != "123456" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad credentials",
		})
	}

	user := User{
		Name:  "Test",
		Email: body.Email,
		Role:  Role{Name: "user"},
	}
	token, err := CreateAccessToken(&user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	user.Token = Token{
		AccessToken: token,
	}
	return ctx.Status(fiber.StatusOK).JSON(&user)
}

func CurrentUserHandler(ctx *fiber.Ctx) error {
	claims, err := ExtractJWT(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(claims)
}

func CurrentUserHandler2(ctx *fiber.Ctx) error {
	claims, err := ExtractJsonWebToken(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(claims)
}

func WelcomeHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"token": "Welcome",
	})
}

func TestTokenHandler(ctx *fiber.Ctx) error {
	token, err := GenerateExampleToken()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"token": token,
	})
}

func ProtectedRouteHandler(c *fiber.Ctx) error {

	return c.JSON(fiber.Map{
		"message": "this route is protected",
	})
}
