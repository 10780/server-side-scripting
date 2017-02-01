package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
	walk string
}

func main() {
	p1 := person{"Moneypenny", 24, "is walking"}
	fmt.Println(p1)
	p1.foo()
}

func (p person) foo() {
	fmt.Println("Hello from", p.last)
	fmt.Println(p.last, p.walk)
}

// func (receiver) identifier(parameters) (returns) {}
