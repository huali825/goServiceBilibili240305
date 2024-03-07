package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	//在线用户的列表
	OnLineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

// NewServer 新建一个 提供创建一个server的接口
func NewServer(ip string, prot int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      prot,
		OnLineMap: make(map[string]*User),
		Message:   make(chan string),
	}

	return server
}

func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		s.mapLock.Lock()
		for _, cli := range s.OnLineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	s.Message <- sendMsg
}

func (s *Server) Handler(conn net.Conn) {
	//在这里写 命中的业务
	//fmt.Println("链接建立成功")

	user := NewUser(conn)

	//用户上线, 将用户加入到oneLineMap中
	s.mapLock.Lock()
	s.OnLineMap[user.Name] = user
	s.mapLock.Unlock()

	s.BroadCast(user, "is online")
}

func (s *Server) Start() {
	l, err := net.Listen("tcp", fmt.Sprint(s.Ip, ":", s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	defer l.Close()

	go s.ListenMessage()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}

		go s.Handler(conn)
	}
}
