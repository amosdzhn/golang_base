package main

import (
	"fmt"
	"time"
)

/**
time包提供了测量和显示时间的功能包

time包中的结构体:
ParseError{Layout string,Value string,LayoutElem string,ValueElem string,Message string} 用于描述解析时间字符串时出现的错误
type Weekday int 代表一周中的某天 const(Sunday Weekday=iota Monday......)
type Month int 代表一年中的某个月 const(January Month = 1 + iota February......)
Location 是一个结构体，代表某个地点或时区
Time 是一个结构体，代表一个纳秒精确的时间点
Duration
Timer
Ticker
*/

//
func testTimeStruct() {
	fmt.Println("test struct:Time")
	now := time.Now()
	// 输入Time类型变量
	fmt.Printf("type:%T value:%v \n", now, now)
	// output: type: time.Time value:2022-07-05 19:17:14.4614138 +0800 CST m=+0.003120101
	nowYear := now.Year()
	nowMonth := now.Month()
	nowDay := now.Day()
	nowHour := now.Hour()
	nowSecond := now.Second()
	nowMinute := now.Minute()
	nowNanosecond := now.Nanosecond()
	fmt.Printf("年:%v 月:%v 日:%v 时:%v 分:%v 秒:%v 纳秒:%v\n", nowYear, nowMonth, nowDay, nowHour, nowSecond, nowMinute, nowNanosecond)
	// output:年:2022 月:July 日:5 时:19 分:3 秒:22 纳秒:823207200
	format := now.Format(time.Stamp)
	fmt.Println(format)
	// 时间戳,获取1970年1月1日0点开始的总毫秒数
	// 秒 纳秒 毫秒 纳秒
	unixTimeStamp := now.Unix()
	unixTimeStampNano := now.UnixNano()
	unixTimeStampMilli := now.UnixMilli()
	unixTimeStampMicro := now.UnixMicro()
	fmt.Printf("时间戳:  unixTimeStamp=%v unixTimeStampNano=%v unixTimeStampMilli=%v unixTimeStampMicro=%v \n", unixTimeStamp, unixTimeStampNano, unixTimeStampMilli, unixTimeStampMicro)
	// output:1657034748 1657034748186296100 1657034748186 1657034748186296
}

// 时间戳转换为时间标准格式
func testTimeStampToStdTime() {
	// 获取一个时间戳
	timeStamp := time.Now().Unix()
	now := time.Unix(timeStamp, 0)
	fmt.Printf("转为为的时间是:%v\n", now)
}
func main() {
	testTimeStruct()
	testTimeStampToStdTime()
}
