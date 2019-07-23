package main

import "fmt"

func main() {
	m := make(map[int]int)

	fmt.Printf("addr in main: %p, len(m): %d\n", &m, len(m))
	test(m)
	fmt.Printf("addr in main: %p, len(m): %d\n", &m, len(m))
}

func test(m map[int]int) {
	fmt.Printf("addr in test: %p, len(m): %d\n", &m, len(m))
	m[1] = 1
	fmt.Printf("addr in test: %p, len(m): %d\n", &m, len(m))
}
