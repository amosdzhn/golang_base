package main

import "fmt"

/**
函数闭包(closure)
可以理解为：闭包=引用函数+返回函数
现在的理解是：通常只能是下层函数中传递变量到上层函数中，
	闭包可以将上层函数的变量传递到下层函数
*/

//
func count() func() int {
	var count int
	fmt.Println("count execute..")
	return func() int {
		count += 1
		return count
	}
}
func main() {
	f1 := count()
	fmt.Printf("第一次调用count(),%v\n", f1()) // 1
	f2 := count()
	fmt.Printf("第二次调用count(),%v\n", f2()) // 1
	// 因为上面两次调用count()函数，从新给count赋值了,每次返回的函数类型返回值都是不同的

	f3 := count()
	fmt.Printf("第一次调用count()且第一次调用f3(),%v\n", f3()) // 1
	fmt.Printf("第一次调用count()且第二次调用f3(),%v\n", f3()) // 2
}
