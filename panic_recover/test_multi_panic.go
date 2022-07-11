package main

import "fmt"

func method1() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	fmt.Println("start method1()")
	method2()
	panic("error message in method1()")
	fmt.Println("end method1()")
}
func method2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("start method2")
	panic("error message in method2()")
	fmt.Println("end method2")
}
func main() {
	fmt.Println("start main()")
	method1()
	fmt.Println("end main()")
}
