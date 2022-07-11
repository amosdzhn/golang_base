package main

import (
	"fmt"
	"time"
)

/**
协程(coroutine):比线程更轻量级
				主协程退出，整个程序退出
*/

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 1000)
	}
}
func main() {
	fmt.Println("start...")
	//showMsg("no coroutine")
	//go showMsg("coroutine")
	fmt.Println("end...")
}
