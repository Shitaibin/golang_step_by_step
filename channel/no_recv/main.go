package main

import (
	"fmt"
 	"time"
	"context"
)

func main(){
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	ch := make(chan int)
	go producer(ctx, ch)

	time.Sleep(time.Second)
	cancel()
	close(ch)

	time.Sleep(time.Second*10)
}

func producer(ctx context.Context, ch chan int) {
	defer fmt.Println("exit")

	for {
		select {
		default:
			ch <- 1
		}

		select {
		case <- ctx.Done():
			return
		}
	}
}