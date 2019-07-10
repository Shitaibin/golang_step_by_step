package main

import "fmt"

func main() {
	i8 := int16(0x123)
	test(i8)
}

func test(x int16) {
	fmt.Printf("addr of x: %p\n", &x)
	panic("want panic")
}

// output：参数直接值传递，有趣了，传递进来的是个地址?实际不是，看到末尾的123了么，为了节省内存，以及内存对齐
// addr of x: 0xc42001607a
// panic: want panic

// goroutine 1 [running]:
// main.test(0xc420080123)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo9-int16.go:12 +0xa8
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo9-int16.go:7 +0x28
// exit status 2
