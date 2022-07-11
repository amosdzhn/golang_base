package main

import (
	"fmt"
)

/**
1.函数类型(函数类型变量)
2.匿名函数
*/

// 使用type声明一个函数类型
type name func(a, b int) int

// 函数类型变量作为形参
func op(a, b int, op func(a, b int) int) int {
	a *= 10
	b *= 10
	return op(a, b)
}

func main() {
	// 1.匿名函数,不调用
	var f1 = func(a int) int {
		return a + 1
	}
	fmt.Println(f1(10)) // 调用匿名函数

	// 2.匿名函数，立刻调用
	r := func(a, b int) int {
		return a + b
	}(3, 5)
	fmt.Printf("匿名函数调用结果:%v\n", r)

	// 3.调用函数类型变量作为参数的函数
	cop := func(a, b int) int {
		return a + b
	}
	result := op(1, 2, cop)
	fmt.Printf("调用op函数返回的结果:%v\n", result)

}
