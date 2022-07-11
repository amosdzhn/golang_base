package main

import "fmt"

/**
测试指针结构体类型初始化
*/
type cat struct {
	name string
	age  int
}

func main() {
	c := new(cat)
	*c = cat{
		name: "小苗",
		age:  20,
	}
	fmt.Println(c)
}
