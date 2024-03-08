package main

import "net"

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

func NewUser(conn net.Conn, s *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,

		server: s,
	}

	//启动监听当前user channel 消息的goroutine
	go user.ListenMessage()

	return user
}

func (user *User) Online() {
	//用户上线, 将用户加入到oneLineMap中
	user.server.mapLock.Lock()
	user.server.OnLineMap[user.Name] = user
	user.server.mapLock.Unlock()

	//广播当前用户上线消息
	user.server.BroadCast(user, "is online")
}

func (user *User) Offline() {

	user.server.mapLock.Lock()
	delete(user.server.OnLineMap, user.Name)
	user.server.mapLock.Unlock()

	user.server.BroadCast(user, "is offline")
}

func (user *User) DoMsg(msg string) {
	user.server.BroadCast(user, msg)
}

func (user *User) ListenMessage() {
	for {
		msg := <-user.C

		user.conn.Write([]byte(msg + "\n"))
	}
}
