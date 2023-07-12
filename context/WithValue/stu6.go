package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func work(ctx context.Context) {
	key := "lcs"
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		println("invalid trace code")
	}
LOOP:
	for {
		println(traceCode, "is running")
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println(traceCode, "end")
	wg.Done()
}

// WithValue返回父节点的副本，其中与key关联的值为val。
func main() {
	ctx1, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	ctx := context.WithValue(ctx1, "lcs", "xxx")
	wg.Add(1)
	go work(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
