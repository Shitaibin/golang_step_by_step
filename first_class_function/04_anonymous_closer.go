package main

import (
	"fmt"
	"time"
)

func test1() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range sl {
		go func() {
			fmt.Printf("%d %d\n", i, v)
		}()
	}

	time.Sleep(time.Second)
}

func test2() {

	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range sl {
		go func(a, b int) {
			fmt.Printf("%d %d\n", a, b)
		}(i, v)
	}
}

func main() {
	test1()
	time.Sleep(time.Second)
	fmt.Println("--------------")
	test2()
	time.Sleep(time.Second)
}

// ➜  first_class_function git:(master) ✗ go run 04_anonymous_closer.go
// 9 9
// 9 9
// 9 9
// 9 9
// 9 9
// 9 9
// 9 9
// 9 9
// 9 9
// 9 9
// --------------
// 0 0
// 2 2
// 5 5
// 9 9
// 4 4
// 7 7
// 6 6
// 8 8
// 1 1
// 3 3
