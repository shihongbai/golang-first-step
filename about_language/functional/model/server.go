package model

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Server struct {
	Addr           string
	Port           int
	Protocol       string
	Timeout        time.Duration
	MaxConnections int
	TLSConfig      *tls.Config
}

func (s Server) ToString() string {
	return fmt.Sprintf("Server { Addr: %s, Port: %d, Protocol: %s, Timeout: %s, MaxConnections: %d, TLSConfig: %v }",
		s.Addr, s.Port, s.Protocol, s.Timeout, s.MaxConnections, s.TLSConfig)
}
func NewServer(addr string, port int, options ...func(*Server)) *Server {
	server := &Server{Addr: addr, Port: port}

	for _, option := range options {
		option(server)
	}

	return server
}

type Option func(*Server)
