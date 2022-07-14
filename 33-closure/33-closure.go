package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3
	fmt.Println(a()) //4
	fmt.Println(b()) //1
	fmt.Println(b()) //2
	fmt.Println(b()) //3
}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
