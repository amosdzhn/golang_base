package main

import "fmt"

func main() {
	fmt.Println("start main()")
	panic("错误信息")
	errMessage := recover()
	fmt.Println(errMessage)
	fmt.Println("end main()")

}
