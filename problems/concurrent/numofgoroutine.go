package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	http.ListenAndServe("0.0.0.0:6060", nil)

	for {
		fmt.Printf("#goroutines: %d\n", runtime.NumGoroutine())
	}
}
