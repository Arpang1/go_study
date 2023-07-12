package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func work(ctx context.Context) {
	go work2(ctx)
LOOP:
	for {
		fmt.Println("work")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			break LOOP
		default:
			continue
		}
	}

	wg2.Done()
}

func work2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println("exit1")
			break LOOP
		default:
			continue
		}
	}

	wg2.Done()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var ctx1 context.Context
	ctx2 := context.TODO()
	ctx3 := context.Background()
	//---with---
	ctx4, err := context.WithTimeout(ctx1, time.Second*3)
	ctx5, err := context.WithDeadline(ctx1, time.Now().Add(time.Second*3))
	ctx6 := context.WithValue(ctx1, "key", "value")
	ctx7, err := context.WithCancel(ctx1)
	fmt.Println(err)
	fmt.Println(ctx2, ctx3, ctx4, ctx5, ctx6, ctx7)
	wg2.Add(1)
	go work(ctx)
	time.Sleep(time.Second * 3)
	cancel()
	wg2.Wait()
	defer ctx.Done()

}
