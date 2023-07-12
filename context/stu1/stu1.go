package main

import (
	"fmt"
	"sync"

	"time"
)

// 如何优雅的实现结束子goroutine ?
//方法一： 使用全局变量
//方法二： 使用channel（管道）

var wg sync.WaitGroup

// 方法一
//var exit = false

// 初始的例子
func worker(e chan bool) {
Loop:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-e:
			fmt.Println("exit")
			break Loop
		default:
			continue
		}
		//if exit {
		//	break
		//}
	}
	// 如何接收外部命令实现退出
	wg.Done()
}

func main() {
	e := make(chan bool, 1)
	wg.Add(1)
	go worker(e)
	// 如何优雅的实现结束子goroutine ?
	//方法一
	//time.Sleep(time.Second * 3)
	//exit = true

	// 方法二
	time.Sleep(time.Second * 6)
	e <- true
	wg.Wait()
	fmt.Println("over")
}
