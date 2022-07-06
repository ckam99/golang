package handler

import (
	"example/fiber/repository"
	"example/fiber/schema"
	"example/fiber/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHanlder struct {
	Repo repository.UserRepository
}

func (r *UserHanlder) GetUsersHandler(c *fiber.Ctx) error {
	queryParam := schema.UserFilterParam{
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
	return c.Status(fiber.StatusOK).JSON(schema.UserListResponse(users))
}

func (r *UserHanlder) CreateUserHandler(c *fiber.Ctx) error {
	body := schema.UserRegister{}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(utils.SetHttpError(err.Error()))
	}
	if user, err := r.Repo.CreateUser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.SetHttpError(err.Error()))
	} else {
		return c.Status(fiber.StatusCreated).JSON(schema.UserReponse(user))
	}
}

func (r *UserHanlder) GetUserHandler(c *fiber.Ctx) error {
	if id, err := c.ParamsInt("id"); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.SetHttpError(err.Error()))
	} else {
		user, err := r.Repo.GetUserByID(uint(id))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(utils.SetHttpError(err.Error()))
		}
		return c.Status(fiber.StatusOK).JSON(schema.UserReponse(user))
	}

}

func (r *UserHanlder) UpdateUserHandler(c *fiber.Ctx) error {
	body := schema.UserUpdate{}
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

func (r *UserHanlder) DeleteUserHandler(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	if err := r.Repo.DeleteUser(uint(id), true); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	return c.Status(fiber.StatusNoContent).JSON(nil)
}
