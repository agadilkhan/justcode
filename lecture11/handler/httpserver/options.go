package httpserver

import "time"

type Option func(*Server)

func WithPort(port string) Option {
	return func(server *Server) {
		server.server.Addr = port
	}
}

func WithShutdownTimeout(timeout time.Duration) Option {
	return func(server *Server) {
		server.shutdownTimeout = timeout
	}
}
