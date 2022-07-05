package main

import (
	"example/fiber/config"
	"fmt"
	"log"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(conf.App.Name, conf.Database.DbName, conf.Database.Timezone)
	fmt.Println("Fiber is running")
}
