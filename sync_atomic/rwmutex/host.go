package rwmutex

import (
	"sync"

	"../peer"
)

// Host 代表当前节点的连接管理
type Host struct {
	peers map[string]*peer.Peer // 连接上的所有Peer，根据Peer.ID索引
	lock  sync.RWMutex          // 保护peers互斥访问
	// 其他字段省略
}

func NewHost() *Host {
	h := &Host{
		peers: make(map[string]*peer.Peer),
	}
	return h
}

func (h *Host) AddPeer(p *peer.Peer) {
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
func (h *Host) GetPeer(pid string) *peer.Peer {
	h.lock.RLock()
	defer h.lock.RUnlock()

	return h.peers[pid]
}
