package main

import "fmt"

/**
探究Go中字符串底层
*/
func main() {
	// 1.字符串零值是空串
	var s1 string
	fmt.Println(s1) // ""
	// 2.使用索引取出的字符是字节，不是字符,不能取字节元素的地址
	var s2 string = "我的中国心"
	fmt.Println(s2)
	fmt.Println(s2[0]) // 230
	fmt.Println(&s2)
	//fmt.Println(&s2[0])  // 可取字符串地址，不能取字符串中字节元素的地址
	//3.字符串是不可变类型，无法修改底层字节数组。
	//可以使用[]rune进行修改，但是重新分配内存空间，并复制字节数组
	s3 := []rune(s2)
	s3[0] = '你'
	s4 := string(s3)
	fmt.Println(s4)
	fmt.Printf("s2地址:%v \ns3[0]地址:%v \ns4地址:%v\n", &s2, &s3[0], &s4)

	stringDefinition()
	stringJoin()
	stringSlice()
	stringFor()
}

func stringDefinition() {
	// 定义字符串
	s1 := "first string"
	s2 := `line1
line2
`
	fmt.Println(s1)
	fmt.Println(s2)
}

func stringJoin() {
	// 字符串拼接
	c1 := "你"
	c2 := "好"
	c3 := "吗"

	s1 := c1 + c2
	fmt.Printf("%v+%v: %v\n", c1, c2, s1)
	var s2 string
	s2 += c3
	s2 += c3
	fmt.Printf("s2: %v\n", s2)
}

func stringSlice() {
	// 通过切片获取字串,字串是切片类型
	var name string = "爱新觉罗.玄烨"
	// go字符串源码使用utf-8编码，在string类型底层是按照字节数组存储，一个字符占3个字节
	var subName string = string([]rune(name)[:4])
	fmt.Printf("截取的子字符串：%v\n", subName)
}

func stringFor() {
	// 遍历字符串
	// 1.string底层是按照字节数组存储的
	var s1 = "A"
	fmt.Printf("len(s1):%v\n", len(s1)) // 1
	var s2 = "中"
	fmt.Printf("len(s2):%v\n", len(s2)) // 3
	var s3 = "ABC中国"
	fmt.Printf("len(s3):%v\n", len(s3)) // 9 = 3 + 3*2
	// 2.
	fmt.Printf("s1[0]:%v\n", s1[0])                 // 65
	fmt.Printf("string(s1[0]):%v\n", string(s1[0])) // "A"
	// 3.遍历
	var content = "ABC,一二三"
	for i := 0; i < len(content); i++ {
		fmt.Println(content[i])
	}
	/*
		65   => A
		66      B
		67      C
		44      ,
		228  ]
		184	 ]  一
		128  ]
		228
		186
		140
		228
		184
		137
	*/
	for i, v := range content {
		fmt.Println(i, v)
	}
	/**
	0 65    A
	1 66	B
	2 67	C
	3 44	,
	4 19968  一
	7 20108  二
	10 19977  三
	*/
	for _, v := range content {
		fmt.Printf("%v\n", string(v))
	}
}
