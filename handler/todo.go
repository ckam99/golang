package handler

import (
	"app/models"
	"app/repository"

	"github.com/gofiber/fiber/v2"
)

// @Summary     Welcome
// @Description  Welcome
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param        limit    query     int  false  "name search by q"
// @Success      200  {array}   models.Todo
// @Failure      404,400,401  {object}   models.Response
// @Router       /todos [get]
func GetTodosHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(repository.TodoList)
}

// @Summary     Get todo by id
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param todo_id   path int true "Todo ID"
// @Success      200  {object}   models.Todo
// @Failure      404  {object}   models.Response
// @Router       /todos/{todo_id} [get]
func GetTodoHandler(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("todo_id", 0)
	if err != nil {
		return ctx.Status(400).JSON(models.Response{
			Message: err.Error(),
		})
	}
	var todo models.Todo
	for _, t := range repository.TodoList {
		if t.ID == id {
			todo = t
			break
		}
	}
	if todo.ID == 0 {
		return ctx.Status(400).JSON(models.Response{
			Message: "Todo not found",
		})
	}
	return ctx.JSON(&todo)
}

// @Summary     CREATE TODO
// @Security  ApiKeyAuth
// @Tags         Todos
// @Accept       json
// @Produce      json
// @Param input  body models.TodoIn true "Todo ID"
// @Success      201  {object}   models.Todo
// @Failure      404,422,400  {object}   models.Response
// @Router       /todos [post]
func CreateTodoHandler(c *fiber.Ctx) error {
	var todo models.Todo
	if err := c.BodyParser(&todo); err != nil {
		return c.Status(422).JSON(&models.Response{
			Message: err.Error(),
		})
	}
	todo.ID = len(repository.TodoList) + 1
	repository.TodoList = append(repository.TodoList, todo)
	return c.Status(201).JSON(todo)
}
