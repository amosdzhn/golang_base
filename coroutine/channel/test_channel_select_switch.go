package main

import "time"

// 测试select switch语句
/*
	select语句类似switch语句，是一个控制结构。用于处理异步IO。
	select会监听case中chan的读写操作,当case中的chan读写状态为非阻塞(可读写)时，会触发相应的操作
	注意：select中的case必须是一个chan操作
		select中的default总是可运行的
		当多个case都可以运行时,select会随机公平的选出一个执行，其他不会执行
		如果没有可运行的case语句,且存在default语句，那么就会执行default语句
		如果既没有可运行的case语句，也没有default语句，那么select会被阻塞，直到某个case通信可以运行
*/
func main() {
	var intChan = make(chan int)
	var strChan = make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		intChan <- 100
		close(intChan)
	}()
}
