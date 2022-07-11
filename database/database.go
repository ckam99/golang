package database

import (
	"example/fiber/entity"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
	Timezone string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s timezone=%s",
		config.Host,
		config.Port,
		config.DbName,
		config.User,
		config.Password,
		config.SSLMode,
		config.Timezone,
	)
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}

func Init(config *Config, runMigration bool, runSeeders bool) *gorm.DB {
	db, err := NewConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to the database %v", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	if runMigration {
		log.Println("Running migrations")
		if err := RunMigrations(db); err != nil {
			log.Fatalf("Failed to connect to the database %v", err.Error())
		} else {
			log.Println("Migrations successfully executed")
		}
	}
	if runSeeders {
		log.Println("Running seeders")
		if err := CreateRolesSeeder(db); err != nil {
			log.Fatalf("Failed to run CreateRolesSeeder:  %v", err.Error())
		}
	}
	//db.Logger = logger.Default.LogMode(logger.Info)
	return db
}
