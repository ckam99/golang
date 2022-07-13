package database

import (
	"fmt"
	"log"
	"os"
	"project-struct/config"
	"project-struct/internal/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DefaultDB *gorm.DB

func NewConnection(dbconfig *config.DbConfig) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s timezone=%s",
		dbconfig.Host,
		dbconfig.Port,
		dbconfig.Name,
		dbconfig.Username,
		dbconfig.Password,
		dbconfig.SSLMode,
		dbconfig.Timezone,
	)
	return gorm.Open(postgres.Open(dns), &gorm.Config{})
}

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&entity.User{})
}

func runSeeders(db *gorm.DB) error {

	if err := CreateRolesSeeder(db); err != nil {
		return err
	}
	return nil
}

func RunDatabaseSeeders(db *gorm.DB) {
	log.Println("Running seeders")
	if err := runSeeders(db); err != nil {
		log.Fatalf("Failed to run CreateRolesSeeder:  %v", err.Error())
	} else {
		log.Println("Seeders successfully executed")
	}
}

func Init(dbconfig *config.DbConfig, runMigration bool) *gorm.DB {
	db, err := NewConnection(dbconfig)
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
	DefaultDB = db
	//db.Logger = logger.Default.LogMode(logger.Info)
	return db
}
