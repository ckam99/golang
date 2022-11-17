package database

import (
	"context"
	"main/pkg/clients/postgresql"
	"main/pkg/migrator"

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
	migrator, err := migrator.New("postgres", db.config.GetURL(), db.migrateDir)
	if err != nil {
		return err
	}
	defer migrator.Close()
	return migrator.Migrate()
}

func (db *Database) Rollback() error {
	migrator, err := migrator.New("postgres", db.config.GetURL(), db.migrateDir)
	if err != nil {
		return err
	}
	defer migrator.Close()
	err = migrator.Rollback()
	if err != nil {
		return err
	}
	return nil
}
