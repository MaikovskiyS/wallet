package server

import (
	"context"
	"net/http"
	"wallet/internal/config"
)

type server struct {
	s *http.Server
}

func New(cfg *config.Config) *server {

	s := &http.Server{
		Addr:         ":" + cfg.HttpServer.Port,
		ReadTimeout:  cfg.HttpServer.ReadTimeout(),
		WriteTimeout: cfg.HttpServer.WriteTimeout(),
		Handler:      nil,
	}
	return &server{s}
}
func (s *server) SetHandler(h http.Handler) {
	s.s.Handler = h
}
func (s *server) ListenAndServe() error {
	return s.s.ListenAndServe()
}
func (s *server) Shutdown() error {
	return s.s.Shutdown(context.Background())
}
