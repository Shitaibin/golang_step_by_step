package main

import (
	"fmt"
	"unsafe"
)

type A1 struct {
	b B1
	c *C1
	d *D1
}

func (a *A1) String() string {
	return fmt.Sprintf("A1{b: %s, c: %s, d: %s}", a.b, a.c, a.d)
}

type B1 struct {
	i    int
	name string
}

func (b B1) String() string {
	return fmt.Sprintf("B1{i: %d, name: %v}", b.i, b.name)
}

type C1 struct {
	s []int
}

func (c *C1) String() string {
	return fmt.Sprintf("C1{s: %v}", c.s)
}

type D1 struct {
	s []string
}

func (d *D1) String() string {
	return fmt.Sprintf("D1{s: %v}", d.s)
}

//------------------------------

type A2 struct {
	b B2
	c *C2
	d *D2
}

type B2 struct {
	i    int
	name string
}

type C2 struct {
	s []int
}

type D2 struct {
	s []string
}

func (a *A2) String() string {
	return fmt.Sprintf("A2{b: %s, c: %s, d: %s}", a.b, a.c, a.d)
}

func (b B2) String() string {
	return fmt.Sprintf("B2{i: %d, name: %v}", b.i, b.name)
}

func (c *C2) String() string {
	return fmt.Sprintf("C2{s: %v}", c.s)
}

func (d *D2) String() string {
	return fmt.Sprintf("D2{s: %v}", d.s)
}

func (d *D2) list() {
	fmt.Printf("d2: ")
	for _, v := range d.s {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

//------------------------------

func test1() {
	b1 := B1{10, "type B1"}
	c1 := &C1{[]int{1, 2, 3}}
	d1 := &D1{[]string{"type D1", "type D1"}}
	a1 := &A1{b1, c1, d1}
	fmt.Println(a1)

	a2 := (*A2)(unsafe.Pointer(a1))
	fmt.Println(a2)

	a2.d.list()
}

func main() {
	test1()
}
