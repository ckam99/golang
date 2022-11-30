// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: author.proto

package protobuf

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AuthorServiceClient is the client API for AuthorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorServiceClient interface {
	GetAuthors(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (AuthorService_GetAuthorsClient, error)
	CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error)
	DeleteAuthor(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
	GetAuthor(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Author, error)
}

type authorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorServiceClient(cc grpc.ClientConnInterface) AuthorServiceClient {
	return &authorServiceClient{cc}
}

func (c *authorServiceClient) GetAuthors(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (AuthorService_GetAuthorsClient, error) {
	stream, err := c.cc.NewStream(ctx, &AuthorService_ServiceDesc.Streams[0], "/protobuf.AuthorService/GetAuthors", opts...)
	if err != nil {
		return nil, err
	}
	x := &authorServiceGetAuthorsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthorService_GetAuthorsClient interface {
	Recv() (*Author, error)
	grpc.ClientStream
}

type authorServiceGetAuthorsClient struct {
	grpc.ClientStream
}

func (x *authorServiceGetAuthorsClient) Recv() (*Author, error) {
	m := new(Author)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authorServiceClient) CreateAuthor(ctx context.Context, in *CreateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/protobuf.AuthorService/CreateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) UpdateAuthor(ctx context.Context, in *UpdateAuthorRequest, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/protobuf.AuthorService/UpdateAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) DeleteAuthor(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/protobuf.AuthorService/DeleteAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorServiceClient) GetAuthor(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Author, error) {
	out := new(Author)
	err := c.cc.Invoke(ctx, "/protobuf.AuthorService/GetAuthor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorServiceServer is the server API for AuthorService service.
// All implementations must embed UnimplementedAuthorServiceServer
// for forward compatibility
type AuthorServiceServer interface {
	GetAuthors(*QueryRequest, AuthorService_GetAuthorsServer) error
	CreateAuthor(context.Context, *CreateAuthorRequest) (*Author, error)
	UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Author, error)
	DeleteAuthor(context.Context, *Id) (*Empty, error)
	GetAuthor(context.Context, *Id) (*Author, error)
	mustEmbedUnimplementedAuthorServiceServer()
}

// UnimplementedAuthorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorServiceServer struct {
}

func (UnimplementedAuthorServiceServer) GetAuthors(*QueryRequest, AuthorService_GetAuthorsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAuthors not implemented")
}
func (UnimplementedAuthorServiceServer) CreateAuthor(context.Context, *CreateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) UpdateAuthor(context.Context, *UpdateAuthorRequest) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) DeleteAuthor(context.Context, *Id) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) GetAuthor(context.Context, *Id) (*Author, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthor not implemented")
}
func (UnimplementedAuthorServiceServer) mustEmbedUnimplementedAuthorServiceServer() {}

// UnsafeAuthorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorServiceServer will
// result in compilation errors.
type UnsafeAuthorServiceServer interface {
	mustEmbedUnimplementedAuthorServiceServer()
}

func RegisterAuthorServiceServer(s grpc.ServiceRegistrar, srv AuthorServiceServer) {
	s.RegisterService(&AuthorService_ServiceDesc, srv)
}

func _AuthorService_GetAuthors_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QueryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthorServiceServer).GetAuthors(m, &authorServiceGetAuthorsServer{stream})
}

type AuthorService_GetAuthorsServer interface {
	Send(*Author) error
	grpc.ServerStream
}

type authorServiceGetAuthorsServer struct {
	grpc.ServerStream
}

func (x *authorServiceGetAuthorsServer) Send(m *Author) error {
	return x.ServerStream.SendMsg(m)
}

func _AuthorService_CreateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuthorService/CreateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).CreateAuthor(ctx, req.(*CreateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_UpdateAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).UpdateAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuthorService/UpdateAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).UpdateAuthor(ctx, req.(*UpdateAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_DeleteAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).DeleteAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuthorService/DeleteAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).DeleteAuthor(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorService_GetAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorServiceServer).GetAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.AuthorService/GetAuthor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorServiceServer).GetAuthor(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthorService_ServiceDesc is the grpc.ServiceDesc for AuthorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.AuthorService",
	HandlerType: (*AuthorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAuthor",
			Handler:    _AuthorService_CreateAuthor_Handler,
		},
		{
			MethodName: "UpdateAuthor",
			Handler:    _AuthorService_UpdateAuthor_Handler,
		},
		{
			MethodName: "DeleteAuthor",
			Handler:    _AuthorService_DeleteAuthor_Handler,
		},
		{
			MethodName: "GetAuthor",
			Handler:    _AuthorService_GetAuthor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAuthors",
			Handler:       _AuthorService_GetAuthors_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "author.proto",
}
