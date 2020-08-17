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

	// go1.13及之前
	// 饿死上面的协程：没给抢占的机会，所以不会有打印
	// go1.14引入了基于信号的抢占，可以有打印了
	for i := 0; ; i++ {
	}
}
