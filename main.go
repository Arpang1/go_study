package main

import (
	"study/panduan"
	"study/tiaozhuan"
	"study/xunhuan"
)

func main() {
	n := 4
	panduan.Iftest(n)
	panduan.Switchtest(n)
	xunhuan.ForTest(n)
	//xunhuan.ForTest2()
	xunhuan.ForRangerTest(n)
	tiaozhuan.GotoDemo()
	tiaozhuan.BreakDemo()
	tiaozhuan.ContinuDemo()
}
