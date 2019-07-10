package main

import (
	"fmt"
)

func main() {
	s := "trace"
	fmt.Printf("addr of s    : %p\n", &s)
	test(s)
}

func test(s string) {
	fmt.Printf("addr of s    : %p, len(s)=0x%x\n", &s, len(s))
	panic("want panic")
}

// output: 也是值传递，但传了1个string，为啥stack trace显示2个参数？
// addr of s    : 0xc42000e1e0
// addr of s    : 0xc42000e1f0, len(s)=0x5
// panic: want panic

// goroutine 1 [running]:
// main.test(0x10c16e7, 0x5)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo2-string.go:15 +0x106
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo2-string.go:10 +0xc3
// exit status 2
