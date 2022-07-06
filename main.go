package main

import (
	"example/fiber/config"
	"example/fiber/database"
	"fmt"
	"log"
)

func main() {
	if conf, err := config.LoadConfig(); err != nil {
		log.Fatalln(err.Error())
	} else {
		database.Init(conf.Database, true) // true for migration database
	}
	fmt.Println("Fiber is running")
}
