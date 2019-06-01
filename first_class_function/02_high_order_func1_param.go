package main

import "fmt"

type Operation func(a, b int) int

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

type Calculator struct {
	v int
}

func (c *Calculator) Do(op Operation, a int) {
	c.v = op(c.v, a)
}

func main() {
	var calc Calculator

	calc.Do(Add, 1) // c.v = 1
	calc.Do(Sub, 2) // c.v = -1

	fmt.Println(calc.v) // -1
}
