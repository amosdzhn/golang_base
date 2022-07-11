package main

import "fmt"

func main() {
	fmt.Println("start main()")
	defer panic("错误信息")
	fmt.Println("end main()")

}
