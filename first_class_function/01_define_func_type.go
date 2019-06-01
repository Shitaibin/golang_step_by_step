package main

import "fmt"

type Operation func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	var op Operation
	op = Add
	fmt.Println(op(1, 2)) // 3
}
