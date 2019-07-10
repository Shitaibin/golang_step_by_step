package main

import "fmt"

func main() {
	s := []string{"hi"}[:]
	fmt.Printf("addr of s: %p\n", &s)
	test(s)
}

func test(s []string) {
	fmt.Printf("addr of s: %p\n", &s)
	panic("want panic")
}

// output：顺便解释一下slice的值传递，为何3个参数？是slice结构体的3个字段吗？
// addr of s: 0xc42000a060
// addr of s: 0xc42000a080
// panic: want panic

// goroutine 1 [running]:
// main.test(0xc42000e1e0, 0x1, 0x1)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo4-slice.go:13 +0xd0
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo4-slice.go:8 +0x106
// exit status 2
