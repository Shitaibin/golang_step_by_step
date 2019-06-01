package main

import "fmt"

type Operation func(b int) int

func Add(b int) Operation {
	addB := func(a int) int {
		return a + b
	}
	return addB
}

func Sub(b int) Operation {
	subB := func(a int) int {
		return a - b
	}
	return subB
}

type Calculator struct {
	v int
}

func (c *Calculator) Do(op Operation) {
	c.v = op(c.v)
}

func main() {
	var calc Calculator

	calc.Do(Add(1)) // c.v = 1
	calc.Do(Sub(2)) // c.v = -1

	fmt.Println(calc.v) // -1
}
