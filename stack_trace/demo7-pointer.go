package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := []string{"hi"}[:]
	fmt.Printf("addr of s: %p\n", &s)
	sp := (*Slice)(unsafe.Pointer(&s))
	fmt.Printf("slice.array: %p, slice.len: 0x%x, slice.cap: 0x%x\n", sp.array, sp.len, sp.cap)
	test(&s)
}

func test(s *[]string) {
	fmt.Printf("addr of s: %p\n", s)
	sp := (*Slice)(unsafe.Pointer(s))
	fmt.Printf("slice.array: %p, slice.len: 0x%x, slice.cap: 0x%x\n", sp.array, sp.len, sp.cap)
	panic("want panic")
}

// 定义同runtime.slice
type Slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

// output：传递slice的指针，就不会新建slice的对象，只传递指针即可，指针本质是1个64位值
// addr of s: 0xc42000a060
// slice.array: 0xc42000e1e0, slice.len: 0x1, slice.cap: 0x1
// addr of s: 0xc42000a060
// slice.array: 0xc42000e1e0, slice.len: 0x1, slice.cap: 0x1
// panic: want panic

// goroutine 1 [running]:
// main.test(0xc42000a060)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo7-pointer.go:20 +0x166
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo7-pointer.go:13 +0x1dd
// exit status 2
