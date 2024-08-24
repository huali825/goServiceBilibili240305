package grpc

import (
	servicegrpc "go20240218/grpc/service"
	"google.golang.org/grpc"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	// 创建一个新的gRPC服务器
	server := grpc.NewServer()
	// 创建一个Server实例
	userServer := &Server{}
	// 将UserServiceServer注册到gRPC服务器上
	servicegrpc.RegisterUserServiceServer(server, userServer)
	// 监听本地 8082 端口
	l, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil {
		// 如果监听失败，则抛出异常
		panic(err)
	}
	// 启动gRPC服务器
	err = server.Serve(l)
	// 如果启动失败，则抛出异常
	panic(err)
}
