package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) getAge() string {
	return p.Name
}

func (p *Person) GetAge() int {
	return p.Age
}

func main() {
	p := &Person{
		"Jack", 10,
	}

	v := reflect.ValueOf(p)
	para := make([]reflect.Value, 0)

	age := v.MethodByName("GetAge").Call(para)
	fmt.Printf("Age: %d\n", age[0].Interface().(int))

	name := v.MethodByName("getName").Call(para)
	fmt.Printf("Name: %s\n", name[0].Interface().(string))
}
