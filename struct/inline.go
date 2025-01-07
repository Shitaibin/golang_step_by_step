package main

import "fmt"

func main() {
	p1 := &Person{
		Animal: Animal{
			Area: "USA",
		},
		Name: "kite",
	}

	p2 := Person{
		Animal: Animal{
			Area: "UK",
		},
		Name: "james",
	}

	fmt.Printf("p1 addr: %p, values: %v\n", &p1, p1)
	p1.Update(p2)
	fmt.Println("After Update")
	fmt.Printf("p1 addr: %p, values: %v\n", &p1, p1)
}

type Person struct {
	Animal
	Name string
}

type Animal struct {
	Area string
}

func (p *Person) Update(o Person) {
	if o.Name != "" {
		p.Name = o.Name
	}
	p.Animal.Update(o.Animal)
}

func (a *Animal) Update(o Animal) {
	if o.Area != "" {
		a.Area = o.Area
	}
}
