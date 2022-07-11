package main

import "fmt"

/**
Go中的面向对象是将函数绑定到结构体进而实现的方法
*/

type computer struct {
	name string
}

func (c computer) work() {
	fmt.Println("work()...")
}

func (c computer) sleep() {
	fmt.Println("sleep()...")
}
func main() {
	com := computer{
		name: "华硕",
	}
	com.work()
	com.sleep()
	fmt.Println(com.name)
}
