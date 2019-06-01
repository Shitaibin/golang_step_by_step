package main

import "fmt"

type Operation func(b int) int

func Add(b int) Operation {
	return func(a int) int {
		return a + b
	}
}

func Sub(b int) Operation {
	return func(a int) int {
		return a - b
	}
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
