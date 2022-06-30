package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	todosJson := /* json */ `
	[
	  {
		"userId": 1,
		"id": 1,
		"title": "delectus aut autem",
		"completed": false
	  },
	  {
		"userId": 1,
		"id": 2,
		"title": "quis ut nam facilis et officia qui",
		"completed": false
	  },
	  {
		"userId": 1,
		"id": 3,
		"title": "fugiat veniam minus",
		"completed": false
	  }
	]`

	var todos []Todo

	err := json.Unmarshal([]byte(todosJson), &todos)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("type: %T \n %v\n", todos, todos)

	var todo Todo
	todo.Id = 4
	todo.UserId = 9
	todo.Title = "Lorem ipsum"
	todo.Completed = true

	todos = append(todos, todo)
	jsonFromTodo, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("type: %T \n %s\n", jsonFromTodo, jsonFromTodo)

}
