package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func testUnix() {
	fmt.Println("Test unix.Exec")
	unix.Exec("/bin/ls", []string{"ls", "-al", "/Users/shitaibin/pprof"}, os.Environ())
}

func testCmd() {
	fmt.Println("Test os.Command")
	cmd := exec.Command("/bin/ls", "-al", "/Users/shitaibin/pprof")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}
	defer stdout.Close()

	cmd.Start()
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		fmt.Println(err)
	} else {
		fmt.Println(string(opBytes))
	}
}

func main() {
	testCmd()
	fmt.Println("--------------")
	testUnix()
}
