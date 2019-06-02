package host

import (
	"fmt"
)

// Peer 代表1个节点
type Peer struct {
	ID string
}

func (p *Peer) WriteMsg(msg string) {
	fmt.Printf("send to: %v, msg: %v\n", p.ID, msg)
}

type Operation func(peers map[string]*Peer)

// Host 代表当前节点的连接管理
type Host struct {
	opCh chan Operation
	stop chan struct{}
}

func NewHost() *Host {
	h := &Host{
		opCh: make(chan Operation),
		stop: make(chan struct{}),
	}
	return h
}

func (h *Host) Start() {
	go h.loop()
}

func (h *Host) Stop() {
	close(h.stop)
}

func (h *Host) loop() {
	peers := make(map[string]*Peer)

	for {
		select {
		case op := <-h.opCh:
			op(peers)
		case <-h.stop:
			return
		}
	}
}

func (h *Host) AddPeer(p *Peer) {
	add := func(peers map[string]*Peer) {
		peers[p.ID] = p
	}
	h.opCh <- add
}

func (h *Host) RemovePeer(pid string) {
	rm := func(peers map[string]*Peer) {
		delete(peers, pid)
	}
	h.opCh <- rm
}

func (h *Host) BroadcastMsg(msg string) {
	broadcast := func(peers map[string]*Peer) {
		for _, p := range peers {
			p.WriteMsg(msg)
		}
	}

	h.opCh <- broadcast
}

// GetPeer 当前Peer不存在时返回nil。
func (h *Host) GetPeer(pid string) *Peer {
	retCh := make(chan *Peer)
	query := func(peers map[string]*Peer) {
		retCh <- peers[pid]
	}

	// 发送查询
	go func() {
		h.opCh <- query
	}()

	// 等待查询结果并返回
	return <-retCh
}

// SendTo 只向某一个Peer发送消息
func (h *Host) SendTo(pid, msg string) {
	p := h.GetPeer(pid)
	p.WriteMsg(msg)
}

// NumOfPeers peer的数量
func (h *Host) NumOfPeers() int {
	retCh := make(chan int)
	query := func(peers map[string]*Peer) {
		retCh <- len(peers)
	}

	go func() {
		h.opCh <- query
	}()

	return <-retCh
}
