package main

import "fmt"

func main() {
	f := 1.22641
	fmt.Println(f)
	// 保留两位小数输出,四舍五入
	fmt.Printf("f=%.2f", f)
}
