// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package cmd

import (
	"account/configs"
	"account/internal/data"
	"account/internal/server"
	"account/internal/service"
	"context"
	"github.com/comeonjy/go-kit/pkg/xlog"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func InitApp(ctx context.Context, logger *xlog.Logger) *App {
	configsInterface := configs.NewConfig(ctx)
	dataData := data.NewData(configsInterface)
	accountRepo := data.NewAccountRepo(dataData)
	accountService := service.NewAccountService(configsInterface, accountRepo, logger)
	grpcServer := server.NewGrpcServer(accountService, configsInterface, logger)
	httpServer := server.NewHttpServer(ctx, configsInterface, logger)
	app := newApp(ctx, grpcServer, httpServer, configsInterface)
	return app
}
