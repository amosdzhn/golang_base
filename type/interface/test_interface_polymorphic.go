package main

import "fmt"

/**
测试接口多态
一个结构体可以实现多个接口
一个接口可以被多个结构体实现
*/
type eater interface {
	eat()
}

type cat struct {
	name string
	age  int
}

func (c *cat) eat() {
	fmt.Println(c)
}

type dog struct {
	name string
}

func (d *dog) eat() {
	fmt.Println(d)
}

func main() {
	d := new(dog)
	c := new(cat)
	d.eat()
	c.eat()
	// 创建接口类型变量，被赋为实现了该接口的结构体变量
	var eater eater
	fmt.Println(eater) // nil
	eater = c
	fmt.Println(eater) //

}
