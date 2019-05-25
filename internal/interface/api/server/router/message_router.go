package router

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/johejo/ges/internal/interface/api/server/handler"
)

type messageRouter struct {
	handler.MessageHandler
	router chi.Router
}

func NewMessageRouter(mh handler.MessageHandler) AppRouter {
	return &messageRouter{
		MessageHandler: mh,
		router:         chi.NewRouter(),
	}
}

func (r *messageRouter) Routes() http.Handler {
	r.router.Get("/{id}", r.MessageHandler.GetMessage)
	r.router.Get("/", r.MessageHandler.GetMessageList)
	r.router.Post("/", r.MessageHandler.CreateMessage)
	return r.router
}
