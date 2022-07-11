package main

import "fmt"

/**
init 函数可用于包的环境初始化操作。
每个源文件中可以有多个init函数
每个包中可以有多个init函数,执行顺序不确定
不同包之间的init函数，按照导包的顺序执行

go语言执行顺序是:变量初始化 -> init() -> main()
*/

// 1.声明一些变量
var version = createVersion()

func createVersion() int {
	fmt.Println("createVersion() execute..")
	return 10
}
func init() {
	fmt.Println("init() 1 execute..")
}
func init() {
	fmt.Println("init() 2 execute..")
}
func init() {
	fmt.Println("init() 3 execute..")
}
func main() {
	fmt.Println("main() execute..")
}
