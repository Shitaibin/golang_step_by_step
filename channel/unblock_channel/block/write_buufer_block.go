package main

import "fmt"

func main() {
	WriteBufChButFull()
}

// 场景2
func WriteBufChButFull() {
	ch := make(chan int, 1)
	// make ch full
	ch <- 100

	ch <- 1
	fmt.Println("write success no block")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
