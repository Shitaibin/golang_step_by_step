package doublecheck

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func (o *Once) DoSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()

	if o.done == 1 {
		return
	}
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func (o *Once) DoFast(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

func Testunsafety() {
	var once Once
	for i := 0; i < 10000; i++ {
		go func() {
			go once.Do(f)
		}()
	}
}

func TestSlow() {
	var once Once
	for i := 0; i < 10000; i++ {
		go func() {
			go once.DoSlow(f)
		}()
	}
}

func TestFast() {
	var once Once
	for i := 0; i < 10000; i++ {
		go func() {
			go once.DoFast(f)
		}()
	}
}

func f() {
	// fmt.Println("f")
}
