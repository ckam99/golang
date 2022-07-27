package models

type User struct {
	ID    int    `json:"id" xml:"id" form:"id" query:"id"`
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}
