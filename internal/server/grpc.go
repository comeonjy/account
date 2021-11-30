package server

import (
	"time"

	"github.com/comeonjy/go-kit/grpc/reloadconfig"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/google/wire"
	"google.golang.org/grpc"

	"github.com/comeonjy/account/api/v1"
	"github.com/comeonjy/account/configs"
	"github.com/comeonjy/account/internal/service"
	"github.com/comeonjy/account/pkg/consts"
)

var ProviderSet = wire.NewSet(NewGrpcServer, NewHttpServer)

func NewGrpcServer(srv *service.AccountService, conf configs.Interface,logger *xlog.Logger) *grpc.Server {
	server := grpc.NewServer(
		grpc.ConnectionTimeout(2*time.Second),
		grpc.ChainUnaryInterceptor(
			xmiddleware.GrpcLogger(consts.TraceName,logger), xmiddleware.GrpcValidate, xmiddleware.GrpcRecover(logger), xmiddleware.GrpcAuth, xmiddleware.GrpcApm(conf.Get().ApmUrl, consts.AppName, consts.AppVersion, xenv.GetEnv(consts.AppEnv))),
	)
	v1.RegisterAccountServer(server, srv)
	reloadconfig.RegisterReloadConfigServer(server,reloadconfig.NewServer(conf))
	return server
}
