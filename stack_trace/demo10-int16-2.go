package main

import "fmt"

func main() {
	test(int16(1), int16(2), int16(3), int16(4))
}

func test(a, b, c, d int16) {
	panic("want panic")
	fmt.Println(a, b, c, d)
}

// output：为了对齐和节省内存，4个参数合并到1个，看到1，2，3了么
// panic: want panic

// goroutine 1 [running]:
// main.test(0x4000300020001)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo10-int16-2.go:10 +0x39
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo10-int16-2.go:6 +0x30
// exit status 2
