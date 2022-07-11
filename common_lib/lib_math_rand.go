package main

import (
	"fmt"
	"math/rand"
)

/**
math.rand 产生随机数
*/
func main() {
	//默认seed是1
	rand.Seed(1)
	// 产生一个随机数
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int())
	}
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int31())
	}
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int63())
	}
	for i := 0; i < 5; i++ {
		fmt.Println(rand.Intn(100))
	}

	//设置种子
	rand.Seed(11)
	for i := 0; i < 5; i++ {
		// 返回一个指定范围内的随机数
		fmt.Println(rand.Intn(10))
	}
}
