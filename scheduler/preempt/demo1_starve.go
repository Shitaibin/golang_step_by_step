package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("go: %d\n", i)
		}
	}()

	// 饿死：没给抢占的机会
	for i := 0; ; i++ {
	}
}
