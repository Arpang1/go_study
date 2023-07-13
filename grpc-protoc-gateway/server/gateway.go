package server

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"grpc-protoc-gateway/header"
	"grpc-protoc-gateway/middleware"
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
	//在gateway的拦截器
	smo := []runtime.ServeMuxOption{
		//所有的请求都会经过这个中间件，相当于拦截器
		runtime.WithForwardResponseOption(middleware.Foeward),
		//当路由发送错误时，会调用这个函数，我们可以处理
		runtime.WithRoutingErrorHandler(middleware.RoutingErrorHandler),
		//保留我们需要的请求头信息，到整个请求的上下文中
		runtime.WithIncomingHeaderMatcher(middleware.IncomingHeaderMatcher),
	}

	//这里添加了一个runtime.NewServeMux()，这个是grpc-gateway的核心，它是一个http的多路复用器，可以注册多个grpc服务
	mux := runtime.NewServeMux(smo...)
	//有什么用呢？我们可以通过mux去定义路由，定义一个上传路由和下载路由，然后通过grpc-gateway去调用grpc服务
	//上传一个文件
	if err := mux.HandlePath(http.MethodPost, "/v1/book/upload", header.Upload); err != nil {
		return err
	}
	//下载一个文件
	if err := mux.HandlePath(http.MethodGet, "/v1/book/download/{name}", header.Download); err != nil {
		return err
	}
	//注册grpc
	err = book.RegisterBookServerHandler(ctx, mux, conn)
	if err != nil {
		return err
	}
	server := http.Server{
		Addr:    "127.0.0.1:31107",
		Handler: mux,
	}
	log.Info("Server Http on %s", server.Addr)

	return server.ListenAndServe()
}
