package main

import (
	"fmt"
	"sync"
	"time"
)

// channel的遍历
func main() {
	// 无缓冲通道，输入 大于 输出
	//testChannel1()
	// 无缓冲通道，输出 小于 输出
	//testChannel2()
	// 无缓冲通道，遍历
	//testChannelFor()
	// 无缓冲通道，for range遍历
	testChannelForRange()
}

//无缓冲通道 管道输入的个数 大于 管道输出的个数
func testChannel1() {
	var c = make(chan int)
	defer close(c)
	var wg = sync.WaitGroup{}

	fmt.Println("测试channel:")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start coroutine")
		for i := 0; i < 10; i++ {
			c <- i
		}
		fmt.Println("end coroutine")
	}()
	wg.Done()
	var r1 int = <-c
	var r2 = <-c
	r3 := <-c
	fmt.Println(r1, r2, r3) // 0 1 2
}

//无缓冲通道 管道输入的个数 小于 管道输出的个数
func testChannel2() {
	var wg = sync.WaitGroup{}
	var c = make(chan int)
	defer close(c) // 从通道读取后关闭通道
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		defer wg.Done()
		for i := 0; i < 2; i++ {
			c <- i
		}
		close(c) // 在输入结束后关闭通道，那么此时r3 将会获取一个默认零值，int为0
	}()
	wg.Wait()
	r1 := <-c
	r2 := <-c
	r3 := <-c // deadlock 当前channel没有输入了，但此时一致阻塞
	fmt.Println(r1, r2, r3)
}

//无缓冲通道 遍历通道
func testChannelFor() {
	var wg = sync.WaitGroup{}
	var channel = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()
	for i := 0; i < 10; i++ {
		r := <-channel
		fmt.Printf("从channel获取值:%v\n", r)
	}
	wg.Wait()
}

// 无缓冲通道 for range遍历
func testChannelForRange() {
	var channel = make(chan int)
	go func() {
		defer close(channel)
		fmt.Println("start goroutine")
		for i := 0; i < 20; i++ {
			channel <- i
		}
		fmt.Println("end goroutine")
	}()
	for v := range channel {
		fmt.Printf("从channel获取值:%v\n", v)
	}
}
