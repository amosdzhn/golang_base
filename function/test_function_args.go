package main

import "fmt"

/**
了解Go函数参数列表
*/
func main() {
	testMethodArgs1(1)

	testMethodArgs2(100, "不该")

	testMethodArgs4(1, 2, 3, 4, 5)
}

// 单个类型的参数列表
func testMethodArgs1(a int) {
	fmt.Println("单个类型参数:")
	fmt.Println(a)
	return
}

// 多个类型的参数
func testMethodArgs2(a int, b string) {
	fmt.Println("多个类型参数:")
	fmt.Println(a, b)
	return
}

// 相同类型的参数简写
func testMethodArgs3(a, b int) {
	fmt.Println("相同类型参数简写:")
	fmt.Println(a, b)
	return
}

// 可变长参数
func testMethodArgs4(args ...int) {
	fmt.Println("可变长参数:")
	for i, v := range args {
		fmt.Println(i, v)
	}
}
