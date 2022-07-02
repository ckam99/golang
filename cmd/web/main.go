package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ckam225/golang/webapp/config"

	"github.com/ckam225/golang/webapp/internal/handlers"
)

func main() {

	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()
	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8000"

	handlers.CreateTemplates(&appConfig)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	fmt.Printf("Server running on http://localhost%v\n", appConfig.Port)
	log.Fatal(http.ListenAndServe(appConfig.Port, nil))

}
