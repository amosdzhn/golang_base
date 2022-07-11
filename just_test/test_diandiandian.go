package main

import "fmt"

/**
golang 中的 ... 是一个语法糖
作用1：用于函数存在多个不定长参数(不定长参数类型参数，可以接收自定义实参，或自定义的slice)
作用2: 将slice作为不定长参数的实际参数进行传递
*/
func method(args ...int) {
	fmt.Printf("args=%v type(args)=%T \n", args, args)
	// output: %T = []int
}

func main() {
	//
	method(1, 2, 3, 4, 5)
	ages := []int{1, 2, 3, 4, 5}
	method(ages...)

	ints := make([]int, 5, 10)
	ints = append(ints, ages...)
	fmt.Println(ints)

}
