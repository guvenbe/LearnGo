package main

import "fmt"

type person struct {
	name string
}
type human interface{ speak() }

func (s *person) speak() {
	fmt.Println("hello:")
}
func saySomething(h human) {
	h.speak()
}
func main() {
	p := person{name: "Bob"}
	p.speak()
	saySomething(p)
	saySomething(&p)
}
