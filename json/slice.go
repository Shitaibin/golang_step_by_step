package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	correct()
	wrong()
}

func correct() {
	buf := []byte(`["a","b"]`)
	var strs []string
	// 切片本身是一个结构体，指向一片buffer，所以入参应当为 strs的指针
	err := json.Unmarshal(buf, &strs)
	fmt.Printf("%v, %v\n", err, strs)
}

func wrong() {
	buf := []byte(`["a","b"]`)
	var strs []string
	// 传入非指针得到nil错误
	err := json.Unmarshal(buf, strs)
	fmt.Printf("%v, %v\n", err, strs)
}
