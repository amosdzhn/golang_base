package main

import "fmt"

/**
interface nested
*/

// 走
type walker interface {
	walk()
}

// 跳
type jumper interface {
	jump()
}

// 跑
type runner interface {
	run()
}

type animal interface {
	walker
	jumper
	runner
	eat()
	sayHello()
}

type person struct {
	name string
	age  int
}

func (p *person) walk() {
	fmt.Println(*p, "walk")
}

func (p *person) jump() {
	fmt.Println(*p, "jump")
}

func (p *person) run() {
	fmt.Println(*p, "run")
}

func (p *person) eat() {
	fmt.Println(*p, "eat")
}
func (p *person) sayHello() {
	fmt.Println(*p, "sayHello")
}

func main() {
	var p = person{
		name: "张飞",
		age:  21,
	}
	fmt.Println("===结构体类型变量调用")
	p.walk()
	p.jump()
	p.run()
	p.eat()
	p.sayHello()

	fmt.Println("===animal类型变量调用")
	var anm animal = &p
	anm.walk()
	anm.jump()
	anm.run()
	anm.eat()
	anm.sayHello()

	fmt.Println("===walker类型变量调用")
	var walk walker = &p
	walk.walk()

	fmt.Println("===runner类型变量调用")
	var run runner = &p
	run.run()

	fmt.Println("===jumper类型变量调用")
	var ju jumper = &p
	ju.jump()

}
