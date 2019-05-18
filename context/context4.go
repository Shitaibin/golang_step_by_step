// 多次取消context，无问题

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Bad way
	// the cancel function returned by context.WithDeadline
	// should be called, not discarded, to avoid a context leakgo-vet
	// ctxTimeout, _ := context.WithDeadline(ctx, time.Now().Add(time.Second))

	// Good way
	ctxTimeout, cancelWork := context.WithDeadline(ctx, time.Now().Add(time.Second))

	go work(ctxTimeout, "reading")

	// context结束后，应尽快取消Work，防止context泄露
	time.Sleep(time.Second * 2)
	cancelWork()
	cancelWork()
	time.Sleep(time.Second)
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
// ➜  context git:(master) ✗ go run context3.go
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
// work is canceled, reason: context deadline exceeded // 超时取消
