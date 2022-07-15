package models

type Todo struct {
	ID        int    `json:"id" example:"1"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodoIn struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
