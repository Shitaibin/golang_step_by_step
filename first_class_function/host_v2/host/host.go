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

// Host 代表当前节点的连接管理
type Host struct {
	add       chan *Peer
	broadcast chan string
	remove    chan string
	stop      chan struct{}
}

func NewHost() *Host {
	h := &Host{
		add:       make(chan *Peer),
		broadcast: make(chan string),
		remove:    make(chan string),
		stop:      make(chan struct{}),
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
		case p := <-h.add:
			peers[p.ID] = p
		case pid := <-h.remove:
			delete(peers, pid)
		case msg := <-h.broadcast:
			for _, p := range peers {
				p.WriteMsg(msg)
			}
		case <-h.stop:
			return
		}
	}
}

func (h *Host) AddPeer(p *Peer) {
	h.add <- p
}

func (h *Host) RemovePeer(pid string) {
	h.remove <- pid
}

func (h *Host) BroadcastMsg(msg string) {
	h.broadcast <- msg
}

// GetPeer 当前Peer不存在时返回nil。
func (h *Host) GetPeer(pid string) *Peer {
	// 只有这3个channel无法实现
	return nil
}

// 使用数据的流动，实现并发
