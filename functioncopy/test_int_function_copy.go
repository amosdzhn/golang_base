package main

import "fmt"

/**
测试函数类型是int:传递int类型参数时是值拷贝
*/
func testIntCopy(a int) {
	a = a + 1
	fmt.Printf("调用函数中:%v\n", a)
}
func main() {
	number := 100
	fmt.Printf("调用函数之前:%v\n", number)
	testIntCopy(number)
	fmt.Printf("调用函数之后:%v\n", number)
}
