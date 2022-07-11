package main

import "fmt"

/**
结构体的初始化:
创建结构体类型变量后若没有进行赋值，则为零值。
可以创建时对全部成员属性进行初始化或对一部分成员你属性初始化
*/
func main() {
	// 1.创建一个匿名结构体变量
	var p1 struct {
		name string
	}
	fmt.Println(p1) // {""}

	// 2.创建一个匿名结构体并赋值
	var p2 = struct {
		name string
	}{name: "张飞"}
	fmt.Println(p2) // {"张飞"}
}
