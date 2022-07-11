package main

import "fmt"

/**
测试在多个嵌套调用函数时的panic与defer recover的使用
*/
func method1() {
	//defer func() {
	//	p := recover()
	//	if nil != p {
	//		fmt.Println("捕获了异常")
	//	}
	//}()
	fmt.Println("start m1")
	method2()
	fmt.Println("end m1")
}
func method2() {
	//defer func() {
	//	p := recover()
	//	if nil != p {
	//		fmt.Println("捕获了异常")
	//	}
	//}()
	fmt.Println("start m2")
	method3()
	fmt.Println("end m2")
}
func method3() {
	fmt.Println("start m3")
	panic("错误信息对象")
	fmt.Println("end m3")
}
func main() {
	fmt.Println("start main")
	method1()
	fmt.Println("end main")
}
