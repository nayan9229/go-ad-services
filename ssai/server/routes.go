package server

import (
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nayan9229/go-ad-services/shared/utilities"
)

func (s *Server) routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))
	r.Use(utilities.Logger())

	r.Get("/", utilities.Health)
	r.Get("/healthz", utilities.Health)
	return r
}
