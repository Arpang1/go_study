package server

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"grpc-protoc-gateway/middleware"
	"grpc-protoc-gateway/pb/book"
	"io"
	"net"
	"os"
)

func GrpcRun() error {
	//打印日志
	log := grpclog.NewLoggerV2(os.Stdout, io.Discard, io.Discard)
	grpclog.SetLoggerV2(log)

	//监听端口 tcp
	l, err := net.Listen("tcp", ":9080")
	if err != nil {
		return err
	}
	defer func(l net.Listener) {
		err := l.Close()
		if err != nil {
			return
		}
	}(l)
	op := []grpc.ServerOption{
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	}
	//创建grpac的server服务
	// grpc.NewServer(opt ...ServerOption)  -->  opt为一些加密密钥
	s := grpc.NewServer(op...)
	//创捷一个book的Server
	g := NewBookService()
	//注册服务
	book.RegisterBookServerServer(s, g)

	log.Info("Serving gRPC on %s", l.Addr())
	return s.Serve(l)
}
