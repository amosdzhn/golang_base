package main

import "fmt"

/**
在Go 或者 JavaScript等语言中，函数是一等公民
在Java等存面向对象语言中，类是一等公民

(函数名，参数列表，返回值类型)构成函数的签名

函数的特性：
1.Go中有三种函数：普通函数，匿名函数，方法(作用在结构体上)
2.Go语言中不允许函数重名(重载)
3.Go语言中函数不能嵌套函数，但是可以嵌套匿名函数
4.Go语言中函数可以赋值给一个变量，使得该变量也成为函数
5.函数可以作为参数，传递给另一个函数
6.函数的返回值可以是一个函数
7.函数调用的时候，如有有参数，先拷贝参数的副本，再将副本传递给函数
8.函数参数可以没有名称
*/
func main() {
	// 测试求和函数
	fmt.Println(sum(1, 1))
	fmt.Println(compare(1, 2))
}

// 定义一个求和函数
func sum(a, b int) (res int) {
	return a + b
}

// 定义一个比较函数
func compare(a, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}
