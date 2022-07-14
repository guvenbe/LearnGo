package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)

	fmt.Println(x())
	fmt.Println(bar()())

}

func bar() func() int {

	return func() int {
		return 451
	}
}
