package database

import (
	"flag"
	"fmt"
	"os"
	"project-struct/internal/entity"
	"project-struct/internal/helper"
	"project-struct/internal/service"

	"gorm.io/gorm"
)

func seederCommandHdler(db *gorm.DB, cmd *flag.FlagSet, all *bool, table *string) {
	if !*all && *table == "" {
		fmt.Println("table name is required or specify -- all for all tables")
		cmd.PrintDefaults()
		os.Exit(1)
	}
	if *all {
		RunDatabaseSeeders(db)
		return
	}

	switch *table {
	case "roles":
		if err := CreateRolesSeeder(db); err != nil {
			panic(err)
		}
		fmt.Println("CreateRoleSeeder successfully seeded")
		return
	default:
		fmt.Printf("Table %s not found\n", *table)
		fmt.Printf("%s\n", *table)
		os.Exit(1)

	}
}

func fakerCommandHdler(db *gorm.DB, cmd *flag.FlagSet, all *bool, table *string, count *int) {
	if !*all && *table == "" {
		fmt.Println("table name is required or specify --all for all tables")
		cmd.PrintDefaults()
		os.Exit(1)
	}
	if *all {
		fmt.Println("Fake all tables")
		return
	}
	if *count == 0 {
		*count = 5
	}

	switch *table {
	case "users":
		if err := service.CreateFakeUsers(db, *count); err != nil {
			panic(err)
		}
		fmt.Println("CreateFakeUsers successfully executed!")
		return
	default:
		fmt.Printf("Table %s not found\n", *table)
		fmt.Printf("%s\n", *table)
		os.Exit(1)
	}
}

func createSuperAdminCommandHdler(db *gorm.DB, cmd *flag.FlagSet, name *string, email *string, pwd *string) {
	println(*name, *pwd, *email)
	if *email == "" && *pwd == "" {
		fmt.Println("table email and password are required to create superadmin user")
		cmd.PrintDefaults()
		os.Exit(1)
	}
	if *name == "" {
		*name = "Administrator"
	}
	password, err := helper.HashPassword(*pwd)
	if err != nil {
		panic(err)
	}

	err = db.Omit("email_confirmed_at", "phone").Create(&entity.User{
		Name:     *name,
		Email:    *email,
		Password: password,
	}).Error
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Super admin successfully created!")

}

func SeederCommand(db *gorm.DB) {
	cmd := flag.NewFlagSet("db:seed", flag.ExitOnError)
	all := cmd.Bool("all", false, "Seed all tables")
	table := cmd.String("table", "", "Seed only table")
	cmd.Parse(os.Args[2:])
	seederCommandHdler(db, cmd, all, table)
}

func FakerCommand(db *gorm.DB) {
	cmd := flag.NewFlagSet("db:fake", flag.ExitOnError)
	all := cmd.Bool("all", false, "Fake all tables")
	table := cmd.String("table", "", "Fake only table")
	count := cmd.Int("count", 5, "Rows needed")
	cmd.Parse(os.Args[2:])
	fakerCommandHdler(db, cmd, all, table, count)
}

func CreateSuperadminCommand(db *gorm.DB) {
	cmd := flag.NewFlagSet("createsuperadmin", flag.ExitOnError)
	email := cmd.String("email", "", "Email address")
	pwd := cmd.String("password", "", "Password")
	name := cmd.String("name", "", "Surname")
	cmd.Parse(os.Args[2:])
	createSuperAdminCommandHdler(db, cmd, name, email, pwd)
}
