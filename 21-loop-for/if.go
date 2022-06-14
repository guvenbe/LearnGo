package main

import "fmt"

func main() {
	if x := 42; x == 42 {
		fmt.Println("001")
	}
	fmt.Println("002")
	fmt.Println("003")

	y := 42
	if y == 42 {
		fmt.Println("42")
	} else if y == 41 {
		fmt.Println("41")
	} else {
		fmt.Println("40")

	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
