package main

import (
	"fmt"
	"sync"
)

var numbers = 0

var wg = sync.WaitGroup{}

// sync.Mutex 互斥锁
// lock() unlock()
var lock = sync.Mutex{}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			for i := 0; i < 10000; i++ {
				//lock.Lock()
				numbers++
				//lock.Unlock()
			}
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("numbers:%v\n", numbers)

}
