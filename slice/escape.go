package main

import (
	"fmt"
	"time"
)

func main() {
	example1()
	example2()
	example3()
	example4()
}

func example1() {
	fmt.Println("Example 1:")
	slice := []int{1,2,3}
	// moved to heap: n， 所以n实际分配在堆上，n的地址就不会变，并不会每次进入for循环都在堆上创建1个新的n
	// n的值依次会是1,2,3，但n的地址不变
	for _, n := range slice { 
		fmt.Printf("addr: %x v: %d\n", &n, n) 
	}
}

func example2() {
	fmt.Println("Example 2:")
	slice := []int{1,2,3}
	for i := range slice {
		fmt.Printf("addr: %x v: %d\n", &slice[i], slice[i])
	}
}

func example3() {
	fmt.Println("Example 3:")
	slice := []int{1,2,3}
	for _, n := range slice {
		fmt.Printf("addr: %x v: %d\n", &n, n)
		time.Sleep(time.Second)
	}
}

func example4() {
	fmt.Println("Example 4:")
	one,two,three := 1,2,3
	slice := []*int{&one,&two,&three}
	for _, ptr := range slice {
		fmt.Printf("addr of ptr: %x, ptr value: %x, ptr referenced value: %d\n", &ptr, ptr, *ptr)
	}
}