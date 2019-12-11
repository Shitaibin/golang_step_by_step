package sync_pkg

import (
	"runtime"
	"sync"
	"testing"
)

type Person struct {
	name string
}

func gcBeforeUse(p *Person) string {
	runtime.GC()
	return p.name
}

func BenchmarkPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} { return new(Person) },
	}

	for i := 0; i < b.N; i++ {
		p := pool.Get().(*Person)
		gcBeforeUse(p)
		pool.Put(p)
	}
}

func BenchmarkPoolUsingAfterPut(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} { return new(Person) },
	}

	for i := 0; i < b.N; i++ {
		p := pool.Get().(*Person)
		pool.Put(p)
		gcBeforeUse(p)
	}
}
