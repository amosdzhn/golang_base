package main

import "fmt"

/**
数组指针:数组的每个元素是指针类型
*/
func main() {
	// 1.数组指针
	var aPtr [3]*int
	var array = [3]int{1, 2, 3}
	for i := 0; i < len(array); i++ {
		aPtr[i] = &array[i]
	}

	// 遍历
	for _, i := range aPtr {
		fmt.Println(*i)
	}
}
