package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Server struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port string `json:"port"`
}

type Database struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
	Timezone string `json:"timezone"`
	SSLMode  string `json:"sslmode"`
}

type Configuration struct {
	Server   *Server
	Database *Database
}

func LoadDbConfig() *Database {
	config := &Database{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Name:     os.Getenv("DB_NAME"),
		Timezone: os.Getenv("TIMEZONE"),
	}
	return config
}

func LoadConfig() (*Configuration, error) {
	err := godotenv.Load(".env")
	config := &Configuration{
		Server: &Server{
			Name: os.Getenv("APP_NAME"),
			Host: os.Getenv("APP_HOST"),
			Port: os.Getenv("APP_PORT"),
		},
		Database: LoadDbConfig(),
	}
	return config, err

}

func (db *Database) GetURL() string {
	// "host=%s port=%s dbname=%s user=%s password=%s sslmode=%s timezone=%s",
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=%s",
		db.Driver,
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		db.SSLMode,
	)
}
