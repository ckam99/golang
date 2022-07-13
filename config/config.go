package config

import (
	"github.com/gofiber/template/html"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port           string `mapstructure:"port"`
	Description    string `mapstructure:"description"`
	TemplateEngine *html.Engine
}

type DbConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Timezone string
	SSLMode  string
}

type Config struct {
	Db     DbConfig     `mapstructure:"database"`
	Server ServerConfig `mapstructure:"server"`
}

var Env *viper.Viper
var AppConfig *Config

func LoadConfig() (Config, error) {
	var config Config
	Env = viper.New()
	Env.AddConfigPath(".")
	// Env.AddConfigPath("./config") // add more config path
	Env.SetConfigName("settings")
	//Env.SetConfigType("json")
	if err := Env.ReadInConfig(); err != nil {
		return Config{}, err
	}
	if err := Env.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	config.Server.TemplateEngine = html.New("./resource/templates", ".tmpl")
	AppConfig = &config
	return config, nil
}
