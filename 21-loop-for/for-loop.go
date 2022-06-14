package main

import "fmt"

func main() {
	fmt.Println("starting")
	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}

	fmt.Println("starting1")
	y := 1
	for {
		if y > 9 {
			break

		}
		fmt.Println(y)
		y++
	}
	fmt.Println("starting2")
	z := 1
	for {
		z++
		if z > 100 {
			break
		}
		if z%2 != 0 {
			continue
		}
		fmt.Println(z)
	}
	for a := 33; a <= 122; a++ {
		fmt.Printf("%v\t%#X\t%#U\n", a, a, a)
	}
}
