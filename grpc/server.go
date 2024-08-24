package grpc

import (
	"context"
	servicegrpc "go20240218/grpc/service"
)

type Server struct {
	servicegrpc.UnimplementedUserServiceServer
}

var _ servicegrpc.UserServiceServer = &Server{} // 检查Server是否实现了UserserviceServer接口

func (s *Server) GetById(ctx context.Context, request *servicegrpc.GetByIdRequest) (*servicegrpc.GetByIdResponse, error) {
	return &servicegrpc.GetByIdResponse{
		User: &servicegrpc.User{
			Age:      123,
			Username: "abcd",
		},
	}, nil
}
