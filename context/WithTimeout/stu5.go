package main

import (
	"context"
	"sync"
	"time"
)

var wg sync.WaitGroup

func work(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			println("work done")
			return
		default:
			println("work")
			time.Sleep(time.Millisecond * 10)
		}
	}
}

// WithTimeout返回WithDeadline(parent, time.Now().Add(timeout))
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	wg.Add(1)
	go work(ctx)
	wg.Wait()
}
