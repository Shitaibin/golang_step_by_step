package main

import (
	"fmt"
	"reflect"
)

func main() {
	setable()
	canSetExample()
	canNotSetExample()
}

func setable() {
	x := 10
	v1 := reflect.ValueOf(x)
	fmt.Println("setable:", v1.CanSet())
	p := reflect.ValueOf(&x)
	fmt.Println("setable:", p.CanSet())
	v2 := p.Elem()
	fmt.Println("setable:", v2.CanSet())

	// 结果
	// setable: false
	// setable: false
	// setable: true
}

// 增加recover
func canNotSetExample() {
	x := 10
	v := reflect.ValueOf(x)
	changeToSeven(v)
	fmt.Println("value outside:", v.Interface())

	// 结果
	// value outside: 7
}

// BUG, it shoule be failed and panic
func canSetExample() {
	x := 10
	v := reflect.ValueOf(&x).Elem()
	changeToSeven(v)
	fmt.Println("value outside:", v.Interface())

	// 结果
	// panic: reflect: reflect.Value.SetInt using unaddressable value
}

// Can not set, if you set a non-setability Value, it will panic
func changeToSeven(v reflect.Value) {
	v.SetInt(7)
}
