package storage

import (
	"fmt"

	"github.com/ckam225/golang/echo/internal/entity"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	postTable = "posts"
)

type postStore struct {
	*sqlx.DB
}

func NewPostStore(db *sqlx.DB) IPostStore {
	return &postStore{
		DB: db,
	}
}

func (p *postStore) GetPosts(limit, offset int) ([]entity.Post, error) {
	var posts []entity.Post
	query := fmt.Sprintf(`SELECT * FROM %s LIMIT $1 OFFSET $2`, postTable)
	if err := p.Select(&posts, query, limit, offset); err != nil {
		return []entity.Post{}, fmt.Errorf("error getting posts: %w", err)
	}
	return posts, nil
}

func (p *postStore) GetPostsByUser(userId uuid.UUID, limit, offset int) ([]entity.Post, error) {
	var posts []entity.Post
	query := fmt.Sprintf(`SELECT * FROM %s WHERE user_id = $1 LIMIT $2 OFFSET $3`, postTable)
	if err := p.Select(&posts, query, userId, limit, offset); err != nil {
		return []entity.Post{}, fmt.Errorf("error getting posts: %w", err)
	}
	return posts, nil
}

func (p *postStore) GetPost(id uuid.UUID) (entity.Post, error) {
	var post entity.Post
	query := fmt.Sprintf(`SELECT * FROM %s WHERE id = $1`, postTable)
	if err := p.Get(&post, query, id); err != nil {
		return entity.Post{}, fmt.Errorf("error getting post: %w", err)
	}
	return post, nil
}

func (p *postStore) CreatePost(t *entity.Post) error {
	query := fmt.Sprintf(`INSERT %s (id, title, content, votes, user_id) VALUES ($1, $2. $3, $4, $5) RETURNING *`, postTable)
	if err := p.Get(&t, query, t.ID, t.Title, t.Content, t.Votes, t.UserId); err != nil {
		return fmt.Errorf("error creating post: %w", err)
	}
	return nil
}

func (p *postStore) UpdatePost(t *entity.Post) error {
	query := fmt.Sprintf(`UPDATE %s SET title=$1, content=$2, votes=$3, user_id=$4 WHERE id=$5 RETURNING *`, postTable)
	if err := p.Get(&t, query, t.Title, t.Content, t.Votes, t.UserId, t.ID); err != nil {
		return fmt.Errorf("error updating post: %w", err)
	}
	return nil
}

func (p *postStore) DeletePost(id uuid.UUID) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id=$1 RETURNING *`, postTable)
	if _, err := p.Exec(query); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
