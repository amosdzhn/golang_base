package main

import "fmt"

/**
测试进制数
*/
func main() {
	// 定义各种进制的数
	var n10 int = 20
	var n8 int = 0100
	var n16 int = 0x300

	// 按照指定进制输入
	fmt.Printf("%d \n", n10)
	fmt.Printf("%b \n", n10)

	fmt.Printf("%o \n", n8)
	fmt.Printf("%x \n", n16)

	var num = 10
	fmt.Printf("%o \n", num)

	var num1 = 16
	fmt.Printf("%x \n", num1)
}
