package array

import "fmt"

func Array() {

	//一维数组
	var a = [5]int{12, 34, 56, 78, 90}
	//也可以这样写
	b := [5]int{21, 43, 65}
	fmt.Println("a的长度和容量:", len(a), cap(a))
	fmt.Println("b的长度和容量:", len(b), cap(b))
	fmt.Println("a数组为：", a, "b数组为:", b)
	//二维数组
	var c = [4][5]int{{1, 2, 3}, {4, 5, 6}}
	//也可以这样使用
	d := [4][5]int{[5]int{1, 2, 3}, [5]int{4, 5, 6}}
	fmt.Println("d的行的长度和容量:", len(d), cap(d))
	fmt.Println("d的某行中列的长度和容量", len(d[1]), cap(d[1]))
	fmt.Println("c数组为:", c, "d数组为:", d)
}
