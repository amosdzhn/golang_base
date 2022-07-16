package main

import "fmt"

/**
测试Go函数的传递类型是值传递还是引用传递(地址)
*/

// int 默认值传递
func testInt(a int) {
	a = 200
}

func testIntPtr(a *int) {
	*a = 200
}

// string 默认值传递
func testString(a string) {
	a = "im ok"

}

func testStringPtr(a *string) {
	*a = "world"
}

// slice 默认地址传递
func testSlice(numbers []int) {
	for i := 0; i < len(numbers); i++ {
		numbers[i] += 1
	}
}

type person struct {
	name string
	age  int
}

// person是值传递
func testStruct(p person) {
	fmt.Printf("p address=%v \n", &p)
	p.name = "ttt"
	p.age = 111
}

func main() {
	n := 20
	testInt(n)
	fmt.Printf("传递int类型变量修改之后的结果是=%v \n", n) // output: 20

	testIntPtr(&n)
	fmt.Printf("传递*int类型变量修改之后的结果是=%v \n", n) // output: 200

	str := "hello"
	testString(str)
	fmt.Printf("传递string类型变量修改后的结果是=%v \n", str) // output: hello

	testStringPtr(&str)
	fmt.Printf("传递*string类型变量修改之后的结果是=%v \n", str) // output: world

	numbers := []int{1, 2, 3, 4, 5}
	testSlice(numbers)
	fmt.Printf("传递slice类型变量修改之后的结果是=%v \n", numbers) // output: [2 3 4 5 6]

	p := person{
		name: "zs",
		age:  0,
	}
	testStruct(p)
	fmt.Printf("传递struct类型变量修改之后的结果是=%v  add=%v \n", p, &p) // output: {zs 0}
}
