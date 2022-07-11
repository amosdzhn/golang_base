package main

import (
	"fmt"
	"strings"
)

/**
go 中的字符串
*/
func main() {
	// 字符串定义
	var str1 string = "独孤求败"
	var str2 = "东方不败"
	str3 := "独孤九剑"
	fmt.Printf("str1:%v \nstr2:%v \nstr3:%v \n", str1, str2, str3)

	// 字符串的拼接
	str4 := str2 + str3
	str5 := strings.Join([]string{"aa", "bb"}, ",")
	fmt.Printf("字符串拼接结果:\nstr4:%v\nstr5:%v\n", str4, str5)

	// strings包下的常用函数

	// string的切片操作
	var name = "东方不败独孤求败"
	firstName := name[:4]
	fmt.Printf("firstName:%s\n", firstName) // 出现乱码
	fmt.Printf("中间不败:%s\n", name[2:4])      // 出现乱码
	// 分析乱码原因
	fmt.Println("分析字符串底层存储细节===")
	// 1.首先string类型底层是一个byte数组
	for i, b := range name {
		fmt.Printf("%v:%v %c\n", i, b, b)
	}
	fmt.Printf("%c%c%c%c\n", 65, 11, 111, 1111)
}
