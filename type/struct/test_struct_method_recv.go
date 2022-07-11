package main

import "fmt"

/**
测试method receiver的类型区别:同普通函数类似理解，当直接传递结构体类型变量时是值传递，
而传递指针结构体类型时是引用传递
*/

type book struct {
	name string
}

// 结构体类型作为receiver 在使用结构体类型变量时是通过值传递(拷贝)
// 在当前方法中对结构体field的修改不会生效
func (b book) update() {
	b.name = "b book"
}

// 指针结构体类型作为receiver 在修改后会生效,是地址传递
func (b *book) update2() {
	b.name = "b *book"
}
func main() {
	aBook := book{
		name: "数据结构",
	}
	fmt.Printf("update()执行前:%v\n", aBook)
	aBook.update()
	fmt.Printf("update()执行后:%v\n", aBook)

	fmt.Printf("update2()执行前:%v\n", aBook)
	aBook.update2()
	fmt.Printf("update2()执行后:%v\n", aBook)
}
