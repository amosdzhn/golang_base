package main

import (
	"fmt"
	"sync"
	"time"
)

/**
async.WaitGroup 用于解决协程之间的同步问题
*/

func main() {
	// 声明一个sync并发包下的WaitGroup结构体类型
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 每当开启一个协程,执行Add
		wg.Add(1)
		go func(a int) {
			// 当一个协程执行结束调用Done
			defer wg.Done()
			time.Sleep(time.Millisecond * 1000) // 睡眠100毫秒
			fmt.Println(a, "执行了")
		}(i)
	}
	// 当前协程需要等到waitGroup中没有协程在执行
	wg.Wait()
	fmt.Println("main end..")

}
