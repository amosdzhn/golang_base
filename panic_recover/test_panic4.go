package main

import "fmt"

func m1() {
	fmt.Println("start m1")
	defer panic("m1 error message")
	fmt.Println("end m1")
}
func main() {
	fmt.Println("start main")
	m1()
	fmt.Println("end main")
}
