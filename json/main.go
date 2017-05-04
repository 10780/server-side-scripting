package main

import (
"encoding/json"
"fmt"
)

type person struct {
First string
Last string
Items []string
}

func main(){
p1 := person {
First: "Harry",
Last: "Enfield",
Items: []string{"Loadsamoney", "Wad of pound notes"},
}

xp := []person{p1}

fmt.Println("go data: %+v\n", xp)

bs, err := json.Marshal(xp)
if err != nil {
fmt.Println(err)
}

fmt.Println("json:", string(bs))
}
