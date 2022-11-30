package rpc

import (
	"context"
	"example/grpc/internal/controller/rpc/converter"
	"example/grpc/internal/controller/rpc/protobuf"
	"example/grpc/internal/core/entity"
	"example/grpc/pkg/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GetAuthors(qr *protobuf.QueryRequest, stream protobuf.AppService_GetAuthorsServer) error {
	authors, err := s.authorService.GetAll(context.Background())
	if err != nil {
		return status.Errorf(codes.Internal, "failed to fetch authors: %s", err)
	}
	for _, author := range authors {
		if err := stream.SendMsg(converter.ConvertAuthor(author)); err != nil {
			return status.Errorf(codes.Internal, "failed to stream authors: %s", err)
		}
	}
	return nil
}

func (s *Server) CreateAuthor(ctx context.Context, in *protobuf.CreateAuthorRequest) (*protobuf.Author, error) {
	author := entity.Author{
		Name:      in.GetName(),
		Biography: in.GetBiography(),
	}
	if err := s.authorService.Create(ctx, &author); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to persist author:%s", err)
	}
	return converter.ConvertAuthor(author), nil
}

func (s *Server) GetAuthor(ctx context.Context, in *protobuf.Id) (*protobuf.Author, error) {
	author, err := s.authorService.GetByID(ctx, in.GetValue())
	if err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get author :%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to delete author:%s", err)
	}
	return converter.ConvertAuthor(author), nil
}

func (s *Server) UpdateAuthor(ctx context.Context, in *protobuf.UpdateAuthorRequest) (*protobuf.Author, error) {
	author := entity.Author{
		ID:        in.Id,
		Name:      in.GetName(),
		Biography: in.GetBiography(),
	}
	if err := s.authorService.Update(ctx, &author); err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get author :%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to save author:%s", err)
	}
	return converter.ConvertAuthor(author), nil
}

func (s *Server) DeleteAuthor(ctx context.Context, in *protobuf.Id) (*protobuf.Empty, error) {
	if err := s.authorService.Delete(ctx, in.GetValue()); err != nil {
		if err == utils.ErrNoEntity {
			return nil, status.Errorf(codes.NotFound, "failed to get author :%s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to delete author:%s", err)
	}
	return &protobuf.Empty{}, nil
}
