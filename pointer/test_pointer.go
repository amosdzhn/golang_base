package main

import "fmt"

/**
指针:指针也是一种数据类型，指针类型的变量的值是一个地址
	只声明变量，未赋值时为nil
	指针类型变量，对应存储的是某一类型变量的地址
	Golang中的指针不能进行运算
*/
func main() {
	// 1.声明一个整数指针类型变量
	var iptr *int
	fmt.Println(iptr == nil)      // true
	fmt.Println(iptr)             // nil
	fmt.Printf("type:%T\n", iptr) // *int

	// 2.给指针变量赋值
	var value = 12
	iptr = &value
	fmt.Printf("%v %v\n", iptr, *iptr) // 0xc0000120f0 12

}
