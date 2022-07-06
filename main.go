package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"example/fiber/entity"
	"example/fiber/utils"
	"fmt"
	"log"
)

func main() {

	if conf, err := config.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	} else {
		db := database.Init(conf.Database, true) // true for migration database
		user := entity.User{
			Name:     "User 7",
			Email:    "user7@mail.ru",
			Phone:    "+757098",
			Password: "DJFS454009",
		}
		user.EncryptPassword()
		if err := db.Omit("email_confirmed_at").Create(&user).Error; err != nil {
			panic(err.Error())
		} else {
			utils.Print2Json(user)
		}
	}
	fmt.Println("Fiber is running")
}
