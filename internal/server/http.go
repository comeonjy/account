package server

import (
	"context"
	"net/http"
	"time"

	"github.com/comeonjy/account/api/v1"
	"github.com/comeonjy/account/configs"
	"github.com/comeonjy/go-kit/pkg/xenv"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/go-kit/pkg/xmiddleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func NewHttpServer(ctx context.Context, conf configs.Interface, logger *xlog.Logger) *http.Server {
	mux := runtime.NewServeMux(runtime.WithErrorHandler(xmiddleware.HttpErrorHandler(logger)))
	server := http.Server{
		Addr:              ":" + xenv.GetEnv(xenv.HttpPort),
		Handler:           xmiddleware.HttpUse(mux, HttpToken, xmiddleware.HttpLogger(xenv.GetEnv(xenv.TraceName), logger)),
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      2 * time.Second,
	}
	if err := v1.RegisterAccountHandlerFromEndpoint(ctx, mux, "localhost:"+xenv.GetEnv(xenv.GrpcPort), []grpc.DialOption{grpc.WithInsecure()}); err != nil {
		panic("RegisterSchedulerHandlerFromEndpoint" + err.Error())
	}
	return &server
}

func HttpToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set(runtime.MetadataHeaderPrefix+"Token", r.Header.Get("Token"))
		next.ServeHTTP(w, r)
	})
}
