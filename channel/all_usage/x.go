package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	// ch2 := make(chan int, 10)

	
	ch1<-1
	ch1<-2
	close(ch1)
	for x :=range handle(ch1) {
		fmt.Println(x)
	}
	time.Sleep(time.Second)
}

func producer(outCh chan<- int)
func consumer(inCh <-chan int)  chan int {
	outCh := make(chan int, 10)
	go func(){
		for x := range inCh {
			outCh <- x*x
		}
	}()
	return outCh
}
