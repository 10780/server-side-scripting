package main

import (
	"fmt"
)

func main(){
xs := []int{1, 2, 3, 7, 8, 9}
fmt.Println(xs, "\n")

//index
for i, _ := range xs{
fmt.Println(i)
}

//index and vaues
for i, _ := range xs{
fmt.Println(i, "=", xs[i])
}
}
