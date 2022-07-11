package main

import (
	"fmt"
)

type person struct {
	first, last string
	age         int
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "MoneyPenny",
		age:   28,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		first: "Collision something",
		ltk:   true,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)
	fmt.Println(sa1.ltk, sa1.person.first, sa1.first, sa1.last, sa1.age)
	//Anonymous Struct
	p3 := struct {
		first, last string
		age         int
	}{
		first: "Bob",
		last:  "James",
		age:   76,
	}
	fmt.Println(p3)
}
