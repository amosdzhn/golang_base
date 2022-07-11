package main

import "fmt"

func m1() {
	fmt.Println("start execute m1()")
	fmt.Println("end execute m1()")
}

func main() {
	fmt.Println("start execute main()")
	defer m1()
	defer func() {
		fmt.Println("start f2")
		fmt.Println("end f2")
	}()
	defer func() {
		fmt.Println("start f3")
		fmt.Println("end f3")
	}()
	fmt.Println("end execute main()")
}
