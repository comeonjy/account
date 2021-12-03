package server

import (
	"time"

	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"account/api/v1"
	"account/configs"
	"account/internal/service"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

func NewGrpcServer(srv *service.AccountService, conf configs.Interface, logger *xlog.Logger) *grpc.Server {
	server := grpc.NewServer(
		grpc.ConnectionTimeout(2*time.Second),
		grpc.ChainUnaryInterceptor(
			xmiddleware.GrpcLogger(xenv.GetEnv(xenv.TraceName), logger), xmiddleware.GrpcValidate, xmiddleware.GrpcRecover(logger), xmiddleware.GrpcAuth, xmiddleware.GrpcApm(conf.Get().ApmUrl, xenv.GetEnv(xenv.AppName), xenv.GetEnv(xenv.AppVersion), xenv.GetEnv(xenv.AppEnv))),
	)
	v1.RegisterAccountServer(server, srv)
	return server
}
