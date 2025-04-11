package server

import (
	"context"
	"net/http"

	"github.com/nayan9229/go-ad-services/ssai/config"
)

// Server wraps http.Server with dependencies
type Server struct {
	AppName string
	Ctx     context.Context
	*http.Server
}

func NewServer(cfg *config.Config) *Server {
	srv := Server{}
	router := srv.routes()

	srv.Server = &http.Server{
		Addr:         cfg.Server.Address,
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	}

	return &srv
}
