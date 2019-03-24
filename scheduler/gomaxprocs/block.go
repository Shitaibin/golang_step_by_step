package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("hello world")
	}()
	go func() {
		for {

		}
	}()
	select {}
}

// go程序启动时会首先创建一个特殊的内核线程 sysmon，用来监控和管理，其内部是一个循环：

// 记录所有 P 的 G 任务的计数 schedtick，schedtick会在每执行一个G任务后递增

// 如果检查到 schedtick 一直没有递增，说明这个 P 一直在执行同一个 G 任务，如果超过10ms，就在这个G任务的栈信息里面加一个 tag 标记

// 然后这个 G 任务在执行的时候，如果遇到非内联函数调用，就会检查一次这个标记，然后中断自己，把自己加到队列末尾，执行下一个G

// 如果没有遇到非内联函数 调用的话，那就会一直执行这个G任务，直到它自己结束；如果是个死循环，并且 GOMAXPROCS=1 的话。那么一直只会只有一个 P 与一个 M，且队列中的其他 G 不会被执行！
