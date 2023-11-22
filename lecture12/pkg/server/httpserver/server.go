package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultAddr            = ":80"
	_defaultShutDownTimeout = 3 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutDownTimeout time.Duration
}

func New(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Addr:         _defaultAddr,
		Handler:      handler,
		ReadTimeout:  _defaultReadTimeout,
		WriteTimeout: _defaultWriteTimeout,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutDownTimeout: _defaultShutDownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) ShutDown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutDownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
