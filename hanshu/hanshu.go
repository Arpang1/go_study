package hanshu

import (
	"errors"
	"fmt"
	"strings"
)

//函数的第一个字母大写

//函数定义
//函数声明
//func 函数名(参数)(返回值){
//	函数体
//}

//一个普通的函数
func Add(x, y int) int {
	return x + y
}
func Sub(x, y int) int {
	return x - y
}

//可变参数
//可变参数变得是参数的数量而不是参数的类型
func KbADD(x ...int) int {
	sum := 0
	for _, i := range x {
		sum += i
	}
	return sum
}

//对象数组
type Person struct {
	Name string
	Sex  string
	Age  int
}

//高阶函数：函数本身作为函数的参数或返回值
//作为参数
func AddPro(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//作为返回值
func Switchpro(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return Add, nil
	case "-":
		return Sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//匿名函数
//匿名函数声明
//func(参数)(返回值){
//	函数体
//}
//因为匿名函数没有函数名，所以匿名函数只能在某个变量或者函数中
//匿名函数多用于实现回调函数和闭包.
func Nm() {
	//匿名函数寄存在变量中，此时add相当于变量名，但只在Nm函数中可以调用。
	add := func(x, y int) int {
		return x + y
	}
	sum := add(20, 30)
	fmt.Println(sum)
	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(50, 20)
}

//闭包
//闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
func Adder() func(int) int {
	var x int
	fmt.Println("x=:", x)
	return func(y int) int {
		fmt.Println("y=:", y)
		x += y
		return x
	}
}

//闭包进阶1
func Adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

//闭包进阶2
func MakeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		//strings.HasSuffix是strings包中的函数，用于检查字符串是否以指定的后缀结尾。 其传入参数为(string,string)两个字符串,返回值为bool一个布尔类型.
		//假如我传入x和y(strings.HasSuffix(x, y)),如果x的后缀为y，那么返回true；否则为false
		if !strings.HasSuffix(name, suffix) {
			return "照片类型-错误"
		}
		return "照片类型-正确"
	}
}

//闭包进阶3
func Calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}
