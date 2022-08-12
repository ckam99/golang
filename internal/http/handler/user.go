package handler

import (
	"strconv"

	"github.com/ckam225/golang/fiber/internal/http/request"
	"github.com/ckam225/golang/fiber/internal/http/response"
	"github.com/ckam225/golang/fiber/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary     User Health Check
// @Description  User Health Check
// @Security  ApiKeyAuth
// @Tags         users
// @Accept        */*
// @Produce      json
// @Success      200  {object}   response.ErrorResponse
// @Failure      400  {object}   response.ErrorResponse
// @Failure      404  {object}   response.ErrorResponse
// @Failure      500  {object}   response.ErrorResponse
// @Router       /users/health [get]
func (h *Handler) UserHealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"message": "ok",
	})
}

// @Summary     Get Users
// @Security  ApiKeyAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "Limit pagination"
// @Param        skip     query     int  false  "Page pagination"
// @Success      200  {array}   response.UserResponse
// @Failure      404,400,401  {object}   response.ErrorResponse
// @Router       /users [get]
func (h *Handler) GetUsersHandler(c *fiber.Ctx) error {
	queryParam := request.UserFilterParam{
		Skip:  0,
		Limit: 100,
	}
	if skip, _ := strconv.Atoi(c.Query("skip", "0")); skip != 0 {
		queryParam.Skip = skip
	}
	if limit, _ := strconv.Atoi(c.Query("limit", "100")); limit != 0 {
		queryParam.Limit = limit
	}
	users, err := h.service.GetAllUsers(queryParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.SetHttpError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(response.ParseUserListEntity(users))
}

// @Summary     Create User
// @Security  ApiKeyAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param input  body request.CreateUser true "User ID"
// @Success      201  {object}  response.UserResponse
// @Failure      404,422,400  {object}   response.ErrorResponse
// @Router       /users [post]
func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	body := request.CreateUser{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(response.SetHttpError(err.Error()))
	}
	if errors := utils.ValidateCredentials(body); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	if user, err := h.service.CreateUser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.SetHttpError(err.Error()))
	} else {
		return c.Status(fiber.StatusCreated).JSON(response.ParseUserEntity(user))
	}
}

// @Summary     Get user by id
// @Security  ApiKeyAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param user_id   path int true "User ID"
// @Success      200  {object}   response.UserResponse
// @Failure      404  {object}   response.ErrorResponse
// @Router       /users/{user_id} [get]
func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.SetHttpError(err.Error()))
	} else {
		user, err := h.service.GetUserByID(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(response.SetHttpError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(response.ParseUserEntity(user))
	}
}

// @Summary     Update user
// @Security  ApiKeyAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param user_id   path int true "User ID"
// @Param input  body request.UpdateUser true "User ID"
// @Success      200  {object}   response.UserResponse
// @Failure      404  {object}   response.ErrorResponse
// @Router       /users/{user_id} [put]
func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	body := request.UpdateUser{}
	id, _ := c.ParamsInt("id", 0)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}
	user, err := h.service.UpdateUser(id, &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.SetHttpError(err.Error()))
	}
	return c.Status(fiber.StatusAccepted).JSON(&user)
}

// @Summary     Delete user
// @Security  ApiKeyAuth
// @Tags         users
// @Accept       json
// @Produce      json
// @Param user_id   path int true "User ID"
// @Success      204
// @Failure      404  {object}   response.ErrorResponse
// @Router       /users/{user_id} [delete]
func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := h.service.DeleteUser(uint(id), true); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
