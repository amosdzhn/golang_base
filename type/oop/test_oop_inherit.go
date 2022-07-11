package main

import "fmt"

/**
oop继承:go中通过结构体的组合实现面向对象的继承
*/

type animal struct {
	age int
}

func (a animal) stdPrintAge() {
	fmt.Println(a.age)
}

type person struct {
	animal
	name string
}

func (p person) stdPrintName() {
	fmt.Println(p.name)
}

type student struct {
	person
	school string
}

func (s student) stdPrintSchool() {
	fmt.Println(s.school)
}

func (s student) stdPrintName() {
	fmt.Println("李四")
}

func main() {
	stu := student{
		person: person{
			animal: animal{age: 20},
			name:   "张三",
		},
		school: "重庆交通大学",
	}

	fmt.Printf("打印属性:%v\n", stu)
	stu.stdPrintAge()
	stu.stdPrintName()
	stu.stdPrintSchool()
	stu.person.stdPrintName()
}
