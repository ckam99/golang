package service

import "github.com/ckam225/golang/echo/internal/database/postgres/storage"

type Service struct {
	IUserService
}

func NewService(store *storage.Store) *Service {
	return &Service{
		IUserService: UserService(store),
	}
}
