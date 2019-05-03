package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		defer wg.Done()
		wg.Add(1)
		defer wg.Done()
	}()
	wg.Wait()
}
