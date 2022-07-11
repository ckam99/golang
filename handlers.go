package main

import (

	// "github.com/dgrijalva/jwt-go/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func LoginHandler(ctx *fiber.Ctx) error {
	var body LoginRequest
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	if body.Email != "test@example" && body.Password != "123456" {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": "Bad credentials",
		})
	}
	token, err := GenerateExampleToken()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}
	user := User{
		Name:  "Test",
		Email: body.Email,
		Token: Token{
			AccessToken: token,
		},
	}
	return ctx.Status(fiber.StatusOK).JSON(&user)
}

func CurrentUserHandler(ctx *fiber.Ctx) error {
	token, err := VerifyJWT(ctx.Get("Authorization"))
	if err != nil {
		return ctx.JSON(err)
	}
	println(token)

	user := ctx.Locals("x-fiber-user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	//now := time.Now().Unix()

	// // Get claims from JWT.
	// claims, err := ExtractJWTMetadata(ctx)
	// if err != nil {
	// 	// Return status 500 and JWT parse error.
	// 	return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }
	// // Set expiration time from JWT data of current book.
	// expires := claims.Expires

	// // Checking, if now time greather than expiration from JWT.
	// if now > expires {
	// 	// Return status 401 and unauthorized error message.
	// 	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	return ctx.JSON(claims)
}

func CurrentUserHandler2(ctx *fiber.Ctx) error {
	user, err := ExtractTokenUser(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return ctx.JSON(user)
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
