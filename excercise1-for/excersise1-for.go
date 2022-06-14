package main

import "fmt"

func main() {
	bd := 1966
	for {
		if bd > 2022 {
			break
		}

		fmt.Println(bd)
		bd++

	}

	for i := 10; i <= 100; i++ {
		fmt.Printf("%v divided by 4 the modulus is: %v\n", i, i%4)
	}
}
