package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	SSLmode  string `json:"sslmode"`
	Timeout  int    `json:"timeout"`
	Database string `json:"database"`
}

func Connect(cfg *Config) (*sqlx.DB, error) {
	databaseUrl := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d",
		cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password, cfg.SSLmode, cfg.Timeout,
	)
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		return nil, fmt.Errorf("error openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	// defer db.Close()
	return db, err
}
