package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/**
runtime.Gosched() 让出处理器，让子协程先执行完毕后，再执行当前协程
*/

var wg = sync.WaitGroup{}

func sayHello(msg string) {
	wg.Add(1)
	defer wg.Done()
	time.Sleep(time.Second * 2)
	for i := 0; i < 2; i++ {
		fmt.Println("hello,", msg)
	}
}
func main() {
	fmt.Printf("当前物理机器的CPU核心数:%v\n", runtime.NumCPU()) // 当前是8
	go sayHello("Golang")
	for i := 0; i < 2; i++ {
		runtime.Gosched()
		fmt.Println("main", i)
	}
	wg.Wait()
	fmt.Println("main end..")
}
