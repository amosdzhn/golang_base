package main

import (
	"fmt"
	"time"
)

func main() {
	var intChan = make(chan int)
	go func() {
		intChan <- 100
		close(intChan)
	}()
	time.Sleep(time.Second * 2)
	r, ok := <-intChan
	if ok {
		fmt.Println("intChan 开启")
	} else {
		fmt.Println("intChan 被关闭")
	}
	fmt.Println(r)
}
