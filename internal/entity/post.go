package entity

import "github.com/google/uuid"

type Post struct {
	ID      uuid.UUID `db:"id"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
	UserId  string    `db:"user_id"`
}

type Comment struct {
	ID      uuid.UUID `db:"id"`
	Content string    `db:"content"`
	Votes   int       `db:"votes"`
	PostId  string    `db:"post_id"`
}
