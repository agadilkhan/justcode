package httpserver

import (
	"net"
	"time"
)

type Option func(s *Server)

func WithPort(port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

func WithReadTimeOut(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func WithShutDownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutDownTimeout = timeout
	}
}
