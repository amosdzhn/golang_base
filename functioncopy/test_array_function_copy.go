package main

import "fmt"

/**
数组是值拷贝
*/
func testArrayCopy(a [3]int) {
	a[0] = 1000
	fmt.Printf("testArrayCopy中的数组地址:%v\n", &a)
	fmt.Printf("函数调用中:%v\n", a)
}
func main() {
	var ages = [3]int{1, 2, 3}
	fmt.Printf("main方法中的数组地址:%v\n", &ages)
	fmt.Printf("函数调用之前:%v\n", ages)
	testArrayCopy(ages)
	fmt.Printf("函数调用之后:%v\n", ages)
}
