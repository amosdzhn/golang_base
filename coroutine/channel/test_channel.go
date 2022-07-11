package main

import (
	"fmt"
	"time"
)

/**
channel(通道)：用于协程之间共享数据
背景：当多个coroutine之间并发执行，且需要协程之间共享数据时，channel充当coroutine之间的管道(通道)并提供
一种机制来保证同步交换。

根据channel的交换行为，有两种通道，无缓冲和有缓冲。无缓冲用于两个协程之间的同步通信。
缓冲通道用于执行异步通信。通道使用后需要关闭

channel由make(chan type)创建
//创建一个无缓冲通道
ch1 := make(chan int)
//向该通道放值
ch1 <- 500
//从通道取值
data := <- ch1
*/
func main() {
	// 创建无缓冲通道
	channel := make(chan int)
	defer close(channel)
	go func() {
		fmt.Println("start coroutine")
		time.Sleep(time.Millisecond * 2000)
		channel <- 200
		fmt.Println("coroutine end..")
	}()
	fmt.Println("main coroutine正在等待获取通道值")
	value := <-channel
	fmt.Println("获取到了值:", value)
	fmt.Println("main end.")
}
