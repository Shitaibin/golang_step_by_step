package atomic

import (
	"sync"
	"sync/atomic"

	"../peer"
)

type Map map[string]*peer.Peer

// Host 代表当前节点的连接管理
type Host struct {
	peers atomic.Value // 连接上的所有Peer，根据Peer.ID索引
	lock  sync.Mutex
	// 其他字段省略
}

func NewHost() *Host {
	h := &Host{}
	h.peers.Store(make(Map))
	return h
}

func (h *Host) AddPeer(p *peer.Peer) {
	ps, _ := h.peers.Load().(Map)
	nps := make(Map)
	for k, v := range ps {
		nps[k] = v
	}
	nps[p.ID] = p
	h.lock.Lock()
	defer h.lock.Unlock()
	h.peers.Store(nps)
}

func (h *Host) RemovePeer(pid string) {
	ps, _ := h.peers.Load().(Map)
	nps := make(Map)
	for k, v := range ps {
		nps[k] = v
	}
	delete(nps, pid)
	h.lock.Lock()
	defer h.lock.Unlock()
	h.peers.Store(nps)
}

// GetPeer 当前Peer不存在时返回nil。
func (h *Host) GetPeer(pid string) *peer.Peer {
	ps, _ := h.peers.Load().(Map)
	return ps[pid]
}
