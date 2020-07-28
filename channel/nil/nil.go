package main

import (
	"fmt"
	"time"
)

func main() {
	// default value of ch is nil
	var ch chan int
	go func() {
		// ch point to an new channel object
		ch = make(chan int, 1)
		ch <- 1
		fmt.Println("g1 exit")
	}()

	// nil channel pass to function
	go func(ch chan int) {
		if ch == nil {
			fmt.Println("ch is nil")
		}

		<-ch
		fmt.Println("g2 exit")
	}(ch)

	time.Sleep(time.Second)
}