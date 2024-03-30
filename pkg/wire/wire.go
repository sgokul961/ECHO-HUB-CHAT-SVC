//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/api"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/api/handler"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/client"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/config"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/db"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/repository"
	"github.com/sgokul961/echo-hub-chat-svc/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.ConnectMongo,
		repository.NewChatRepository,
		usecase.NewChatUseCase,
		handler.NewChatHandler,
		client.NewAuthServiceClient,
		api.NewServerHttp)
	return &api.ServerHTTP{}, nil
}
