package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"example/fiber/entity"
	"example/fiber/repository"
	"example/fiber/utils"
	"fmt"
	"log"
)

func main() {
	if conf, err := config.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	} else {
		db := database.Init(conf.Database, true) // true for migration database
		userRepo := repository.UserRepository{
			Query: db,
			Filter: repository.UserFilterParam{
				Limit: 100,
				Skip:  0,
			},
		}
		users, err := userRepo.GetAllUsers()
		if err != nil {
			panic(err)
		}
		utils.PrintJson(entity.ToUserListResponse(users))
	}
	fmt.Println("Fiber is running")
}
