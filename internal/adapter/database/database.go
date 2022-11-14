package database

import (
	"context"
	"main/pkg/clients/postgresql"
	"main/pkg/migrate"

	_ "github.com/lib/pq"
)

type Database struct {
	postgresql.Client
	config     postgresql.Config
	migrateDir string
}

func Connection(ctx context.Context, cfg postgresql.Config) (*Database, error) {
	c, err := postgresql.NewClient(ctx, cfg, 3)
	//rootDir, _ := os.Getwd()
	db := &Database{
		Client:     c,
		config:     cfg,
		migrateDir: "./internal/migration",
	}
	return db, err
}

func (db *Database) Migrate() error {
	migration, err := migrate.New("postgres", db.config.GetURL(), db.migrateDir, &migrate.Config{})
	if err != nil {
		return err
	}
	return migration.Migrate()
}

func (db *Database) Rollback() error {
	migration, err := migrate.New("postgres", db.config.GetURL(), db.migrateDir, &migrate.Config{})
	if err != nil {
		return err
	}
	err = migration.Rollback()
	if err != nil {
		return err
	}
	return nil
}
