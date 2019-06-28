package main

import "fmt"

func main() {
	ch := make(chan int)

	// 增加关闭
	close(ch)

	select {
	case v, ok := <-ch:
		fmt.Printf("v: %v, ok: %v\n", v, ok)
	}
}

// v: 0, ok: false
