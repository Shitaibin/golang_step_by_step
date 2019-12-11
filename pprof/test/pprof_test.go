package test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func Entrance() {
	// 设置block采样频率，单位ns
	runtime.SetBlockProfileRate(1 * 1000 * 1000)
	// 设置mutex采样频率
	runtime.SetMutexProfileFraction(1)

	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go LockAndSleep(&mu, &wg)
	go LockAndSleep(&mu, &wg)
	wg.Wait()
}

func LockAndSleep(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	time.Sleep(time.Second)
	fmt.Println("Exit")
}

func TestEntrance(t *testing.T) {
	Entrance()
}

func BenchmarkEntrace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Entrance()
	}
}
