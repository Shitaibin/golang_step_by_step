package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	// 增加关闭
	close(ch)

	// not in select是阻塞读
	v, ok := <-ch
	fmt.Printf("v: %v, ok: %v\n", v, ok)
}

// v: 0, ok: false
