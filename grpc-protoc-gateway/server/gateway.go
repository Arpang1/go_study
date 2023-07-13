package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"grpc-protoc-gateway/pb/book"
	"io"
	"net/http"
	"os"
)

func GatewayRun() error {
	//开启日志
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	ctx := context.Background()

	//创建grpc连接
	conn, err := grpc.DialContext(ctx, "localhost:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	mux := runtime.NewServeMux()

	//注册grpc
	err = book.RegisterBookServerHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	server := http.Server{
		Addr:              "127.0.0.1:31107",
		Handler:           mux,
		TLSConfig:         nil,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	log.Info("Server Http on %s", server.Addr)

	return server.ListenAndServe()
}
