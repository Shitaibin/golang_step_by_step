package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 创建请求的通道和银行
	reqCh := make(chan *Request, 100)
	bank := NewBank()

	// 银行处理请求
	go bank.Loop(reqCh)

	// 小明和小刚2个线程同时存取钱
	var wg sync.WaitGroup
	wg.Add(2)
	go xiaoming(&wg, reqCh)
	go xiaogang(&wg, reqCh)

	// 等待小明和小刚完成
	wg.Wait()
	close(reqCh)

	// 等待看银行是否退出
	time.Sleep(time.Second)
}

type Bank struct {
	saving map[string]int // 每账户的存款金额
}

// Request 银行存取操作
type Request struct {
	op    string       // 存、取、查
	name  string       // 操作的账号
	value int          // 操作金额
	retCh chan *Result // 存放银行处理结果的通道
}

// Result 执行结果
type Result struct {
	success bool // 成功
	value   int  // 查询时使用：余额
}

func NewBank() *Bank {
	b := &Bank{
		saving: make(map[string]int),
	}
	return b
}

// Loop 银行处理客户请求
func (b *Bank) Loop(reqCh chan *Request) {
	for req := range reqCh {
		switch req.op {
		case "deposite":
			b.Deposit(req)
		case "withdraw":
			b.Withdraw(req)
		case "query":
			b.Query(req)
		default:
			// 响应
			ret := &Result{
				false,
				0,
			}
			req.retCh <- ret
		}
	}

	// 无请求时银行退出
	fmt.Println("Bank exit")
}

// Deposit 存款
func (b *Bank) Deposit(req *Request) {
	name := req.name
	amount := req.value

	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount

	// 响应
	ret := &Result{
		true,
		0,
	}
	req.retCh <- ret
}

// Withdraw 取款，不足时取款失败
func (b *Bank) Withdraw(req *Request) {
	name := req.name
	amount := req.value

	var status bool
	if balance, ok := b.saving[name]; !ok || balance < amount {
		status = false
		amount = 0
	} else {
		status = true
		b.saving[name] -= amount
	}

	// 响应
	ret := &Result{
		status,
		amount,
	}
	req.retCh <- ret
}

// Query 查询余额
func (b *Bank) Query(req *Request) {
	name := req.name

	var (
		ok      bool
		balance int
	)

	if balance, ok = b.saving[name]; !ok {
		balance = 0
	}

	// 响应
	ret := &Result{
		true,
		balance,
	}
	req.retCh <- ret
}

func xiaoming(wg *sync.WaitGroup, reqCh chan<- *Request) {
	name := "xiaoming"
	retCh := make(chan *Result)
	defer func() {
		close(retCh)
		wg.Done()
	}()

	depReq := &Request{
		"deposite",
		name,
		100,
		retCh,
	}
	withDrawReq := &Request{
		"withdraw",
		name,
		20,
		retCh,
	}
	queryReq := &Request{
		"query",
		name,
		0,
		retCh,
	}

	// 顺序3个请求：存100，花20，剩80
	reqs := []*Request{depReq, withDrawReq, queryReq}
	expRets := []int{0, 0, 80} // 期望Result中返回的值
	for i, req := range reqs {
		reqCh <- req
		waitResp(req, expRets[i])
	}
}

// waitResp 等待请求响应req, expVal是查询时使用的期待值
func waitResp(req *Request, expVal int) {
	ret := <-req.retCh
	if ret.success {
		if req.op != "query" {
			fmt.Printf("%s %s %d success\n", req.name, req.op, req.value)
		} else {
			if ret.value != expVal {
				fmt.Printf("%s query result error, got %d want %d\n", req.name, ret.value, expVal)
			} else {
				fmt.Printf("%s has %d\n", req.name, ret.value)
			}
		}
		return
	}

	if req.op != "query" {
		fmt.Printf("%s %s %d failed\n", req.name, req.op, req.value)
	} else {
		fmt.Printf("%s query failed\n", req.name)
	}
}

// xiaogang 存、花、查
func xiaogang(wg *sync.WaitGroup, reqCh chan<- *Request) {
	name := "xiaogang"
	retCh := make(chan *Result)
	defer func() {
		close(retCh)
		wg.Done()
	}()

	depReq := &Request{
		"deposite",
		name,
		100,
		retCh,
	}
	withDrawReq := &Request{
		"withdraw",
		name,
		200,
		retCh,
	}
	queryReq := &Request{
		"query",
		name,
		0,
		retCh,
	}

	// 顺序3个请求：存100，花200失败，剩100
	reqs := []*Request{depReq, withDrawReq, queryReq}
	expRets := []int{0, 0, 100} // 期望Result中返回的值
	for i, req := range reqs {
		reqCh <- req
		waitResp(req, expRets[i])
	}
}
