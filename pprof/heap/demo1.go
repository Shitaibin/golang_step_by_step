// 展示内存增长
package main

import (
	"time"
)

// 运行一段时间：fatal error: runtime: out of memory
func main() {
	tick := time.Tick(time.Second / 100)
	var buf []byte
	for range tick {
		buf = append(buf, make([]byte, 1024*1024)...)
	}
}
