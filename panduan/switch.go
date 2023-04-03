package panduan

import "fmt"

func Switchtest(n int) {
	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

//一个分支可以有多个值，多个case值中间使用英文逗号分隔

func testSwitch2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：
func switchtest3(n int) {
	switch {
	case n < 25:
		fmt.Println("好好学习吧")
	case n > 25 && n < 35:
		fmt.Println("好好工作吧")
	case n > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}
