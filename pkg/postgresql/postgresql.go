package postgresql

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string
	Username string
	Password string
	Database string
	Port     string
	SSLMode  string
}

type Client interface {
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Ping(ctx context.Context) error
	Close()
}

type client struct {
	*pgxpool.Pool
	dsn string
}

func NewClient(ctx context.Context, cfg Config, maxAttempts int) (Client, error) {
	return Connection(ctx, cfg.GetURL(), maxAttempts)
}

func Connection(ctx context.Context, dsn string, maxAttempts int) (Client, error) {
	var pool *pgxpool.Pool

	var err error
	err = tryAttempt(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 5*time.Second)
	if err != nil {
		log.Fatal("error do with tries postgresql")
	}
	return &client{
		Pool: pool,
		dsn:  dsn,
	}, err
}

func tryAttempt(fn func() error, attemtps int, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--
			continue
		}
		return nil
	}
	return
}

func (cfg *Config) GetURL() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode)
}
