package main

import "fmt"

const (
	a         = 42
	b float64 = 42.78
	c         = "Bora Guven"
)

const (
	a=iota
	b
	c
)

fmt.println("------")
fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
const PI = 3.14

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")
	for i, v := range s {
		fmt.Println(i, v)
	}

	for i, v := range s {
		fmt.Printf("Index: %d => hex:%#U", i, v)
	}
	s1 := `H`
	fmt.Println(s1)
	n := []byte(s1)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
}
