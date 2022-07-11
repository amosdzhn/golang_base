package main

import "fmt"

/**
数组:相同数据类型的元素的集合
	数组一旦定义，长度不可修改,定义时必须使用常量来定义数组大小
	通过索引下标访问元素
	Go中数组类型形如: [Size]Type
*/
func main() {
	// 1.数组的定义与初始化: var varName = [size]type{}
	var names [3]string = [3]string{"张三", "李四"}
	fmt.Println(names)
	// 2.单独定义一个数组，未初始化，默认零值
	var names2 [3]int
	fmt.Println(names2)
	// 3.类型推到，长度堆到定义并初始化数组
	var names3 = [...]int{1, 2, 3}
	fmt.Printf("type:%T\n", names3) // [3]int
	// 4.访问数组元素
	var student = [...]string{"张三", "李四", "王五"}
	// 下标访问(存在下标越界)
	stu1 := student[0]
	fmt.Println(stu1)
	// 获取数组长度使用内置函数len()
	stuLen := len(student)
	fmt.Println(stuLen)
	// for 循环遍历数组
	for i := 0; i < len(student); i++ {
		fmt.Println(student[i])
	}
	// for range 循环遍历数组
	for i, v := range student {
		fmt.Printf("i:%v v:%v\n", i, v)
	}
	//
	fmt.Println(&student)
	fmt.Println(&student[0])
	fmt.Println(&student[1])
	fmt.Println(&student[2])
}
