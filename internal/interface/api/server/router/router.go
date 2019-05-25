package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

type AppRouter interface {
	Routes() http.Handler
}

func New(mr AppRouter) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.AllowAll().Handler)

	r.Route("/api/v0", func(r chi.Router) {
		r.Mount("/messages", mr.Routes())
	})

	return r
}
