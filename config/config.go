package config

import (
	"example/fiber/database"
	"html/template"
	"os"

	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Name        string
	Description string
	HtmlEngine  *html.Engine
}

type Configuration struct {
	Server   *ServerConfig
	Database *database.Config
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

func LoadConfig() (*Configuration, error) {
	engine := html.New("./resource/templates", ".tmpl")
	engine.AddFunc(
		// add unescape function
		"unescape", func(s string) template.HTML {
			return template.HTML(s)
		},
	)
	err := godotenv.Load(".env")
	config := &Configuration{
		Server: &ServerConfig{
			Name:        os.Getenv("APP_NAME"),
			Description: os.Getenv("APP_DESCRIPTION"),
			HtmlEngine:  engine,
		},
		Database: LoadDbConfig(),
	}
	return config, err

}
