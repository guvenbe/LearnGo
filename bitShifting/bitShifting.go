package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println("vim-go")
	x := 4
	fmt.Printf("%d\t%b\t%#X\n", x, x, x)
	y := x << 2
	fmt.Printf("%d\t%b\t%#X\n", y, y, y)
	z := x << 3
	fmt.Printf("%d\t%b\t%#X\n", z, z, z)
	v := x << 4
	fmt.Printf("%d\t%b\t%#X\n", v, v, v)
	fmt.Printf("%d\t\t\t%b\t\t\t%#X\n", kb, kb, kb)
	fmt.Printf("%d\t\t\t%b\t\t\t%#X\n", mb, mb, mb)
	fmt.Printf("%d\t\t%b\t\t%#X\n", gb, gb, gb)
}
