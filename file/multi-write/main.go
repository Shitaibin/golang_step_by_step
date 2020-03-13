package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	go pprof()

	go write("file1")
	go write("file2")

	select {}
}

func write(name string) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Printf("Create file failed: %v", err)
	}

	for {
		if _, err := f.Write([]byte("hello\n")); err != nil {
			fmt.Printf("Write %v error: %v", name, err)
			return
		}
	}
}

func pprof() {
	ip := "0.0.0.0:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
