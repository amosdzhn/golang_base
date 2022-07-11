package main

import "fmt"

/**
匿名结构体
*/
func main() {
	// 定义一个匿名结构体
	var p struct {
		name string
		age  int
	}
	fmt.Printf("p1:%v type:%T\n", p, p)

	// 定义匿名结构体并赋值
	p2 := struct {
		name string
		age  int
	}{name: "张三", age: 20}
	fmt.Printf("p2:%v type:%T\n", p2, p2)
}
