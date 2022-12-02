package http

import (
	"encoding/json"
	"example/grpc/pkg/security"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type Server struct {
	*http.ServeMux
}

func NewHTTPServer() *Server {
	server := &Server{
		ServeMux: http.NewServeMux(),
	}
	fs := http.FileServer(http.Dir("./docs/swagger"))
	server.Handle("/token", http.HandlerFunc(server.GetTokenHandler))
	server.Handle("/swagger/", http.StripPrefix("/swagger/", fs))
	return server
}

func (s *Server) GetTokenHandler(w http.ResponseWriter, r *http.Request) {
	token, err := security.GenerateToken(map[string]any{
		"id":  time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}, os.Getenv("SECRET_KEY"))

	if err != nil {
		json.NewEncoder(w).Encode(map[string]any{
			"error": err,
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]any{
		"access_token": token,
	})
	return
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
