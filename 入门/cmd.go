package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("os.Args type: %T\n", os.Args)
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("第%v个参数是:%v\n", i, os.Args[i])
	}
}
