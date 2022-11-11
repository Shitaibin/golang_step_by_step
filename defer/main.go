package main

import "fmt"

func main() {
	test()
}

func test() (msg string) {
	defer func() {
		fmt.Printf("catch: %s\n", msg)
	}()

	if msg := genMsg(); len(msg) != 0 {
		fmt.Printf("msg: %s\n", msg)
		return msg
	}
	return "no valid msg"
}

func genMsg() string {
	return "hello"
}
