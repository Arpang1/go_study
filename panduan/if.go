package panduan

import "fmt"

func Iftest(n int) {
	if n >= 4 {
		fmt.Println("优秀")
	} else if n >= 2 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}
}
