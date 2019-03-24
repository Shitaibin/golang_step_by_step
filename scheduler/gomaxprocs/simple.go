package main

//
// GOMAXPROCS=1 go run simple.go
// go run simple.go

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	done := make(chan struct{})

	go test(1, done, 1, testNoSleep)
	<-done

	go test(2, done, 8, testNoSleep)
	<-done

	go test(3, done, 1, testWithSleeep)
	<-done

	go test(4, done, 8, testWithSleeep)
	<-done
}

func test(id int, done chan struct{}, n int, f func(chan int)) {
	fmt.Printf("---------------- test %d ---------------------\n", id)
	runtime.GOMAXPROCS(n)
	ch := make(chan int, 1)
	fmt.Printf("num of CPU: %d\n", runtime.NumCPU())
	fmt.Printf("begin: num of goroutine: %d\n", runtime.NumGoroutine())

	f(ch)
	time.Sleep(time.Microsecond * 100)
	fmt.Printf("middle in test: num of goroutine: %d\n", runtime.NumGoroutine())

	time.Sleep(time.Second)
	fmt.Printf("end: num of goroutine: %d\n", runtime.NumGoroutine())

	// 看运行程序时的GOMAXPROCS
	fmt.Printf("old GOMAXPROCS: %d\n", runtime.GOMAXPROCS(1))

	done <- struct{}{}
}

func testNoSleep(ch chan int) {
	go func() {
		ch <- 1
	}()

	go func() {
		fmt.Printf("read: %d\n", <-ch)
	}()

	fmt.Printf("middle: num of goroutine: %d\n", runtime.NumGoroutine())
}

func testWithSleeep(ch chan int) {
	go func() {
		fmt.Printf("read: %d\n", <-ch)
		time.Sleep(time.Microsecond * 400)
	}()

	fmt.Printf("middle: num of goroutine: %d\n", runtime.NumGoroutine())

	go func() {
		ch <- 1
		time.Sleep(time.Microsecond * 400)
	}()
}
