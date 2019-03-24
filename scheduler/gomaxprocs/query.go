package main

import (
	"fmt"
	"runtime"
)

func main() {
	old := runtime.GOMAXPROCS(-1)
	fmt.Printf("init: %d\n", old)

	for i := 0; i < 10000; i++ {
		old := runtime.GOMAXPROCS(i)
		fmt.Printf("old: %d\n", old)
	}
}
