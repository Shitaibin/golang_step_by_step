package main

import (
	"log"
	"runtime"
	"runtime/debug"
	"time"
)

var lastTotalFreed uint64

func main() {
	printMemStats("start")
	sleep()

	l := make([]int32, 1024*512)
	printMemStats("after create")
	sleep()

	// 使用1下防止回收
	for i := 0; i < len(l); i++ {
		l[i] = 10
	}

	l = nil
	printMemStats("after free")
	sleep()

	runtime.GC()
	printMemStats("after gc")
	sleep()

	debug.FreeOSMemory()
	printMemStats("after free os")
	sleep()
}

func sleep() {
	time.Sleep(time.Second * 10)
}

func printMemStats(tag string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%15s: Alloc = %vMB TotalAlloc = %vMB  Just Freed = %vMB Sys = %vMB NumGC = %v\n",
		tag, m.Alloc/1024, m.TotalAlloc/1024, ((m.TotalAlloc-m.Alloc)-lastTotalFreed)/1024, m.Sys/1024, m.NumGC)

	lastTotalFreed = m.TotalAlloc - m.Alloc
}
