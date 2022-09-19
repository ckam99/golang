package database

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all methods db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) Transaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if err2 := tx.Rollback(); err2 != nil {
			return fmt.Errorf("tx error: %v, rollback: %v", err, err2)
		}
		return err
	}
	return tx.Commit()
}
