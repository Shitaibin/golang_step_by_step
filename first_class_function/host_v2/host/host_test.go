package host

import (
	"testing"
	"time"
)

func TestHost(t *testing.T) {
	h := NewHost()
	h.Start()

	h.AddPeer(&Peer{"peer1"})
	h.AddPeer(&Peer{"peer2"})

	p := h.GetPeer("peer1")
	if p == nil {
		t.Errorf("want find peer1, but got nothing")
	}

	p = h.GetPeer("peer3")
	if p != nil {
		t.Errorf("want got nothing of peer3, but got: %v", p.ID)
	}

	h.RemovePeer("peer1")
	p = h.GetPeer("peer1")
	if p != nil {
		t.Errorf("want got nothing of peer1, but got: %v", p.ID)
	}

	time.Sleep(time.Second)
	h.Stop()
}
