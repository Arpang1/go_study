package main

import "study/hanshu"

func main() {
	//n := 4
	////判断
	//panduan.Iftest(n)
	//panduan.Switchtest(n)
	////循环
	//xunhuan.ForTest(n)
	////xunhuan.ForTest2()
	//xunhuan.ForRangerTest(n)
	////跳转
	//tiaozhuan.GotoDemo()
	//tiaozhuan.BreakDemo()
	//tiaozhuan.ContinuDemo()
	////数组
	//array.Array()
	// 普通函数
	//fmt.Println(hanshu.Add(10, 20))
	//可变参数
	//fmt.Println(hanshu.KbADD(10, 20, 30))
	//对象数组
	//c := make([]hanshu.Person, 0, 100)
	//c = append(c, hanshu.Person{"apang", "man", 22})
	//c = append(c, hanshu.Person{"Lcs", "nan", 23})
	//c = append(c, hanshu.Person{"Mina", "woman", 20})
	//for _, x := range c {
	//	if x.Name == "Lcs" {
	//		fmt.Println(x)
	//		break
	//	}
	//}
	//高阶函数
	//作为参数
	//fmt.Println(hanshu.AddPro(20, 20, hanshu.Add))
	//作为返回值
	//fmt.Println(hanshu.Switchpro("+"))

	//匿名函数
	//hanshu.Nm()

	//闭包
	//f := hanshu.Adder()
	//fmt.Println(f(10))
	//fmt.Println(f(20))
	//
	//fmt.Println("Next f1")
	//f1 := hanshu.Adder()
	//fmt.Println(f1(30))
	//fmt.Println(f1(40))
	//闭包进阶1
	f2 := hanshu.Adder2(10)
	println(f2(5))
	println(f2(4))
	println(f2(7))
	//闭包进阶2
	//hz := hanshu.MakeSuffixFunc(".png")
	//println(hz("1.png"))
	//println(hz("2.png"))
	//println(hz("3.jpg"))
	//println(hz("4.png"))
	//println(hz("5.jpg"))
	//output
	//照片类型-正确
	//照片类型-正确
	//照片类型-错误
	//照片类型-正确
	//照片类型-错误

	//闭包进阶3
	//add1, sub1 := hanshu.Calc(10)
	//fmt.Println(add1(5))
	//fmt.Println(sub1(3))
	//fmt.Println(add1(4))
	//fmt.Println(sub1(2))
	//output
	//15
	//12
	//16
	//14
}
