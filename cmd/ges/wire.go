//+build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/johejo/ges/internal/application/usecase"
	"github.com/johejo/ges/internal/infrastructure/mysql"
	"github.com/johejo/ges/internal/interface/api/server"
	"github.com/johejo/ges/internal/interface/api/server/handler"
	"github.com/johejo/ges/internal/interface/api/server/router"
)

func NewServer() server.Server {
	wire.Build(
		mysql.OpenConnection,
		mysql.NewMessageRepository,
		usecase.NewMessageUseCase,
		handler.NewMessageHandler,
		router.NewMessageRouter,
		router.New,
		server.New,
	)
	return nil
}
