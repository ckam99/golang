package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ckam225/golang/fiber/internal/config"
	"github.com/ckam225/golang/fiber/internal/http/handler"
	"github.com/ckam225/golang/fiber/internal/jobs"
)

// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// load environment variable
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// http server
	server := handler.NewHandler(cfg)
	// register jobs
	jobs.RegisterNotificationChannel()
	// unregister jobs
	defer jobs.UnregisterNotificationChannel()

	log.Fatal(server.Listen(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))))
}
