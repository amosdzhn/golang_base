package main

import "fmt"

type desk struct {
	high int
}

func (d desk) m1() {
	fmt.Printf("(d desk) %p\n", &d)
}

func (d *desk) m2() {
	fmt.Printf("(d *desk) %p\n", d)
}

/**
其实receiver 是在调用方法隐式传入的第一个参数，在java中是this
*/
func main() {
	// 还可以这样调用方法，显示的传递
	d := desk{high: 20}
	fmt.Printf("main中的d变量地址:%p\n", &d) // 0xc0000120a8
	desk.m1(d)                         // 0xc0000120a8
	(*desk).m2(&d)                     // 0xc0000120a8
}
