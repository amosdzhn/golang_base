package main

import "fmt"

/**
测试结构体嵌套
*/
type door struct {
	name  string
	color string
	size  int
}

type room struct {
	name string
	door door
	size int
}

func main() {
	d := door{
		name:  "门",
		color: "黑色",
		size:  20,
	}

	r := room{
		name: "教室",
		door: d,
		size: 20,
	}
	fmt.Println(r)

}
