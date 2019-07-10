package main

import "fmt"

func main() {
	test(16)
}

func test(x int) {
	panic("want panic")
	fmt.Println(x)
}

// output：参数直接值传递
// panic: want panic

// goroutine 1 [running]:
// main.test(0x10)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo1.go:10 +0x39
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo1.go:6 +0x2a
// exit status 2
