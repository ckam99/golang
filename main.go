package main

import (
	"example/gorm/database"
	"example/gorm/models"
	"example/gorm/repositories"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln(err.Error())
	}
	config := &database.Config{
		Host:     os.Getenv("POSTGRES_HOSTNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("POSTGRES_DB"),
		TimeZone: os.Getenv("TIMEZONE"),
	}
	db := database.Init(config, true)

	userRepo := repositories.UserRepository{
		Db: db,
	}

	/* crrate single user */
	if user, err := userRepo.CreateUser(&models.User{
		Name:  "Alain SHuer",
		Email: "alainsh 90@mail.ru",
	}); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(user)
	}

	/*
	 create many users
	*/
	if users, err := userRepo.CreateUsers(&[]models.User{
		{Name: "User 15", Email: "use15@mail.ru"},
		{Name: "User 16", Email: "user16@mail.ru"},
	}); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Created many users", users)
	}

	/* Get list users */
	fmt.Println(userRepo.GetUsers(100, 0))

	/* fetch single user by Object */

	testUser := models.User{}
	testUser.ID = 1
	if user, err := userRepo.GetUser(&testUser); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("======fetch single user by Object =======")
		fmt.Println(user)
	}

	/* Get User by ID */
	if user, err := userRepo.GetUserByID(2); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("======Get User by ID =========")
		fmt.Println(user)
	}

	/* Get User by Email */
	if user, err := userRepo.GetUserByEmail("user7@mail.ru"); err != nil {
		panic(err.Error())
	} else {
		fmt.Println("======Get User by Email ========")
		fmt.Println(user)
	}

	/* update user with payload */
	if user, err := userRepo.UpdateUser(
		1, models.UpdateUserSchema{Name: "Mariam Sala"}); err != nil {
		panic(err)
	} else {
		fmt.Println("======update user with payload ========")
		fmt.Println(user)
	}

}
