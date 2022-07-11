package main

import "fmt"

/**
测试指针结构体:使用new()创建一个指针结构体类型变量
*/
type dog struct {
	name string
	age  int
}

func main() {
	var d *dog = new(dog)
	fmt.Printf("值:%v 类型:%T 地址:%p\n", d, d, d)
	// 由于d是一个指针结构体类型,在取值时可以省略*
	(*d).name = "张飞"
	d.age = 20
	fmt.Printf("值:%v 类型:%T 地址:%p\n", d, d, d)
}
