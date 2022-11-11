// 持续append，然后缩小，内存使用量是否会随着时间不断增大？
// 观察1小时，会增长，但也会下降。老的slice会被gc，新的slice大小跟数据量有关，数据量不大时，新slice就可以变小。
//
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

// 运行一段时间：fatal error: runtime: out of memory
func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	tick := time.Tick(time.Second / 100)
	var buf []int
	for range tick {
		buf = appendAndReduce(buf)
	}
}

func appendAndReduce(buf []int) []int {
	buf = append(buf, 1)
	buf = buf[1:]
	return buf
}
