package grpc

import (
	"context"
	"go20240218/grpc/service"
)

type Server struct {
	service.UnimplementedUserServiceServer
}

var _ service.UserServiceServer = &Server{} // 检查Server是否实现了UserserviceServer接口

func (s *Server) GetById(ctx context.Context, request *service.GetByIdRequest) (*service.GetByIdResponse, error) {
	return &service.GetByIdResponse{
		User: &service.User{
			Age:      123,
			Username: "abcd",
		},
	}, nil
}
