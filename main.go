package main

import (
	"fmt"

	"github.com/ckam225/golang/config"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(conf)
}
