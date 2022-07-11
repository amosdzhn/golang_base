package main

import "fmt"

/**
测试 for range循环
	可以对数组、切片、字符串、map、chan进行遍历
	数组、切片、字符串返回索引和值
	map 返回 key 和 value
	channel 返回通道中的值
*/
func main() {
	var arr = [...]string{"张飞", "关羽", "黄忠", "赵云", "马超"}
	fmt.Println("==对数组进行 for range遍历")
	fmt.Printf("arr type:%T\n", arr) // [5]int
	for i, v := range arr {
		fmt.Printf("index:%v value:%v\n", i, v)
	}

	var content = "空山新雨后，天气晚来秋"
	fmt.Println("==对字符串进行 for range遍历")
	fmt.Printf("content type:%T\n", content) // string
	for i, v := range content {
		fmt.Printf("index:%v value:%v string(value):%v\n", i, v, string(v))
	}

	var slice = []string{"重庆", "四川", "上海"}
	fmt.Println("==对切片进行 for range遍历")
	fmt.Printf("slice type:%T\n", slice)
	for i, v := range slice {
		fmt.Printf("index:%v value:%v\n", i, v)
	}

	var students = make(map[string]int)
	students["张三"] = 19
	students["李四"] = 20
	students["王五"] = 22
	fmt.Printf("students type:%T\n", students)
	fmt.Println("==对map进行 for range遍历")
	for k, v := range students {
		fmt.Printf("k:%v v:%v\n", k, v)
	}
}
