package channel

import "../peer"

type Operation func(peers map[string]*peer.Peer)

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
	peers := make(map[string]*peer.Peer)

	for {
		select {
		case op := <-h.opCh:
			op(peers)
		case <-h.stop:
			return
		}
	}
}

func (h *Host) AddPeer(p *peer.Peer) {
	add := func(peers map[string]*peer.Peer) {
		peers[p.ID] = p
	}
	h.opCh <- add
}

func (h *Host) RemovePeer(pid string) {
	rm := func(peers map[string]*peer.Peer) {
		delete(peers, pid)
	}
	h.opCh <- rm
}

// GetPeer 当前Peer不存在时返回nil。
func (h *Host) GetPeer(pid string) *peer.Peer {
	retCh := make(chan *peer.Peer)
	query := func(peers map[string]*peer.Peer) {
		retCh <- peers[pid]
	}

	// 发送查询
	go func() {
		h.opCh <- query
	}()

	// 等待查询结果并返回
	return <-retCh
}

// NumOfPeers peer的数量
func (h *Host) NumOfPeers() int {
	retCh := make(chan int)
	query := func(peers map[string]*peer.Peer) {
		retCh <- len(peers)
	}

	go func() {
		h.opCh <- query
	}()

	return <-retCh
}
