package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	select {
	case v, ok := <-ch:
		fmt.Printf("v: %v, ok: %v\n", v, ok)
	}
}

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [chan receive]:
// main.main()
//         /Users/shitaibin/Workspace/golang_step_by_step/channel/ok/demo_select1.go:8 +0x69
// exit status 2
