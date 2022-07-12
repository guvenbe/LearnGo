package main

import "fmt"

func main() {

	func() {
		fmt.Printf("Anonymous function ran\n")
	}()
	func(x int) {
		fmt.Printf("Anonymous function ran %d\n", x)
	}(44)
}
