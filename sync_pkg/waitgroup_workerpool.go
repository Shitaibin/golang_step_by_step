package main

import (
	"fmt"
	"sync"
)

func main() {
	jobCh := genJob(10)
	retCh := make(chan string, 10000)
	workerPool(5, jobCh, retCh)

	for ret := range retCh {
		fmt.Println(ret)
	}
}

func genJob(n int) <-chan int {
	jobCh := make(chan int, 200)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()

	return jobCh
}

func workerPool(n int, jobCh <-chan int, retCh chan<- string) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(&wg, i, jobCh, retCh)
	}

	wg.Wait()
	close(retCh)
}

func worker(wg *sync.WaitGroup, id int, jobCh <-chan int, retCh chan<- string) {
	cnt := 0
	for job := range jobCh {
		cnt++
		ret := fmt.Sprintf("worker %d processed job: %d, it's the %dth processed by me.", id, job, cnt)
		retCh <- ret
	}

	wg.Done()
}
