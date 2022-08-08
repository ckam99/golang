package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ckam225/golang/echo/internal/dto"
	"github.com/ckam225/golang/echo/internal/entity"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetUsersHandler(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 1
	}
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	users, err := h.service.GetUsers(limit, offset)

	if err != nil {
		return h.fatal(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, &users)
}

func (h *Handler) GetUserHandler(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.fatal(http.StatusBadRequest, err)
	}
	user, err := h.service.GetUser(userId)
	if err != nil {
		return h.fatal(http.StatusInternalServerError, err)
	}
	if user.ID == uuid.Nil {
		return h.fatal(http.StatusNotFound, errors.New("user not found"))
	}
	return c.JSON(http.StatusOK, &user)
}

func (h *Handler) CreateUserHandler(c echo.Context) error {
	userDTO := dto.CreateUser{}
	if err := c.Bind(&userDTO); err != nil {
		return h.fatal(http.StatusBadRequest, err)
	}
	if err := c.Validate(userDTO); err != nil {
		return h.fatal(http.StatusUnprocessableEntity, err)
	}
	if err := h.service.CreateUser(&entity.User{
		ID:    uuid.New(),
		Name:  userDTO.Name,
		Email: userDTO.Email,
	}); err != nil {
		return h.fatal(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *Handler) DeleteUserHandler(c echo.Context) error {
	userId, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.fatal(http.StatusBadRequest, err)
	}
	if err := h.service.DeleteUser(userId); err != nil {
		return h.fatal(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusNoContent, nil)
}
