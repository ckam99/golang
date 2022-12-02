package http

import (
	"fmt"
	"net"
	"net/http"
)

type Server struct {
	*http.ServeMux
}

func NewHTTPServer() *Server {
	server := &Server{
		ServeMux: http.NewServeMux(),
	}
	fs := http.FileServer(http.Dir("./docs/swagger"))
	server.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	return server
}

func (s *Server) Serve(address string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("cannot create http network listener:%w", err)
	}
	if err = http.Serve(listener, s); err != nil {
		return fmt.Errorf("cannot start HTTP server: %w", err)
	}
	return nil
}
