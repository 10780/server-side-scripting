package main

import (
	"fmt"
)

type person struct{
fname string
lname string
food []string
}

func main() {
p1 := person{"Perry", "Ball", []string{"steak", "vegetables", "skittles"}}
fmt.Println(p1.fname, "\n")
fmt.Println(p1.food, "\n")

for pvalFood, _ := range p1.food{
fmt.Println(pvalFood)
}
}
