package main

import "fmt"

func main() {
	// slice 创建slice操作
	sli := make([]int, 2, 5)
	// slice 添加元素操作
	fmt.Printf("init slice: %v  add:%v \n", sli, &sli[0])
	// output: init slice: [0 0]  add:0xc00000a390
	sli2 := append(sli, 1, 2)
	fmt.Printf("sli2: %v  add:%v \n", sli2, &sli2[0])
	// output: sli2: [0 0 1 2]  add:0xc00000a390

	// 此时slice capacity不足，则创建一个新的slice
	sli3 := append(sli2, 3, 4, 5)
	fmt.Printf("sli3: %v  add:%v capacity=%v \n", sli3, &sli3[0], cap(sli3))
	// output: sli3: [0 0 1 2 3 4 5]  add:0xc000014320

	// slice 修改元素操作
	sli3[0] = 99
	fmt.Printf("sli3: %v  add:%v capacity=%v \n", sli3, &sli3[0], cap(sli3))

	// slice 遍历,查询操作
	for i := 0; i < len(sli3); i++ {
		fmt.Printf("idx:%v value:%v ", i, sli3[i])
	}
	// output: idx:0 value:99 idx:1 value:0 idx:2 value:1 idx:3 value:2 idx:4 value:3 idx:5 value:4 idx:6 value:5
	fmt.Println()
	for idx, val := range sli3 {
		fmt.Printf("idx:%v value:%v ", idx, val)
	}
	// output: idx:0 value:99 idx:1 value:0 idx:2 value:1 idx:3 value:2 idx:4 value:3 idx:5 value:4 idx:6 value:5

	// slice 删除元素操作(golang本身并没有直接提供删除slice元素的语法或函数，需要使用slice本身的特性来删除元素)
	//
}

/**
删除slice元素方法1 截取法(修改原切片)
*/
//func deleteSliceEl1(sourceSlice []int, delEl int) []int {
//	for i := 0; i < len(sourceSlice); i++ {
//		if sourceSlice[i] == delEl {
//			sourceSlice = append(sourceSlice[:i], sourceSlice[i+1]...)
//			i--
//		}
//	}
//	return sourceSlice
//}
