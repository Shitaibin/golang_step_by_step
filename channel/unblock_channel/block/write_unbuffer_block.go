package main

import "fmt"

func main() {
	WriteNoBufCh()
}

// 场景2
func WriteNoBufCh() {
	ch := make(chan int)

	ch <- 1
	fmt.Println("write success no block")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
