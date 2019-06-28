package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	v, ok := <-ch
	fmt.Printf("v: %v, ok: %v\n", v, ok)
}

// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan receive]:
