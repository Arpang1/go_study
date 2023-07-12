package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type ServiceA struct{}

// Add函数为ServiceA的方法
func (t *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	service := new(ServiceA)
	// 注册rpc服务
	err := rpc.Register(service)
	if err != nil {
		return
	}
	// rpc服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 启动http服务
	err = http.ListenAndServe(":31107", nil)
	if err != nil {
		log.Printf("http start filed %#v\n", err)
		return
	}
}
