package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Welcome godoc
// @Summary     Welcome
// @Description  Welcome
// @Tags         healths
// @Accept        */*
// @Produce      json
// @Success      200  {object}   models.Response
// @Failure      400  {object}   models.Response
// @Failure      404  {object}   models.Response
// @Failure      500  {object}   models.Response
// @Router       / [get]
func HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "welcome",
	})
}
