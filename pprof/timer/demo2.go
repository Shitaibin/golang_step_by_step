package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	// 开启pprof
	go func() {
		ip := "0.0.0.0:6060"
		if err := http.ListenAndServe(ip, nil); err != nil {
			fmt.Printf("start pprof failed on %s\n", ip)
			os.Exit(1)
		}
	}()

	// 每s起10个goroutine，goroutine会阻塞，不释放内存
	tick := time.Tick(time.Second / 10)
	i := 0
	for range tick {
		i++
		fmt.Println(i)
	}
}
