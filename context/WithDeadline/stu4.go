package main

import (
	"context"
	"fmt"
	"time"
)

// context.WithDeadline 控制超时时间
func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(50*time.Millisecond))
	//尽管我们设置了5o毫秒超时，ctx会过期，cacel会关闭，但是还是关闭 cacel比较好
	defer cancel()
	select {
	// time.After(1 * time.Second) 表示创建了一个通道，这个通道会在 1 秒后发送一个信号
	case <-time.After(100 * time.Millisecond):
		println("overslept")
	case <-time.After(10 * time.Millisecond):
		println("before overslept")
	case <-ctx.Done():
		fmt.Println("已经出错", ctx.Err())
	}
}
