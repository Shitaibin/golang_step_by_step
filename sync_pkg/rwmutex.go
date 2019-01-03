package main

import (
	"fmt"
	"sync"
	"time"
)

type Bank struct {
	sync.RWMutex
	saving map[string]int // 每账户的存款金额
}

func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b
}

// Deposit 存款
func (b *Bank) Deposit(name string, amount int) {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount
}

// Withdraw 取款，返回实际取到的金额
func (b *Bank) Withdraw(name string, amount int) int {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		return 0
	}
	if b.saving[name] < amount {
		amount = b.saving[name]
	}
	b.saving[name] -= amount

	return amount
}

// Query 查询余额
func (b *Bank) Query(name string) int {
	b.RLock()
	defer b.RUnlock()

	if _, ok := b.saving[name]; !ok {
		return 0
	}

	return b.saving[name]
}

func main() {
	b := NewBank()
	go b.Deposit("xiaoming", 100)
	go b.Withdraw("xiaoming", 20)
	go b.Deposit("xiaogang", 2000)

	time.Sleep(time.Second)
	print := func(name string) {
		fmt.Printf("%s has: %d\n", name, b.Query(name))
	}

	nameList := []string{"xiaoming", "xiaogang", "xiaohong", "xiaozhang"}
	for _, name := range nameList {
		go print(name)
	}

	time.Sleep(time.Second)
}
