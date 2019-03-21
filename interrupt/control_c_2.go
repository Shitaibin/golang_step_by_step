package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	stop := make(chan struct{})
	go func() {
		t := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			// 处理中断
			case <-stop:
				fmt.Printf("Goroutine exit\n")
				return
			case <-t.C:
				fmt.Println("Running")
			}
		}
	}()

	// 捕获中断，然后通知协程退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, os.Kill)
	<-sigCh
	close(stop)
}

// Output 1：ctrl-c
// ➜  interrupt git:(master) ✗ go run control_c_2.go
// Running
// Running
// Running
// Running
// Running
// Running
// Running
// ^C%

// Output 2: kill 命令
// ➜  interrupt git:(master) ✗ go run control_c_2.go
// Running
// Running
// Running
// Running
// Running
// Running
// signal: killed
