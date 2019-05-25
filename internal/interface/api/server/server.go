package server

import (
	"fmt"
	"net/http"

	"github.com/johejo/gohejo/envutils"
	"github.com/johejo/gohejo/logutils"
)

var logger = logutils.New()

type Server interface {
	Run() error
}

type server struct {
	router http.Handler
}

func New(router http.Handler) Server {
	return &server{
		router: router,
	}
}

func (s *server) Run() error {
	host := envutils.GetEnv("APP_HOST", "")
	port := envutils.GetEnv("APP_PORT", "8080")
	addr := fmt.Sprintf("%s:%s", host, port)

	logger.Printf("start: %s", addr)
	return http.ListenAndServe(addr, s.router)
}
