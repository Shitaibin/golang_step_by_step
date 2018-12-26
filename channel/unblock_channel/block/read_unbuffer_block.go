package main

import "fmt"

func main() {
	ReadNoDataFromNoBufCh()
}

// 场景1
func ReadNoDataFromNoBufCh() {
	noBufCh := make(chan int)

	<-noBufCh
	fmt.Println("read from no buffer channel success")

	// Output:
	// fatal error: all goroutines are asleep - deadlock!
}
