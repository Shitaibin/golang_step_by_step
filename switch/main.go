package main

import "fmt"

func main() {
	a := 0

	switch a {
	// 0 will do nothing
	case 0:
	case 1:
		fmt.Println("1")
	}
}
