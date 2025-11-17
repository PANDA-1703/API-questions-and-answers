package http

import (
	"API-quest-answ/internal/config"
	"context"
	"fmt"
	"net/http"
)

type HttpServer struct {
	httpServer *http.Server
}

func NewServer(cfg *config.HttpServerConfig, handler http.Handler) *HttpServer {
	return &HttpServer{
		httpServer: &http.Server{
			Addr:           fmt.Sprintf(":%d", cfg.Port),
			Handler:        handler,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderBytes,
		},
	}
}

func (s *HttpServer) ListAndServe() error {
	return s.httpServer.ListenAndServe()
}

func (s *HttpServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
