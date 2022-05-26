package main

import (
	"fmt"
	"time"
)

type Point struct {
	X, Y int
}

type Bootcamp struct {
	Lat, Lon float64
	Date     time.Time
}

var (
	p = Point{1, 2}  // has type Point
	q = &Point{1, 2} // has type *Point
	r = Point{X: 1}  // Y:0 is implicit
	s = Point{}      // X:0 and Y:0
)

func main() {
	fmt.Println(p, q, r, s)

	event := Bootcamp{
		Lat: 34.012679,
		Lon: -118.456123,
	}

	event.Date = time.Now()
	fmt.Printf("Event on %s, location (%f, %f)",
		event.Date, event.Lat, event.Lon)
}
