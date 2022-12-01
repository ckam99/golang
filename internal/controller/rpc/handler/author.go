package handler

import (
	"context"
	"example/grpc/internal/controller/rpc/converter"
	"example/grpc/internal/controller/rpc/pb"
	"example/grpc/internal/core/entity"
	"example/grpc/internal/core/ports"
	"example/grpc/internal/core/service"
	"example/grpc/internal/provider/postgres"
	"example/grpc/pkg/postgresql"
	"example/grpc/pkg/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthorServer struct {
	pb.UnimplementedAuthorServiceServer
	service ports.AuthorService
}

func NewAuthorServer(db postgresql.Client) *AuthorServer {
	return &AuthorServer{
		service: service.NewAuthorService(
			postgres.NewAuthorRepository(db),
		),
	}
}
func (s *AuthorServer) GetAuthors(rq *pb.QueryRequest, stream pb.AuthorService_GetAuthorsServer) error {
	authors, err := s.service.GetAll(context.Background())
	if err != nil {
		return status.Errorf(codes.Internal, "failed to fetch authors: %s", err)
	}
	for _, author := range authors {
		if err := stream.Send(converter.ConvertAuthor(author)); err != nil {
			return status.Errorf(codes.Internal, "failed to stream authors: %s", err)
		}
	}
	return nil
}

func (s *AuthorServer) CreateAuthor(ctx context.Context, in *pb.CreateAuthorRequest) (*pb.Author, error) {
	author := entity.Author{
		Name:      in.GetName(),
		Biography: in.GetBiography(),
	}
	if err := s.service.Create(ctx, &author); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to persist author:%s", err)
	}
	return converter.ConvertAuthor(author), nil
}

func (s *AuthorServer) UpdateAuthor(ctx context.Context, in *pb.UpdateAuthorRequest) (*pb.Author, error) {
	author := entity.Author{
		ID:        in.Id,
		Name:      in.GetName(),
		Biography: in.GetBiography(),
	}
	if err := s.service.Update(ctx, &author); err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get author :%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to save author:%s", err)
	}
	return converter.ConvertAuthor(author), nil
}

func (s *AuthorServer) DeleteAuthor(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	if err := s.service.Delete(ctx, in.GetValue()); err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get author :%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to delete author:%s", err)
	}
	return nil, nil
}

func (s *AuthorServer) GetAuthor(ctx context.Context, in *pb.Id) (*pb.Author, error) {
	author, err := s.service.GetByID(ctx, in.GetValue())
	if err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get author :%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to delete author:%s", err)
	}
	return converter.ConvertAuthor(author), nil
}
