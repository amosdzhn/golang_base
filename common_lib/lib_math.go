package main

import (
	"fmt"
	"math"
)

/**
math 中提供了一些数学常量(PI等)，在math/const.go中
以及一些有用的数学函数,sin,cos,绝对值，平方根等
*/

//测试math包中的一些常量
func testMathConst() {
	fmt.Printf("maxInt8=%v \n", math.MaxInt8) // 127
	fmt.Printf("maxInt16=%v \n", math.MaxInt16)
	fmt.Printf("maxInt32=%d \n", math.MaxInt32)
	fmt.Printf("maxInt64=%d \n", math.MaxInt64)
	fmt.Printf("maxInt=%v \n", math.MaxInt) // 9223372036854775807

	fmt.Printf("PI=%.0f \n", math.Pi)
	fmt.Printf("PI=%.0f \n", math.Pi)
	fmt.Printf("PI=%.10f \n", math.Pi)
}
func main() {
	testMathConst()

	// 测试绝对值 Abs
	num1 := -200
	fmt.Printf("%v 的绝对值=%v \n", num1, math.Abs(float64(num1)))

	// 测试次方操作 Pow
	fmt.Printf("%v的2次方=%v \n", num1, math.Pow(float64(num1), 2))

	// 开平方Sqrt
	fmt.Printf("100 开平方=%v \n", math.Sqrt(float64(100)))

	// 开立方Cbrt
	fmt.Printf("100 开立方=%v \n", math.Cbrt(float64(100)))

	// 向上取整
	fmt.Printf("2.1向上取整=%v\n", math.Ceil(2.1))

	// 向下取整
	fmt.Printf("2.6向下取整=%v\n", math.Floor(2.6))

	// 取余数
	fmt.Printf("10除以3的余数是=%v\n", math.Mod(float64(10), float64(3)))

	// 获取小数部分和整数部分
	modf, frac := math.Modf(10.33)
	fmt.Printf("10.33 的整数部分=%v 小数部分=%.5f \n", modf, frac)
}
