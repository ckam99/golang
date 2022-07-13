package main

import (
	"fmt"
	"os"
	"project-struct/config"
	"project-struct/internal/database"
)

const StringHelp = `
Usage:
  console help - Help
  console db:migrate - Migration database
  console db:seed - Database seeding
  console db:fake - Create fake data
  console createsuperadmin --email <email> --password <password> - Create super admin user
`

func main() {

	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	db, err := database.NewConnection(&conf.Db)
	if err != nil {
		panic(err)
	}

	if len(os.Args) < 2 {
		fmt.Printf("%s\n", StringHelp)
		os.Exit(1)
	}

	// get command

	switch os.Args[1] {
	case "help":
		fmt.Printf("%s", StringHelp)
	case "db:migrate":
		database.RunMigrations(db)
	case "db:seed":
		database.SeederCommand(db)
	case "db:fake":
		database.FakerCommand(db)
	case "createsuperadmin":
		database.CreateSuperadminCommand(db)
	default:
		fmt.Println("Option not found")
	}

}
