package main

import "fmt"

func main() {
	var m map[int]int
	_, ok := m[1]
	if !ok {
		fmt.Println("not panic")
	}
}
