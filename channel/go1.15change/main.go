package main

import "fmt"

func main() {
	ch := make(chan int)

	// v, ok := <-ch
	// fmt.Printf("Read from open channel: v = %v, ok = %v\n", v, ok)

	close(ch)
	v, ok := <-ch
	fmt.Printf("Read from close channel: v = %v, ok = %v\n", v, ok)
}
