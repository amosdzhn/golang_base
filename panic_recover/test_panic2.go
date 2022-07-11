package main

import "fmt"

func main() {
	fmt.Println("start main()")
	errMessage := "错误信息"
	panic(errMessage)
	fmt.Println("end main()")
}
