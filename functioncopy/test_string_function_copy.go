package main

import "fmt"

/**
函数传递string类型是值拷贝
*/
func testStringCopy(s string) {
	s = s + "!"
	fmt.Printf("函数执行中:%v\n", s)
}
func main() {
	name := "东方不败"
	fmt.Printf("函数调用之前：%v\n", name)
	testStringCopy(name)
	fmt.Printf("函数调用之后：%v\n", name)
}
