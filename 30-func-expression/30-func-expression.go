package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("my first func expression")
	}

	f()

	g := func(x int) {
		fmt.Printf("my first func expression %d", x)
	}

	g(1984)
}
