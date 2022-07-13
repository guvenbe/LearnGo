package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(loopFact(4))

}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFact(n int) int {

	total := 1
	for ; n > 0; n-- {
		total *= n

	}
	return total
}
