package main

import (
	"database/sql"
	"example-asyncq/handler"
	"example-asyncq/worker"
	"log"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/hibiken/asynq"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	DbUrl    string
	RedisUrl string
}

func loadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	DB_URL := os.Getenv("DB_URL")
	REDIS_URL := os.Getenv("REDIS_URL")
	if DB_URL == "" {
		DB_URL = "postgres:/postgres:postgres@host.docker.internal:5432/demo?ssl_mode=disable"
	}
	if REDIS_URL == "" {
		REDIS_URL = "host.docker.internal:6379"
	}
	return &Config{
		DbUrl:    DB_URL,
		RedisUrl: REDIS_URL,
	}
}

func main() {

	cfg := loadConfig()

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// distributor
	redisOption := asynq.RedisClientOpt{
		Addr: cfg.RedisUrl,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOption)

	// http route handlers
	app := fiber.New()
	app.Get("/", adaptor.HTTPHandlerFunc(handler.Greet))
	email := handler.NewEmailHandler(db, taskDistributor)
	app.Get("/email", email.SendVerifyEmail)

	// task processor
	go startTaskProcessor(redisOption, db)
	log.Fatal(app.Listen(":9000"))
}

func startTaskProcessor(redisOption asynq.RedisClientOpt, db *sql.DB) {
	tProcess := worker.NewRedisTaskProcessor(redisOption, db)
	log.Println("task processor started")
	if err := tProcess.Start(); err != nil {
		log.Fatalf("FAILED TO START TASK PROCESSOR: %s", err)
	}
}
