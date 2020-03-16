package main

import (
	"log"
	"runtime"
	"time"

	"github.com/google/gops/agent"
)

func main() {
	if err := agent.Listen(agent.Options{
		Addr: "0.0.0.0:8848",
		// ConfigDir:       "/home/centos/gopsconfig", // 最好使用默认
		ShutdownCleanup: true}); err != nil {
		log.Fatal(err)
	}

	_ = make([]int, 1000, 1000)
	runtime.GC()

	_ = make([]int, 1000, 2000)
	runtime.GC()

	time.Sleep(time.Hour)
}
