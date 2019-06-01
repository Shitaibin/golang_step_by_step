package host

import (
	"fmt"
	"sync"
)

// Peer 代表1个节点
type Peer struct {
	ID string
	// Peer其他信息，比如网络连接、地址、协议类型等等
}

func (p *Peer) WriteMsg(msg string) {
	fmt.Printf("send to: %v, msg: %v\n", p.ID, msg)
}

// Host 代表当前节点的连接管理
type Host struct {
	peers map[string]*Peer // 连接上的所有Peer，根据Peer.ID索引
	lock  sync.RWMutex     // 保护peers互斥访问
	// 其他字段省略
}

func NewHost() *Host {
	h := &Host{
		peers: make(map[string]*Peer),
	}
	return h
}

func (h *Host) AddPeer(p *Peer) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.peers[p.ID] = p
}

func (h *Host) RemovePeer(pid string) {
	h.lock.Lock()
	defer h.lock.Unlock()

	delete(h.peers, pid)
}

// GetPeer 当前Peer不存在时返回nil。
func (h *Host) GetPeer(pid string) *Peer {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.peers[pid]
}

func (h *Host) BroadcastMsg(msg string) {
	h.lock.RLock()
	defer h.lock.RUnlock()

	for _, p := range h.peers {
		p.WriteMsg(msg)
	}
}

// 并发情况下，需要Mutex保护数据，保护临界区，特点：每增加1个操作，都要获取锁。这并不是Go的思维，这只是用Go语言写了C/C++，Java思维的代码
