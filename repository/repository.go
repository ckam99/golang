package repository

import (
	"app/models"
)

var TodoList []models.Todo = []models.Todo{
	{ID: 1, Title: "Eat", Completed: false},
	{ID: 2, Title: "Sport", Completed: true},
	{ID: 3, Title: "Cinema", Completed: false},
	{ID: 4, Title: "Sleep", Completed: false},
}
