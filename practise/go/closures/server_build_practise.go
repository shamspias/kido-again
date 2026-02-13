package main

import "fmt"

type Server struct {
	Host string
	Port int
}

type Option func(*Server)

func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}
func NewServer(opts ...Option) *Server {
	svr := &Server{Host: "localhost", Port: 8080}

	for _, opt := range opts {
		opt(svr)
	}
	return svr
}

func main() {
	myServer := NewServer(WithPort(8000))
	fmt.Printf("Server starting on %s:%d\n", myServer.Host, myServer.Port)
}
