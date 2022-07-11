package main

import "fmt"

/**
类型定义：type newTypeName TypeName
类型别名：type newTypeName = TypeName

类型定义和类型别名的区别:
1.类型定义相当于一个全新的定义，与之前的类型不同
2.类型别名只在代码中存在，在编译完成之后不会存在该别名
3.类型别名只是原类型的一个别名，可以使用原类型所拥有的方法，但是类型定义就不能使用(压根就没关系)
*/

// 类型定义
func testTypeDefinition() {
	// 类型定义
	type myInt int
	var a myInt = 20
	fmt.Printf("值:%v 类型:%T\n", a, a)
}

// 类型别名
func testTypeAlias() {
	type myInt = int
	var a myInt = 20
	fmt.Printf("值:%v 类型:%T\n", a, a)
}
func main() {
	testTypeDefinition() // 20  main.myInt  (一种新的类型)
	testTypeAlias()      // 20  int
}
