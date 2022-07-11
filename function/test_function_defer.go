package main

import "fmt"

/**
defer 后面的语句，可以延迟到函数return之前执行
		若单个函数中存在多个defer，按照栈的规则执行(先进后出)
*/

func sayHello() int {
	defer fmt.Println("step1..")
	defer fmt.Println("step2..")
	defer func() {
		fmt.Println("匿名函数中的 step3...")
	}()
	return 1
}
func main() {
	hello := sayHello()
	fmt.Println(hello)
	/**
	输出顺序:
		匿名函数中的 step3...
		step2..
		step1..
		1
	*/
}
