package internal

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) Run (port string, handler http.Handler) error {
	s.server = &http.Server{
		Addr: ":" + port,
		Handler: handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout: 30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Println("Server is running")
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown (ctx context.Context) error{
	return s.server.Shutdown(ctx)
}