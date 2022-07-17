package main

import (

	// "github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const hash = "$2a$10$0ngjo9r3Sl3PfUGKSqbxYO4nTDLixreZf83O0gTWJkxthTiWI1aLa" // hash of 123456

func LoginHandler(ctx *fiber.Ctx) error {
	var body LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if body.Email != "test@example.com" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad credentials: email does not match",
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(body.Password)); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad credentials: password does not match",
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
