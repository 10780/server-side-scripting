package main

import (
	"fmt"
)

type person struct{
fname string
lname string
}

func main() {
p1 := person{"Perry", "Ball"}
fmt.Println(p1, "\n")

n := aNum()
fmt.Println(n)
}

/*IMPORTANT DO NOT FORGET
func (reciever) identifier(parameter) (returns){
code
}
*/
func aNum() int{
return 24
}
