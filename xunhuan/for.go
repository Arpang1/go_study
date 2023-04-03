package xunhuan

import (
	"fmt"
	"strconv"
	"time"
)

func ForTest(n int) {
	for i := 0; i <= n; i++ {
		fmt.Println(i)
	}
}

//无限循环
func ForTest2() {
	str := "!"
	for {
		fmt.Println("无限循环" + str)
		str = str + "!"
		time.Sleep(1 * time.Second)
	}
}

// for range (键值循环)
func ForRangerTest(m int) {
	n := make([]string, m)
	for i := 0; i < m; i++ {
		n[i] = "str" + strconv.Itoa(i)
	}
	for _, i := range n {
		str := i
		fmt.Println("range:" + str)
	}
}
