package main

import (
	"fmt"
)

func main() {
m := map[string]int{"one": 1, "two": 2, "three": 3, "five": 5, "six": 6, "seven": 7}
fmt.Println(m, "\n")

for k, _ := range m{
fmt.Println(k, "\n")
}

for k, v := range m{
fmt.Println(k, ",", v)
}

}
