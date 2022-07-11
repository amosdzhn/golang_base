package main

import "fmt"

func main() {
	fmt.Println("start")
	panic("错误信息")
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()
	fmt.Println("end")
}
