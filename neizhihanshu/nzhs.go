package neizhihanshu

import (
	"fmt"
)

// 内置函数
// defer语句：
func Defer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)
}

// Panic语句
func Panic() {
	fmt.Println("A")
	panic("B")
	fmt.Println("C")
}

//recover语法

func Recover() {
	fmt.Println("A")
	panic("B")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("C")
}
