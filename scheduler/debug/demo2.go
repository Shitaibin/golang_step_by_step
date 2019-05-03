package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Hello goroutine")
	}()
	fmt.Println("Hello schedule")
	time.Sleep(time.Millisecond)
}
