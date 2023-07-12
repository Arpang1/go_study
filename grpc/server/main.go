package main

import (
	"google.golang.org/grpc"
	"net"
	"server/controller"
	"server/pb/book"
)

func main() {
	//创建一个rpc连接
	controller.NewBook("安徒生", "安徒生", "26.00", "安徒生童话故事")
	controller.NewBook("爱丽丝梦游仙境", "刘易斯", "66.00", "爱丽丝梦游仙境")
	controller.NewBook("白雪公主", "小白", "29.00", "白雪公主")
	controller.NewBook("灰姑娘", "小灰", "56.00", "灰姑娘")
	controller.NewBook("小红帽", "小红", "23.00", "小红帽")
	controller.NewBook("青蛙王子", "小青", "46.00", "青蛙王子")

	listen, err := net.Listen("tcp", ":31107")
	if err != nil {
		panic("failed to listen: " + err.Error())
	}
	//创建 grpc
	grpcserver := controller.NewGrpcServer()

	// 创建grpc服务
	grpcServer := grpc.NewServer()
	book.RegisterBookServiceServer(grpcServer, grpcserver)
	//启动服务
	if err := grpcServer.Serve(listen); err != nil {
		panic("failed to serve: " + err.Error())
	}
}
