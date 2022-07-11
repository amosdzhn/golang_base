package main

import "fmt"

/**
slice是引用传递
*/
func testSliceCopy(a []int) {
	if len(a) > 1 {
		a[0] = 200
	}
	fmt.Printf("函数调用中:%v\n", a)
}
func main() {
	slice := []int{1, 2, 3}
	fmt.Printf("函数调用之前:%v\n", slice)
	testSliceCopy(slice)
	fmt.Printf("函数调用之后:%v\n", slice)
}
