package main

import "fmt"

/**
切片类型：
	数组固定长度，带来了一些限制
	切片可以理解为可变长的数组，底层使用数组实现，存在容量增加功能
	切片是一个存放相同类型元素的可变长容器
*/
func main() {
	// 1.声明一个切片
	testSliceDefinition()
	// 2.测试切片长度和容量
	testSliceLenCap()
	// 3.测试切片的遍历
	testSliceFor()
	// 4.切片的CRUD操作
	testSliceCRUD()
	// 5.切片拷贝
	testSliceCopy()

}

func testSliceDefinition() {
	// 定义切片
	fmt.Println("定义切片:")
	var sl1 []int
	// 定义切片并初始化
	var sl4 []int = []int{1, 2, 3}
	// 声明一个整形元素类型切片，且初始长度为2，容量为3
	var sl2 []int = make([]int, 2, 3)
	sl3 := make([]string, 1, 3)
	// 从数组类型到切片类型
	arr := [...]int{1, 2, 3, 4, 5}
	sl5 := arr[:]
	fmt.Printf("type sl1:%T sl2:%T sl3:%T sli4:%T sl5:%T\n", sl1, sl2, sl3, sl4, sl5)
}

func testSliceLenCap() {
	// 测试切片的长度和容量
	fmt.Println("测试切片长度和容量:")
	slice1 := make([]int, 5, 10)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)
}

func testSliceFor() {
	// 测试切片的遍历(查询)
	names := []string{"张三", "李四"}
	for i := 0; i < len(names); i++ {
		println(names[i])
	}

	for i, name := range names {
		fmt.Println(i, name)
	}
}

func testSliceCRUD() {
	// 1.创建
	students := make([]string, 2, 5)
	fmt.Printf("创建后:%v\n", students)
	// 2.添加
	students2 := append(students, "张飞")
	fmt.Printf("添加后:%v\n", students2)
	fmt.Printf("添加前和添加后的两个切片的地址比较:%v,%v\n", &students, &students2)
	// 一般添加之后将新的切片赋值给旧的切片变量
	students2 = append(students2, "曹操")
	fmt.Printf("再次添加后:%v\n", students2)
	// 3.删除 以删除"曹操"为例子  曹操索引是3
	// 将student [0,3) 添加 student [4,...) 那么就将将曹操删除了
	students2 = append(students2[:3], students2[4:]...)
	fmt.Printf("删除之后:%v\n", students2)
	// 4.修改
	students2[2] = "刘备"
	fmt.Printf("修改之后:%v\n", students2)
}

func testSliceCopy() {
	// 测试切片拷贝
	var numbers = []int{1, 2, 3}
	numbers2 := numbers
	// 当修改切片元素时，原切片元素也被修改
	// 切片的赋值是赋值原切片的地址
	numbers2[0] = 4
	fmt.Printf("原切片值:%v 地址:%v  第一个元素地址:%v\n", numbers, &numbers, &numbers[0])
	fmt.Printf("现切片值:%v 地址:%v  第一个元素地址:%v\n", numbers2, &numbers2, &numbers2[0])

	// 使用copy() 内置函数是深拷贝
	var number3 = []int{1, 2, 3, 4}
	var number4 []int = make([]int, 4)
	result := copy(number4, number3)
	fmt.Printf("copy返回值:%v\n", result) // 返回值是copy完成的个数
	number4[0] = 100
	fmt.Printf("原切片:%v 地址:%v 第一个元素地址:%v\n", number3, &number3, &number3[0])
	fmt.Printf("现切片:%v 地址:%v 第一个元素地址:%v\n", number4, &number4, &number4[0])
}
