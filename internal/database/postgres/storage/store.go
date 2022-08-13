package storage

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	dataSource string
	IUserStore
	IPostStore
	ICommentStore
}

func NewStore(dataSource string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("error openning database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return &Store{
		dataSource:    dataSource,
		IUserStore:    NewUserStore(db),
		IPostStore:    NewPostStore(db),
		ICommentStore: NewCommentStore(db),
	}, nil
}

func (s *Store) RunMigrations() error {
	m, err := migrate.New(
		"file://internal/database/postgres/migrations",
		//"postgres://postgres:postgres@host.docker.internal/golang_echo?sslmode=disable"
		s.dataSource,
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil {
		return err
	}
	return nil
}

func (s *Store) GetURL() string {
	return s.dataSource
}
