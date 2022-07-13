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

var DefaultDb *gorm.DB

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

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}

func RunMigrations(db *gorm.DB) {
	log.Println("Running migrations")
	if err := Migrate(db); err != nil {
		log.Fatalf("Failed to connect to the database %v", err.Error())
	} else {
		log.Println("Migrations successfully executed")
	}
}

func RunDatabaseSeeders(db *gorm.DB) {
	log.Println("Running seeders")
	if err := DatabaseSeeder(db); err != nil {
		log.Fatalf("Failed to run DatabaseSeeder:  %v", err.Error())
	} else {
		log.Println("Seeders successfully executed")
	}
}

func Init(config *Config, runMigration bool) *gorm.DB {
	db, err := NewConnection(config)
	if err != nil {
		log.Fatalf("Failed to connect to the database %v", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database successfully")
	if runMigration {
		RunMigrations(db)
	}
	DefaultDb = db
	//db.Logger = logger.Default.LogMode(logger.Info)
	return db
}
