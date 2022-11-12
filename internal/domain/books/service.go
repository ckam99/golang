package books

import (
  "database/sql"
  "context"
)


type service struct{
  repo Repository
}

type NewService(db *sql.DB) Service{
  return &service{
    repo: NewRepository(db)
  }
}

func(s *service) GetAll(ctx context.Context, param *QueryFilterDTO)([]Book,error){
  return s.repo.GetAll(param),nil
}

func(s *service) Create(ctx context.Context, dto CreateDTO) error{
  panic("not implemented")
}

func(s *service) Find(ctx context.Context, id int64)(Book,error){
  panic("not implemented")
}

func(s *service) Update(ctx context.Context, dto UpdateDTO) error{
  panic("not implemented")
}

func(s *service) Delete(ctx context.Context, id int64) error{
  panic("not implemented")
}