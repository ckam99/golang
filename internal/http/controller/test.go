package controller

import (
	"github.com/ckam225/golang/fiber/internal/entity"
	"github.com/ckam225/golang/fiber/pkg/mailer"

	"github.com/gofiber/fiber/v2"
)

type TestController struct {
}

func (C *TestController) TestMailHandler(ctx *fiber.Ctx) error {
	var user entity.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.JSON(fiber.Map{
			"message": "Email sent successfully",
		})
	}
	mailer.NotificationChannel <- &mailer.Notification{
		To: []string{user.Email},
		Data: map[string]string{
			"name":  user.Email,
			"email": user.Name,
		},
		Subject:  "Tesy notify",
		Template: "mail/register.tmpl",
	}
	return ctx.JSON(fiber.Map{
		"message": "Email sent successfully",
	})
}

func (C *TestController) TestPushNotificationHandler(ctx *fiber.Ctx) error {
	var users []entity.User

	if err := ctx.BodyParser(&users); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var to []string
	for _, user := range users {
		to = append(to, user.Email)
	}
	mailer.NotificationChannel <- &mailer.Notification{
		To: to,
		Data: map[string]string{
			"name":  "Welcome",
			"email": "how is your life",
		},
		Subject:  "Tesy notify",
		Template: "mail/register.tmpl",
	}
	return ctx.JSON(fiber.Map{
		"message": "Email sent successfully",
	})
}
