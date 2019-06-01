package main

import "./host"

func main() {
	h := host.NewHost()
	h.AddPeer(&host.Peer{"peer1"})
	h.AddPeer(&host.Peer{"peer2"})
	h.AddPeer(&host.Peer{"peer3"})

	h.RemovePeer("peer2")
	h.BroadcastMsg("hi hangzhou")
}
