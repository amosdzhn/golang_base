package main

import "fmt"

/**
1.定义一个结构体类型,类型名为person
2.创建结构体类型变量再赋值，先创建并初始化
*/
type person struct {
	name  string
	age   int
	email string
}

type Student struct {
	name, email string
	age, price  int
}

func main() {
	//2.声明变量再赋值
	var p1 person
	fmt.Println(p1) // {"" 0 ""}
	p1 = person{
		name: "张昭",
		age:  20,
	}
	fmt.Println(p1) // {"张昭",20,""}

	//3.声明结构体类型变量并赋值
	var p2 = person{
		name:  "张良",
		age:   19,
		email: "没有",
	}
	fmt.Println(p2)

	p3 := new(person)
	p3.name = "你好"
	fmt.Println(*p3)
}
