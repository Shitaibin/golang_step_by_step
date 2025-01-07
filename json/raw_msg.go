package main

import (
	"encoding/json"
	"fmt"
)

// 测试marshal生成json rawMessage
type Person struct {
	Name  string
	Other *json.RawMessage
}

func main() {
	p := &Person{
		Name:  "John",
		Other: &json.RawMessage{},
	}

	attr := map[string]string{
		"age": "30",
	}
	attrBuf, err := json.Marshal(attr)
	if err != nil {
		panic(err)
	}
	p.Other = (*json.RawMessage)(&attrBuf)

	pBuf, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pBuf))
}
