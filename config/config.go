package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port        string `mapstructure:"port"`
	Description string `mapstructure:"description"`
}

type DbConfig struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Db     DbConfig     `mapstructure:"database"`
	Server ServerConfig `mapstructure:"server"`
}

var vp *viper.Viper

func LoadConfig() (Config, error) {
	var config Config
	vp = viper.New()
	vp.AddConfigPath(".")
	// vp.AddConfigPath("./config") // add more config path
	vp.SetConfigName("settings")
	vp.SetConfigType("json")
	if err := vp.ReadInConfig(); err != nil {
		return Config{}, err
	}
	if err := vp.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func TestViper() {
	vp := viper.New()
	vp.AddConfigPath(".")
	vp.SetConfigName("settings")
	vp.SetConfigType("json")
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(vp.Get("database"))
	fmt.Println(vp.Get("server.port"))
	fmt.Println(vp.Get("template.0"), vp.Get("template.1"))

	vp.Set("name", "richard")
	vp.WatchConfig()
	fmt.Println(vp.Get("name"))

	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config changed:", in.Name)
	})
	vp.WatchConfig()

	for {

	}
}
