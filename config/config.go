package config

import (
	"example/fiber/database"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Name        string
	Description string
}

type Config struct {
	App      *AppConfig
	Database *database.Config
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	config := &Config{
		App: &AppConfig{
			Name:        os.Getenv("APP_NAME"),
			Description: os.Getenv("APP_DESCRIPTION"),
		},
		Database: LoadDbConfig(),
	}
	return config, err

}

func LoadDbConfig() *database.Config {
	config := &database.Config{
		Host:     os.Getenv("POSTGRES_HOSTNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("POSTGRES_DB"),
		Timezone: os.Getenv("TIMEZONE"),
	}
	return config
}
