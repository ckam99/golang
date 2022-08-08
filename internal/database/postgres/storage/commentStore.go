package storage

import (
	"fmt"

	"github.com/ckam225/golang/echo/internal/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	commentTable = "comments"
)

type commentStore struct {
	*sqlx.DB
}

func NewCommentStore(db *sqlx.DB) ICommentStore {
	return &commentStore{DB: db}
}

func (s *commentStore) GetComments(limit, offset int) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := s.Select(&comments, `SELECT FROM $1 LIMIT $2 OFFSET $3`, commentTable, limit, offset); err != nil {
		return []entity.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}
	return comments, nil
}

func (s *commentStore) GetCommentsByPost(postId uuid.UUID, limit, offset int) ([]entity.Comment, error) {
	var comments []entity.Comment
	if err := s.Select(&comments, `SELECT FROM $1 WHERE post_id = $2 LIMIT $3 OFFSET $4`, commentTable, postId, limit, offset); err != nil {
		return []entity.Comment{}, fmt.Errorf("error getting comments: %w", err)
	}
	return comments, nil
}

func (s *commentStore) GetComment(id uuid.UUID) (entity.Comment, error) {
	var comment entity.Comment
	if err := s.Get(&comment, `SELECT FROM $1 WHERE id = $2`, commentTable, id); err != nil {
		return entity.Comment{}, fmt.Errorf("error getting comment: %w", err)
	}
	return comment, nil
}

func (s *commentStore) CreateComment(t *entity.Comment) error {
	if err := s.Get(&t, `INSERT $1 (id, content, votes, post_id) VALUES ($2, $3. $4, $5) RETURNING *`,
		commentTable, t.ID, t.Content, t.Votes, t.PostId); err != nil {
		return fmt.Errorf("error creating comment: %w", err)
	}
	return nil
}

func (s *commentStore) UpdateComment(t *entity.Comment) error {
	if err := s.Get(&t, `UPDATE $1 SET  content=$2, votes=$3, post_id=$4 WHERE id=$5 RETURNING *`,
		commentTable, t.Content, t.Votes, t.PostId, t.ID); err != nil {
		return fmt.Errorf("error updating comment: %w", err)
	}
	return nil
}

func (s *commentStore) DeleteComment(id uuid.UUID) error {
	if _, err := s.Exec(`DELETE FROM $1 WHERE id = $2`, commentTable, id); err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}
	return nil
}
