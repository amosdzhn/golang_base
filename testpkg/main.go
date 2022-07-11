package main

import (
	"fmt"
	"testpkg/p1"
)

/**
在一个mod中，跑
*/

func init() {
	fmt.Println("main init" +
		"")
}
func main() {
	fmt.Println("项目启动")
	p1.SayHello()
	p1.SayBye()
}
