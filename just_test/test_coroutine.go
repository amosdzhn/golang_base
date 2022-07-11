package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	intChan := make(chan int)
	defer close(intChan)
	wg.Add(1)
	go func() {
		intChan <- 20
		wg.Done()
	}()
	wg.Wait()
	res := <-intChan
	fmt.Printf("res=%v \n", res)
}
