package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	// conf, err := config.LoadConfig()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(conf)

	viper.AddConfigPath("./config")
	viper.SetConfigName("settings")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	fmt.Println(viper.Get("server.port"))
	fmt.Println(viper.Get("REDIS_HOST"))

}
