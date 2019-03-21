// WithCancel

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctxCancel, cancelWork := context.WithCancel(ctx)

	go work(ctxCancel, "reading")

	// 运行一段时间后取消work
	time.Sleep(time.Second)
	cancelWork()

	// 等待work打印cancel的结果
	time.Sleep(time.Millisecond * 500)
}

func work(ctx context.Context, workName string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("work is canceled, reason: %s\n", ctx.Err())
			return
		default:
			fmt.Printf("%s work is running\n", workName)
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// Output
// ➜  context git:(master) ✗ go run context1.go
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// reading work is running
// work is canceled, reason: context canceled // 主动取消
