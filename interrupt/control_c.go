package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	sigCh := make(chan os.Signal, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		t := time.NewTicker(time.Millisecond * 500)

		for {
			select {
			// 处理中断
			case sig := <-sigCh:
				fmt.Printf("Catch %v\n", sig.String())
				wg.Done()
				return
			case <-t.C:
				fmt.Println("Running")
			}
		}
	}()

	// 捕获中断
	signal.Notify(sigCh, os.Interrupt, os.Kill)

	wg.Wait()
}

// Output
// ➜  interrupt git:(master) ✗ go run control_c.go
// Running
// Running
// Running
// ^CCatch interrupt
