package repository

import (
	"github.com/ckam225/golang/sqlx/database"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
	IPersonRepository
}

func New(cfg *database.Config) (*Repository, error) {
	db, err := database.Connect(cfg)
	if err != nil {
		return nil, err
	}
	repo := &Repository{
		db:                db,
		IPersonRepository: NewPersonRepository(db),
	}
	return repo, nil
}

func (r *Repository) Destroy() {
	defer r.db.Close()
}
