package main

import "fmt"

/**
map的数据类型是：map[keyType]valueType
*/
func main() {
	testMapDefinition()

	testMapSimpleUse()

	testMapCURD()
}

// 测试声明map
// 声明变量 var name map[k_type]v_type  默认为nil
// 只开辟空间 name = make(map[k_type]v_type)
// 开辟空间并赋值 name = map[k_type]v_type{}
func testMapDefinition() {
	// 1.先声明变量，再开辟内存空间
	var m1 map[string]int // m1 现在只是nil
	fmt.Printf("只声明的m1的值:%v 类型:%T 地址:%v\n", m1, m1, &m1)
	m1 = make(map[string]int)
	fmt.Printf("只声明的m1的值:%v 类型:%T 地址:%v\n", m1, m1, &m1)
	// 2.声明变量并开辟内存空间
	var m2 = make(map[string]int)
	fmt.Printf("只声明的m2的值:%v 类型:%T 地址:%v\n", m2, m2, &m2)
	// 3.声明变量开辟内存空间，再赋值
	var m3 map[string]string = map[string]string{"k1": "v1"}
	fmt.Printf("只声明的m3的值:%v 类型:%T 地址:%v\n", m3, m3, &m3)
}

// 简单的使用map
func testMapSimpleUse() {
	//先声明，再初始化
	var stu map[string]int
	stu = make(map[string]int)
	stu["张三"] = 19
	stu["李四"] = 29
	fmt.Printf("map: %v  size:%v", stu, len(stu))
}

// 测试Map的CURD操作
func testMapCURD() {
	// 1.创建map(创建map变量，开辟空间，初始化赋值)
	var person = map[string]int{"张三": 19, "李四": 20}
	fmt.Printf("创建后:%v \n", person)
	// 2.查询,判断是否存在
	fmt.Printf("查询王五:%v \n", person["王五"]) // 不存在 0
	value, ok := person["王五"]
	fmt.Printf("是否存在:%v 值:%v\n", ok, value) // false 0
	value, ok = person["张三"]
	fmt.Printf("是否存在:%v 值:%v\n", ok, value) // true 19
	// 3.增加和修改元素
	person["王五"] = 99
	fmt.Printf("增加后:%v \n", person) // 增加不存在的相当于增加
	person["王五"] = 100
	fmt.Printf("增加后:%v \n", person) // 增加存在的相当于修改
	// 4.遍历map for range
	for key := range person {
		fmt.Println(key)
	}
	for key, value := range person {
		fmt.Println(key, value)
	}
}
