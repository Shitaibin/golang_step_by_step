package main

// 两种运行方式：
// GOMAXPROCS=1 go run block.go
// 只有1个P，2个goroutine都在同1个P上，所以不可能同时运行，打印结果是，交替打印3000+
// 个“go：”，然后3000多个“main：”
//
// go run block.go
// 使用默认的，本机有8个核，所以有8个P，有2个goroutine，因为打印是标准输出，存在syscall，
// 如果两个goroutine刚开始在同一个P上，syscall造成线程M阻塞时，会把另外1个goroutine转移
// 到P‘上，P’又被另外线程M‘运行，M和M’“可能”在不同核上，是并行的，打印就变成了打印几行“go：”，
// 再打印几行“main：”，就这样往复

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("go: %d\n", i)
		}
	}()

	// 不饿死
	// for i := 0; ; i++ {
	// 	fmt.Printf("main: %d\n", i)
	// }

	// 饿死：没给抢占的机会
	for i := 0; ; i++ {
	}

	time.Sleep(time.Second)
}
