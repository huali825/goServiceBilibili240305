package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// NewServer 新建一个 提供创建一个server的接口
func NewServer(ip string, prot int) *Server {
	server := &Server{
		Ip:   ip,
		Port: prot,
	}

	return server
}

func (s *Server) Handler(conn net.Conn) {
	//在这里写 命中的业务
	fmt.Println("链接建立成功")
}

func (s *Server) Start() {
	l, err := net.Listen("tcp", fmt.Sprint(s.Ip, ":", s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}

		go s.Handler(conn)
	}
}
