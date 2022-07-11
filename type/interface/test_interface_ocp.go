package main

import "fmt"

/**
OCP设计原则：对扩展开放，对修改关闭
*/

type swimmer interface {
	swim()
}

type pig struct {
	name string
}

func (p *pig) swim() {
	fmt.Printf("我是%v,我在游泳\n", p.name)
}

type fish struct {
	name string
}

func (f *fish) swim() {
	fmt.Printf("我是%v,我在游泳\n", f.name)
}

// 该方法的参数是一个接口类型，只要实现了该接口的结构体类型变量都可以作为实际参数进行传递

func whoIsSwim(sw swimmer) {
	sw.swim()
}
func main() {
	p := pig{name: "猪猪侠"}
	f := fish{name: "小鲤鱼"}
	whoIsSwim(&p)
	whoIsSwim(&f)
}
