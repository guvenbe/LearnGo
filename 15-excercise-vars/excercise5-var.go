package main

import "fmt"

var x int

type hotdog int

var y hotdog
var z int

func main() {
	fmt.Printf("%d %T\n", x, x)
	x = 42
	fmt.Printf("%d %T\n", x, x)

	fmt.Printf("%d %T\n", y, y)
	y = 42
	fmt.Printf("%d %T\n", y, y)

	fmt.Printf("%d %T\n", z, z)
	z = 45
	fmt.Printf("%d %T\n", z, z)

	y = hotdog(z)
	fmt.Printf("%d %T\n", y, y)

	a := int(y)
	fmt.Printf("%d %T\n", a, a)
}
