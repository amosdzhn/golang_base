package main

import "fmt"

/**
测试捕获系统panic
*/
func main() {
	fmt.Println("start main")
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	//fmt.Println(10 / 0)
	panic("aaa")
	fmt.Println("end main")
}
