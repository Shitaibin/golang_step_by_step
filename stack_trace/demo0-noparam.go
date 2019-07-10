package main

func main() {
	test()
}

func test() {
	panic("want panic")
}

// output
// panic: want panic

// goroutine 1 [running]:
// main.test(...)
//         /Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo0.go:8
// main.main()
//         /Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo0.go:4 +0x3a
// exit status 2
