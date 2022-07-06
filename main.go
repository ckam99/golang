package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"example/fiber/entity"
	"example/fiber/utils"
	"fmt"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
)

func main() {
	if conf, err := config.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	} else {
		db := database.Init(conf.Database, true) // true for migration database

		// fake data example 1
		user := entity.User{
			Password:         "0000",
			EmailConfirmedAt: time.Time{},
		}
		user.EncryptPassword()
		faker.FakeData(&user)
		if err := db.Omit("id", "email_confirmed_at", "deleted_at").Create(&user).Error; err != nil {
			panic(err)
		}
		utils.Print2Json(user)

		// example faker 2

		user1 := entity.User{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Phone:    faker.Phonenumber(),
			Password: utils.HashPassword("123456"),
		}
		if err := db.Omit("email_confirmed_at").Create(&user1).Error; err != nil {
			panic(err)
		}
		utils.Print2Json(user1)

	}
	fmt.Println("Fiber is running")
}
