package main

import (
	"fmt"
)

func main() {
	// 整数
	// test uint8/16/32/64
	var un1 uint8 = 100
	var un2 uint16 = 200
	var un3 uint32 = 300
	var un4 uint64 = 400
	fmt.Printf("un1:%v un2:%v un3:%v un4:%v\n", un1, un2, un3, un4)
	fmt.Printf("uint type:%T\n", un1) // uint8

	// test int8/16/32/64
	var n1 int8 = 100
	var n2 int16 = 200
	var n3 int32 = 300
	var n4 int64 = 400
	fmt.Printf("n1:%v n2:%v n3:%v n4:%v \n", n1, n2, n3, n4)

	// test 机器字大小 根据不同的平台不同可能32位可能64位
	var unn int = 100
	var nn uint = 100
	fmt.Printf("unn:%v nn:%v \n", unn, nn)

	// byte 与 int8等价
	var b byte = 100
	fmt.Printf("b:%v type:%T \n", b, b)
	// rune 和 int32等价 一个Unicode字符中32位即4字节表示一个码点
	var r rune = 300
	fmt.Printf("r:%v type:%T \n", r, r)

	var str rune = 'A'
	fmt.Printf("str:%v type:%T\n", str, str) // 65 int32

	//*************
	var uip uintptr = 100
	fmt.Printf("uip:%v type:%T\n", uip, uip)

	// 数字的运算
	numberOp()
	// 逻辑运算
	logicOp()
	// 比较运算
	compareOp()

}

// 数字之间的运算
func numberOp() {
	// + - * / 可用于(整数，浮点数，复数) % 只能用于整数
	var n1 = 100
	var n2 = 10
	fmt.Printf("加法运算:%v\n", n1+n2)
	fmt.Printf("减法运算:%v\n", n1-n2)
	fmt.Printf("乘法运算:%v\n", n1*n2)
	fmt.Printf("除法运算:%v\n", n1/n2)
}

// 逻辑运算
func logicOp() {

}

// 比较运算
func compareOp() {

}
