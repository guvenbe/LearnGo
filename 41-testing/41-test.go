// Package demo for sum
package main

import "fmt"

func main() {
	fmt.Println("2+3 =", MySum(2, 3))
	fmt.Println("4+7 =", MySum(4, 7))
	fmt.Println("5+9 =", MySum(5, 9))
}

//Sum adds an unlimited number of values of type int
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

func MySubraction(xi ...int) int {
	sub := 0
	for _, v := range xi {
		sub -= v
	}
	return sub
}
