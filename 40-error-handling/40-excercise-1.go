package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First, Last string
	Sayings     []string
}

func main() {

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken not stirred", "Any last wishes?", "Never say Never"},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Panic("JSON did not marshal:", err)
	}
	fmt.Println(string(bs))
}
