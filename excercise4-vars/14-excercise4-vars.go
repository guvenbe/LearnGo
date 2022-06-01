package main

import "fmt"

var x int

type hotdog int

var y hotdog

func main() {
	fmt.Printf("%d %T\n", x, x)
	x = 42
	fmt.Printf("%d %T\n", x, x)

	fmt.Printf("%d %T\n", y, y)
	y = 42
	fmt.Printf("%d %T\n", y, y)
}
