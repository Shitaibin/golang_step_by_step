package main

import (
	"fmt"
	"reflect"
)

// 使用Value.Kind()类型判断，并获取真实值进行打印
func printValue(i interface{}) {
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Int:
		fmt.Println("value:", v.Int())
	case reflect.Float64:
		fmt.Println("value:", v.Float())
	case reflect.String:
		fmt.Println("value:", v.String())
	case reflect.Bool:
		fmt.Println("value:", v.Bool())
	default:
		fmt.Println("type is:", v.Type())
	}
}

func main() {
	a := 1
	b := 2.1234
	c := "three"
	d := false
	e := []int{1, 2, 3}

	printValue(a)
	printValue(b)
	printValue(c)
	printValue(d)
	printValue(e)

	// 特殊，Kind是Int
	type MyInt int
	var x MyInt = 10
	printValue(x)
	fmt.Println("type:", reflect.TypeOf(x))
}

// 结果：
// value: 1
// value: 2.1234
// value: three
// value: false
// type is: []int
// value: 10
// type: main.MyInt
