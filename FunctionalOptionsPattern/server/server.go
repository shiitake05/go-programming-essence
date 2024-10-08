package server

import (
	"log"
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

type Option func(*Server)

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		s.logger = logger
	}
}

func New(host string, port int, options ...Option) *Server {
	svr := &Server{
		host: host,
		port: port,
	}
	for _, option := range options {
		option(svr)
	}
	return svr
}

func (s *Server) Start() error {
	if s.logger != nil {
		s.logger.Printf("server started")
	}
	// do something
	return nil
}
