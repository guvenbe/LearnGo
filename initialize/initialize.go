package main

import "fmt"

type Bootcamp struct {
	Lat, Lon float64
}

func main() {
	x := new(Bootcamp)
	y := &Bootcamp{}
	fmt.Println(x, y)
	fmt.Println(*x == *y)
}
