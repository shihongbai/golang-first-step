package golang_IM_system

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
)

type Server struct {
	IP   string
	Port int

	OnlineUser map[string]*User
	mapLock    sync.RWMutex

	// 消息广播的channel
	Message chan string
}

func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message

		s.mapLock.Lock()

		// 广播
		for _, cli := range s.OnlineUser {
			cli.C <- msg
		}

		s.mapLock.Unlock()
	}
}

func (s *Server) BroadCast(user *User, msg string) {
	info := fmt.Sprintf("The current user is online, user name %s, msg %s", user.Name, msg)

	s.Message <- info
}

func (s *Server) Handler(conn net.Conn) {
	// 业务逻辑

	// 用户上线，并广播
	s.mapLock.Lock()
	user := NewUser(conn)
	s.OnlineUser[user.Name] = user
	s.mapLock.Unlock()

	// 广播消息
	s.BroadCast(user, "已上线")

	// 阻塞
	select {}
}

func (s *Server) Start() {
	// socket listen
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		logrus.Error(err)
	}
	defer listen.Close()

	for {
		// accept
		conn, err := listen.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}

		go s.ListenMessage()

		// do handler
		go s.Handler(conn)
	}

	// close listen
}

func NewServer(IP string, port int) *Server {
	return &Server{
		IP:         IP,
		Port:       port,
		OnlineUser: make(map[string]*User),
		Message:    make(chan string),
	}
}
