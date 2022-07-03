package database

import (
	"example/gorm/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
	TimeZone string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.Host,
		config.Port,
		config.User,
		config.Password,
		config.DBName,
		config.SSLMode,
		config.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}

func Init(config *Config, migrate bool) *gorm.DB {
	db, err := NewConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to the database %v", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	if migrate {
		log.Println("Running migrations")
		if err := RunMigrations(db); err != nil {
			log.Fatalf("Failed to connect to the database %v", err.Error())
		} else {
			log.Println("Migrations successfully executed")
		}
	}
	//db.Logger = logger.Default.LogMode(logger.Info)
	return db
}

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Book{})

	return err
}
