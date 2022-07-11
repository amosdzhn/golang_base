package main

import "fmt"

// 测试chan类型的创建
func main() {
	// 声明一个存放int数据的双向通道
	var a chan int
	fmt.Println(nil == a)
	close(a)
	fmt.Println(a) // nil
}
