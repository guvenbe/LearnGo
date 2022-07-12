package main

import "fmt"

type person struct {
	first, last string
}

type secretAgent struct {
	person
	ltk bool
}
type human interface {
	speak()
}

type hotdog int

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, " - the person speak")
}
func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed in to bar xxxxp", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed in to bar xxxxs", h.(secretAgent).first)
	}
	fmt.Println("I was passed in to bar", h)
	h.speak()
}
func main() {
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond"},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "Moneypenny"},
		ltk: true,
	}

	p1 := person{
		first: "Dr",
		last:  "yes",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	sa1.speak()
	sa2.speak()

	fmt.Println("---------------------")
	bar(sa1)
	bar(sa2)
	bar(p1)

	//coversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
