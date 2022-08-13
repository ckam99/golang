package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ckam225/golang/fiber-sqlx/internal/dto"
	"github.com/ckam225/golang/fiber-sqlx/internal/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) GetUsersHandler(c *fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Params("limit", "1"))
	if limit == 0 {
		limit = 1
	}
	offset, _ := strconv.Atoi(c.Params("offset", "0"))

	users, err := h.service.GetUsers(limit, offset)

	if err != nil {
		return h.Raise(c, http.StatusInternalServerError, err)
	}
	return c.JSON(&users)
}

func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return h.Raise(c, http.StatusBadRequest, err)
	}
	user, err := h.service.FindUser(userId)
	if err != nil {
		return h.Raise(c, http.StatusInternalServerError, err)
	}
	if user.ID == uuid.Nil {
		return h.Raise(c, http.StatusNotFound, errors.New("user not found"))
	}
	return c.JSON(&user)
}

func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	userDTO := dto.CreateUser{}
	if err := c.BodyParser(&userDTO); err != nil {
		return h.Raise(c, http.StatusBadRequest, err)
	}
	if err := h.Validate(userDTO); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err)
	}

	if h.service.IsEmailExist(userDTO.Email) {
		return h.Raise(c, http.StatusBadRequest, fmt.Errorf("email %s not available", userDTO.Email))
	}

	if err := h.service.CreateUser(&entity.User{
		ID:    uuid.New(),
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}); err != nil {
		return h.Raise(c, http.StatusInternalServerError, err)
	}

	return c.Status(http.StatusNoContent).SendString("")
}

func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	userId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return h.Raise(c, http.StatusBadRequest, err)
	}
	if err := h.service.DeleteUser(userId); err != nil {
		return h.Raise(c, http.StatusInternalServerError, err)
	}
	return c.Status(http.StatusNoContent).SendString("")
}
