package timer_demo

import (
	"sync"
	"testing"
	"time"
)

func waitTimer() {
	n := 1000000
	var wg sync.WaitGroup

	wait := func() {
		defer wg.Done()
		// defer fmt.Println("exit")

		ch := time.Tick(time.Second)
		<-ch
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go wait()
	}

	wg.Wait()
	// fmt.Println("All exit")
}

func TestTest(t *testing.T) {
	waitTimer()
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		waitTimer()
	}
}
