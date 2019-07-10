package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[1] = 2
	fmt.Printf("addr of m: %p\n", &m)
	// mp := (*Slice)(unsafe.Pointer(&m))
	// fmt.Printf("slice.array: %p, slice.len: 0x%x, slice.cap: 0x%x\n", mp.array, mp.len, mp.cap)
	test(m)
}

func test(m map[int]int) {
	fmt.Printf("addr of m: %p\n", &m)
	// mp := (*Slice)(unsafe.Pointer(&m))
	// fmt.Printf("slice.array: %p, slice.len: 0x%x, slice.cap: 0x%x\n", mp.array, mp.len, mp.cap)
	panic("want panic")
}

// 定义同runtime.hmap
// type Hmap struct {
// 	// Note: the format of the Hmap is encoded in ../../cmd/internal/gc/reflect.go and
// 	// ../reflect/type.go. Don't change this structure without also changing that code!
// 	count     int // # live cells == size of map.  Must be first (used by len() builtin)
// 	flags     uint8
// 	B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
// 	noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
// 	hash0     uint32 // hash seed

// 	buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
// 	oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
// 	nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

// 	extra *mapextra // optional fields
// }

// output：顺便解释一下slice的值传递，为何3个参数？是slice结构体的3个字段吗？
// addr of m: 0xc42000a060
// slice.array: 0xc42000e1e0, slice.len: 0x1, slice.cap: 0x1
// addr of m: 0xc42000a080
// slice.array: 0xc42000e1e0, slice.len: 0x1, slice.cap: 0x1
// panic: want panic

// goroutine 1 [running]:
// main.test(0xc42000e1e0, 0x1, 0x1)
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo5-slice2.go:20 +0x1c9
// main.main()
// 	/Users/shitaibin/Workspace/golang_step_by_step/stack_trace/demo5-slice2.go:13 +0x1f2
// exit status 2
