package option

import (
	"crypto/tls"
	"golang-first-step/about_language/functional/model"
	"time"
)

func Protocol(p string) model.Option {
	return func(s *model.Server) {
		s.Protocol = p
	}
}
func Timeout(timeout time.Duration) model.Option {
	return func(s *model.Server) {
		s.Timeout = timeout
	}
}
func MaxConns(maxconns int) model.Option {
	return func(s *model.Server) {
		s.MaxConnections = maxconns
	}
}
func TLS(tls *tls.Config) model.Option {
	return func(s *model.Server) {
		s.TLSConfig = tls
	}
}
