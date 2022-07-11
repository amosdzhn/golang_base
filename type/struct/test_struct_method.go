package main

import "fmt"

/**
测试结构体方法

方法的receiver type 并不一定要是结构体类型。type定义的别名都可以
struct结合它的方法就等价于面向对象中的类，只不过struct可以和它的方法分开，并非要属于同一个源文件
方法就是函数，Go中没有方法的重载，所有的方法名必须唯一
方法有两种接收类型:(T type) 或者 (T *type) 它们之间有区别
如果receiver是一个指针类型，则会自动解除引用
方法和类型是分开的，意味着实例的行为(method)和数据存储(field)是分开的,但是它们通过receiver建立联系
*/

type sky struct {
	name  string
	color string
}

// 将方法绑定到结构体中，结构体为接收者receiver
func (recv sky) fmtName() {
	fmt.Println(recv.name)
}
func main() {
	s := new(sky)
	s.name = "张三"
	s.fmtName()
}
