package main

import (
	"fmt"
	"log"
)

type northgateMathError struct {
	lat  string
	long string
	err  error
}

func (n northgateMathError) Error() string {
	return fmt.Sprintf("Northgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("square root of negative number: %v", f)
		return 0, northgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	return 42, nil
}
