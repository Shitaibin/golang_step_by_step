package benchtest

import (
	"sync"
	"sync/atomic"
	"testing"

	"github.com/stretchr/testify/assert"
)

var N = 100000

func atom(b *testing.B) {
	var ops int32

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			atomic.AddInt32(&ops, 1)
			wg.Done()
		}()
	}

	wg.Wait()
	// opsFinal := atomic.LoadUint64(&ops)
	assert.Equal(b, N, int(ops), "not equal")
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atom(b)
	}
}

func ch(b *testing.B) {
	var ops int
	ch := make(chan bool, 1)

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			ch <- true
			ops = ops + 1
			<-ch
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(b, N, ops, "not equal")
}

func mute(b *testing.B) {
	var ops int
	var m sync.Mutex

	var wg sync.WaitGroup
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			m.Lock()
			ops = ops + 1
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	assert.Equal(b, N, ops, "not equal")
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch(b)
	}
}

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mute(b)
	}
}

// goos: darwin
// goarch: amd64
// BenchmarkAtomic-8    	      50	  34943502 ns/op	     859 B/op	       7 allocs/op
// BenchmarkChannel-8   	      10	 125640858 ns/op	 2651630 B/op	   13491 allocs/op
// BenchmarkMutex-8     	      50	  35282788 ns/op	     894 B/op	      13 allocs/op
// 讨论链接：https://github.com/developer-learning/reading-go/issues/364
