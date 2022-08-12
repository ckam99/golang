package storage

import (
	"github.com/ckam225/golang/echo/internal/entity"
	"github.com/google/uuid"
)

type IPostStore interface {
	GetPosts(limit, offset int) ([]entity.Post, error)
	GetPostsByUser(userId uuid.UUID, limit, offset int) ([]entity.Post, error)
	GetPost(id uuid.UUID) (entity.Post, error)
	CreatePost(t *entity.Post) error
	UpdatePost(t *entity.Post) error
	DeletePost(id uuid.UUID) error
}

type ICommentStore interface {
	GetComments(limit, offset int) ([]entity.Comment, error)
	GetCommentsByPost(postId uuid.UUID, limit, offset int) ([]entity.Comment, error)
	GetComment(id uuid.UUID) (entity.Comment, error)
	CreateComment(t *entity.Comment) error
	UpdateComment(t *entity.Comment) error
	DeleteComment(id uuid.UUID) error
}

type IUserStore interface {
	GetUsers(limit, offset int) ([]entity.User, error)
	FindUser(id uuid.UUID) (entity.User, error)
	FindUserBy(field string, value interface{}) (entity.User, error)
	CountUserBy(field string, value interface{}) (int, error)
	CreateUser(t *entity.User) error
	UpdateUser(t *entity.User) error
	DeleteUser(id uuid.UUID) error
}

type IStore interface {
	IUserStore
	IPostStore
	ICommentStore
}
