package main

import (
	"grpc-protoc-gateway/server"
	"log"
	"time"
)

func main() {
	go func() {
		err := server.GrpcRun()
		if err != nil {
			panic("开启grpc失败，" + err.Error())
		} else {
		}
	}()
	time.Sleep(2 * time.Second)
	err := server.GatewayRun()
	if err != nil {
		log.Printf("启动网关错误")
	}
}
