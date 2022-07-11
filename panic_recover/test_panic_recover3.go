package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer func()")
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	panic("错误信息")
	fmt.Println("end")

}
