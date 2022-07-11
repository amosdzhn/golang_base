package main

import "fmt"

/**
map 是引用传递
*/
func testMapCopy(m map[string]string) {
	m["k1"] = "1000"
	fmt.Printf("函数调用中:%v\n", m)
}
func main() {
	var cm = map[string]string{"k1": "v1", "k2": "v2"}
	fmt.Printf("函数调用之前:%v\n", cm)
	testMapCopy(cm)
	fmt.Printf("函数调用之后:%v\n", cm)
}
