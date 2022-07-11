package main

import "fmt"

/**
了解Go函数返回值
*/
func main() {

	r1 := testMethodReturn1()
	fmt.Println(r1) // 20

	r21, r22 := testMethodReturn2()
	fmt.Println(r21, r22)

	r31, r32 := testMethodReturn3()
	fmt.Println(r31, r32)

	r41, r42 := testMethodReturn4()
	fmt.Println(r41, r42)

	r5 := testMethodReturn5()
	r55 := r5()
	fmt.Println(r55)
}

// 无返回值
func testMethodReturn0() {
	return
}

// 单返回值
func testMethodReturn1() int {
	return 20
}

// 多返回值
func testMethodReturn2() (int, int) {
	return 1, 1
}

// 多返回值变量
func testMethodReturn3() (a int, b int) {
	a = 1
	b = 1
	return
}

// 多返回值变量简写
func testMethodReturn4() (a, b int) {
	a = 1
	b = 2
	return a, b
}

// 返回一个函数
func testMethodReturn5() func() int {
	return func() int {
		return 100
	}
}
