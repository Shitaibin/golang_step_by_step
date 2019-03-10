package main

import (
	"fmt"
)

func main() {
	x := 10
	fmt.Printf("addr of x: %x\n", &x)

	toInterface(x)
}

func toInterface(i interface{}) {
	fmt.Printf("addr of i: %x\n", &i)
}
