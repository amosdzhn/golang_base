package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func()")
	}()
	panic("错误信息")
	fmt.Println("end main")
}
