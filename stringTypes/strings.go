package main

import "fmt"

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
	s := `H`
	fmt.Println(s)
	n := []byte(s)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%#X\n", n)
}
