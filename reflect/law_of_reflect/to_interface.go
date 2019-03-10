package main

import (
	"fmt"
	"reflect"
)

// 使用Value.Interface()获取真实值并打印
func printValue(v reflect.Value) {
	fmt.Println("value:", v.Interface())
}

func main() {
	a := 1
	b := 2.1234
	c := "three"
	d := false
	e := []int{1, 2, 3}

	printValue(reflect.ValueOf(a))
	printValue(reflect.ValueOf(b))
	printValue(reflect.ValueOf(c))
	printValue(reflect.ValueOf(d))
	printValue(reflect.ValueOf(e))

	// 特殊，Kind是Int
	type MyInt int
	var x MyInt = 10
	printValue(reflect.ValueOf(x))
	fmt.Println("type:", reflect.TypeOf(x))
}

// 结果：
// value: 1
// value: 2.1234
// value: three
// value: false
// value: [1 2 3]
// value: 10
// type: main.MyInt
