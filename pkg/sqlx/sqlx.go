package sqlx

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database struct {
	dsn   string
	query string
	DB    *sqlx.DB
}

func Postgres(host string, port int, dbname, username, password, sslmode string, timeout int) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s connect_timeout=%d",
		host, port, dbname, username, password, sslmode, timeout,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return &Database{
		dsn: dsn,
		DB:  db,
	}, nil
}

func (s *Database) RunMigrations(migrationDir string) error {
	m, err := migrate.New(
		migrationDir,
		s.dsn,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (s *Database) RollbackMigrations(migrationDir string) error {
	m, err := migrate.New(migrationDir, s.dsn)
	if err != nil {
		return err
	}
	if err := m.Down(); err != nil {
		return err
	}
	return nil
}

func (s *Database) GetDNS() string {
	return s.dsn
}

func (s *Database) Close() error {
	return s.DB.Close()
}
