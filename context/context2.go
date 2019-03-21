// WithTimeout

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()

	// Bad way
	// the cancel function returned by context.WithTimeout
	// should be called, not discarded, to avoid a context leakgo-vet
	// ctxTimeout, _ := context.WithTimeout(ctx, time.Second)

	// Good way
	ctxTimeout, cancelWork := context.WithTimeout(ctx, time.Second)

	go work(ctxTimeout, "reading")

	// context结束后，应尽快取消Work，防止context泄露
	time.Sleep(time.Second * 2)
	cancelWork()
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
// ➜  context git:(master) ✗ go run context2.go
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
// work is canceled, reason: context deadline exceeded
