package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

func main() {
	// 连接rpc服务
	client, err := rpc.DialHTTP("tcp", "localhost:31107")
	if err != nil {
		fmt.Println("dial http failed,err:", err)
		return
	}
	// 同步调用
	arg := &Args{7, 8}
	var reply int
	err = client.Call("ServiceA.Add", arg, &reply)
	if err != nil {
		log.Printf("call ServiceA.Add failed,err:%#v\n", err)
		return
	}
	log.Printf("reply:%#v\n", reply)
	log.Printf("ServiceA.Add: %d+%d=%d\n", arg.A, arg.B, reply)

	// 异步调用
	var reply2 int
	divcall := client.Go("ServiceA.Add", arg, &reply2, nil)
	replyCall := <-divcall.Done
	log.Printf("reply:%#v, reply error:%#v\n", reply2, replyCall.Error)
}
