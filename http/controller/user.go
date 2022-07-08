package controller

import (
	"example/fiber/http/request"
	"example/fiber/http/response"
	"example/fiber/repository"
	"example/fiber/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo repository.UserRepository
}

func (r *UserController) GetUsersHandler(c *fiber.Ctx) error {
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
	users, err := r.Repo.GetAllUsers(queryParam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.SetHttpError(err.Error()))
	}
	return c.Status(fiber.StatusOK).JSON(response.ParseUserListEntity(users))
}

func (r *UserController) CreateUserHandler(c *fiber.Ctx) error {
	body := request.CreateUser{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.SetHttpError(err.Error()))
	}
	if errors := utils.ValidateSchema(body); errors != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(errors)
	}
	if user, err := r.Repo.CreateUser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.SetHttpError(err.Error()))
	} else {
		return c.Status(fiber.StatusCreated).JSON(response.ParseUserEntity(user))
	}
}

func (r *UserController) GetUserHandler(c *fiber.Ctx) error {
	if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.SetHttpError(err.Error()))
	} else {
		user, err := r.Repo.GetUserByID(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(utils.SetHttpError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(response.ParseUserEntity(user))
	}
}

func (r *UserController) UpdateUserHandler(c *fiber.Ctx) error {
	body := request.UpdateUser{}
	id, _ := c.ParamsInt("id", 0)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}
	user, err := r.Repo.UpdateUser(uint(id), &body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.SetHttpError(err.Error()))
	}
	return c.Status(fiber.StatusAccepted).JSON(&user)
}

func (r *UserController) DeleteUserHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := r.Repo.DeleteUser(uint(id), true); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
