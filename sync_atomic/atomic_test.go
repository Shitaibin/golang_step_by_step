package benchtest

import (
	"testing"

	atomic "./atomic"
	channel "./channel"
	mu "./mutex"
	peer "./peer"
	rwmu "./rwmutex"
)

func BenchmarkMutexHost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mutexTest()
	}
}

func BenchmarkRWMutexHost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rwmutexTest()
	}
}

func BenchmarkAtomicHost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		atomicTest()
	}
}

func BenchmarkChannelHost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channelTest()
	}
}

type Op struct {
	op    string
	param string
}

var opList []Op

func init() {
	print("init opList\n")
	opList = append(opList, Op{"add", "peer1"})
	for i := 0; i < 1000000; i++ {
		opList = append(opList, Op{"get", "peer1"})
	}
	opList = append(opList, Op{"rm", "peer1"})
}

func mutexTest() {
	h := mu.NewHost()
	doTest(h.AddPeer, h.RemovePeer, h.GetPeer)
}

func rwmutexTest() {
	h := rwmu.NewHost()
	doTest(h.AddPeer, h.RemovePeer, h.GetPeer)
}

func channelTest() {
	h := channel.NewHost()
	h.Start()
	defer h.Stop()

	doTest(h.AddPeer, h.RemovePeer, h.GetPeer)
}

func atomicTest() {
	h := atomic.NewHost()
	doTest(h.AddPeer, h.RemovePeer, h.GetPeer)
}

func doTest(addFunc func(*peer.Peer), rmFunc func(string), getFunc func(string) *peer.Peer) {
	for _, op := range opList {
		switch op.op {
		case "add":
			go addFunc(&peer.Peer{ID: op.param})
		case "rm":
			go rmFunc(op.param)
		case "get":
			go getFunc(op.param)
		}
	}
}
