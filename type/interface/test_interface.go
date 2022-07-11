package main

import "fmt"

/**
接口是一堆未实现方法的集合。
在Go中不需要结构体指定实现某个接口，只要
实现了一个接口中所有的方法就称为实现了这个接口
*/

// 定义一个usb接口
type usb interface {
	read()
	write()
}

// 定义一个电脑结构体
type computer struct {
	name  string
	price float32
}

func (c *computer) read() {
	fmt.Println("read()...")
}

func (c *computer) write() {
	fmt.Println("write()...")
}
func main() {

}
