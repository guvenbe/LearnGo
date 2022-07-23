package main

import (
	"errors"
	"log"
)

var MyError = errors.New("square root of negative number")

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, MyError
	}
	return 42, nil
}
