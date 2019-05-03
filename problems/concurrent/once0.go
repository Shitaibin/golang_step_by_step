package main

import (
	"fmt"
	"sync"
	"time"
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

func call(o *Once) {
	go o.Do(f)
}

func main() {
	var once Once
	for {
		go call(&once)
	}

	time.Sleep(time.Second)
}

func f() {
	fmt.Println("f")
}
