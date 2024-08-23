package grpc

import "go20240218/grpc/service"

type Server struct {
	service.UnimplementedUserServiceServer
}
