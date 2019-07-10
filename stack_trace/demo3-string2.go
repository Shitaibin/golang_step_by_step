package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "trace"
	fmt.Printf("addr of s    : 0x%p\n", &s)
	test(s)
}

func test(s string) {
	fmt.Printf("addr of s    : %p, len(s)=0x%x\n", &s, len(s))

	// 打印string底层的存储地址
	sp := (*StringStruct)(unsafe.Pointer(&s))
	fmt.Printf("addr of s.str: %p, len of s: 0x%x\n", sp.str, sp.len)

	panic("want panic")
}

type StringStruct struct {
	str unsafe.Pointer
	len int
}

// output: 也是值传递，入参为string的实际地址和长度
// addr of s    : 0x0xc42000e1e0
// addr of s    : 0xc42000e1f0, len(s)=0x5
// addr of s.str: 0x10c17a7, len of s: 0x5
// panic: want panic

// goroutine 1 [running]:
// main.test(0x10c17a7, 0x5)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo3-string2.go:20 +0x1bd
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo3-string2.go:11 +0xc3
// exit status 2

// string的定义
// type stringStruct struct {
// 	str unsafe.Pointer
// 	len int
// }
