package storage

import (
	"fmt"

	"github.com/ckam225/golang/fiber-sqlx/internal/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	userTable = "users"
)

type userStore struct {
	*sqlx.DB
}

func NewUserStore(db *sqlx.DB) IUserStore {
	return &userStore{
		DB: db,
	}
}

func (s *userStore) GetUsers(limit, offset int) ([]entity.User, error) {
	var users []entity.User
	query := fmt.Sprintf(`SELECT * FROM %s LIMIT $1 OFFSET $2`, userTable)
	if err := s.Select(&users, query, limit, offset); err != nil {
		return []entity.User{}, fmt.Errorf("error getting list users: %w", err)
	}
	return users, nil
}

func (s *userStore) FindUser(id uuid.UUID) (entity.User, error) {
	var user entity.User
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, userTable)
	if err := s.Get(&user, query, id); err != nil {
		return entity.User{}, fmt.Errorf("error getting user: %w", err)
	}
	return user, nil
}

func (s *userStore) FindUserBy(field string, value interface{}) (entity.User, error) {
	var user entity.User
	stmt, err := s.Preparex(fmt.Sprintf(`SELECT * FROM %s WHERE $1 = $2`, userTable))
	if err != nil {
		return entity.User{}, fmt.Errorf("error in statement getting user by %w", err)
	}
	if err := stmt.Get(&user, field, value); err != nil {
		return entity.User{}, fmt.Errorf("error getting user by %s: %w", field, err)
	}
	return user, nil
}

func (s *userStore) CreateUser(t *entity.User) error {
	query := fmt.Sprintf(`INSERT INTO %s (id, email, name) VALUES ($1, $2, $3) RETURNING *`, userTable)
	if err := s.Get(t, query, t.ID, t.Email, t.Name); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *userStore) UpdateUser(t *entity.User) error {
	query := fmt.Sprintf(`UPDATE %s SET email=$1, name=$2 WHERE id=$3 RETURNING *`, userTable)
	if err := s.Get(t, query, t.Email, t.Name, t.ID); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (s *userStore) DeleteUser(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 RETURNING *`, userTable)
	if _, err := s.Exec(query, id); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (s *userStore) CountUserBy(field string, value interface{}) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) %s WHERE $1 = $2", userTable)
	var count int
	if err := s.Get(&count, query, field, value); err != nil {
		return count, err
	}
	return count, nil
}
