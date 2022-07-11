package main

import "fmt"

/**
结构体类型是值传递
*/
type person struct {
	name string
	age  int
}

func testStructCopy(p person) {
	p.age = 100
	fmt.Printf("执行中的person:%v\n", p)
}

func testStructPointerCopy(p *person) {
	p.age = 200
	fmt.Printf("执行中的person:%v\n", p)
}
func main() {
	p := person{
		name: "张三",
		age:  20,
	}

	fmt.Printf("函数调用之前的person:%v\n", p)
	testStructCopy(p)
	fmt.Printf("函数调用之后的person:%v\n", p)

	p2 := new(person)
	fmt.Printf("执行指针传递之前:%v\n", p2)
	testStructPointerCopy(p2)
	fmt.Printf("执行指针传递之前:%v\n", p2)

}
