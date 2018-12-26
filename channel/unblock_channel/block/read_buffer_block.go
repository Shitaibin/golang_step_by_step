package main

import "fmt"

func main() {
	ReadNoDataFromBufCh()
}

// 场景1
func ReadNoDataFromBufCh() {
	bufCh := make(chan int, 1)

	<-bufCh
	fmt.Println("read from no buffer channel success")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
