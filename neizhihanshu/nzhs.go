package neizhihanshu

import (
	"fmt"
	"strconv"
)

//内置函数
//defer语句：
func Defer() {
	fmt.Println(1)
	fmt.Println(2)
	c := 1 + 1 + 1 + 1 + 2
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	fmt.Println(6)
	defer fmt.Println("c:" + strconv.Itoa(c))
	a := []int{1, 2, 3, 4, 5, 6}

	c = 20

	defer fmt.Println(a)
}
