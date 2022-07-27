package handlers

import (
	"net/http"
	"strconv"

	"github.com/ckam225/golang/echo/internal/database"
	"github.com/ckam225/golang/echo/internal/models"
	"github.com/ckam225/golang/echo/internal/schemas"
	"github.com/labstack/echo/v4"
)

func GetUsersHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, database.UserList)
}

func GetUserHandler(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &schemas.ErrorResponse{
			Message: err.Error(),
		})
	}
	var user models.User
	for _, v := range database.UserList {
		if v.ID == userId {
			user = v
			break
		}
	}
	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, &schemas.ErrorResponse{
			Message: "User not found",
		})
	}
	return c.JSON(http.StatusOK, &user)
}

func CreateUserHandler(c echo.Context) error {
	user := schemas.CreateUser{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(user); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}
	newUser := models.User{
		ID:    len(database.UserList) + 1,
		Name:  user.Name,
		Email: user.Email,
	}
	database.UserList = append(database.UserList, newUser)
	return c.JSON(http.StatusCreated, newUser)
}

func DeleteUserHandler(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, &schemas.ErrorResponse{
			Message: err.Error(),
		})
	}
	index := -1
	for i, v := range database.UserList {
		if v.ID == userId {
			index = i
			break
		}
	}
	if index == -1 {
		return echo.NewHTTPError(http.StatusNotFound, &schemas.ErrorResponse{
			Message: "User not found",
		})
	}
	database.UserList = append(database.UserList[:index], database.UserList[index+1:]...)
	return c.JSON(http.StatusNoContent, nil)
}
