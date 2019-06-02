package host

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

func TestHostWithSequence(t *testing.T) {
	h := NewHost()
	h.Start()
	defer h.Stop()

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
}

func TestHostWithConcurrency(t *testing.T) {
	for i := 0; i < 10000; i += rand.Intn(500) {
		testHostWithConcurrencyN(t, i)
	}
}

// testHostWithConcurrencyN 添加n个peer，然后随机删除已经加入的peer，最后统计peer数量是否正确
func testHostWithConcurrencyN(t *testing.T, n int) {
	h := NewHost()
	h.Start()
	defer h.Stop()

	// 添加n个peer
	pidCh := func() <-chan string {
		pidCh := make(chan string, n)

		go func() {
			defer close(pidCh)
			for i := 0; i < n; i++ {
				pid := "peer" + strconv.Itoa(i)
				go h.AddPeer(&Peer{pid})
				pidCh <- pid
			}
		}()

		return pidCh
	}()

	// 随机删除m个已添加的peer
	m := 0
	for pid := range pidCh {
		if rand.Intn(n)%2 == 0 {
			if h.GetPeer(pid) != nil {
				m += 1
				h.RemovePeer(pid)
			}
		}
	}

	// 统计剩下的peer是否为n-m个
	got := h.NumOfPeers()
	assert.Equal(t, got, n-m, "left peer mismatch")
}
