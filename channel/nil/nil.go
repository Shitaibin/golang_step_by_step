package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
		fmt.Println("g1 exit")
	}()

	go func(ch chan int) {
		if ch == nil {
			fmt.Println("ch is nil")
		}

		<-ch
		fmt.Println("g2 exit")
	}(ch)

	time.Sleep(time.Second)
}